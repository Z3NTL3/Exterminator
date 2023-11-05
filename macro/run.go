package macro

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
)

type Composer interface {
	Installer
	Exterminator
}

type Installer interface {
	RunInstallationCmd(cmd CmdCtx, out io.Writer, id int)
}

type Exterminator interface {
	Run(cmd CmdCtx, out io.Writer, ctx context.Context, done chan int)
}

type Client struct {
	Err error
}

type CmdCtx struct {
	Name string
	Args []string
}

/*
The installer function that takes care of the entire installation

“cmd“ is the command to execute and “out“ is the output which the process's stdout is bind to
*/
func (c *Client) RunInstallationCmd(cmd CmdCtx, out io.Writer, id int) {
	// Abort installer on any failure
	if c.containsErr() {
		os.Exit(-1)
	}

	command := exec.Command(cmd.Name, cmd.Args...)
	{
		command.Stdout = out
		command.Stderr = out
	}

	if err := command.Run(); err != nil {
		c.Err = err
		return
	}

	log.Printf("[INFO] Installation thread[%d], succeeded!\r\n", id)
}

// Scan and check if any error occured on the installation process
func (c *Client) containsErr() bool {
	if c.Err != nil && c.Err != io.EOF {
		return true
	}
	return false
}

func (c *Client) Run(cmd CmdCtx, out io.Writer, ctx context.Context, done chan int) {
	if err := exec.Command(cmd.Name, "scrape").Run(); err != nil {
		out.Write([]byte(err.Error()))
		log.Fatal(err)
	}

	if err := exec.Command("bash", "-c", "-n", "999999").Run(); err != nil {
		out.Write([]byte(err.Error()))
		log.Fatal(err)
	}

	command := exec.CommandContext(ctx, cmd.Name, cmd.Args...)
	{
		command.Stdout = out
		command.Stderr = out
	}

	if err := command.Start(); err != nil {
		command.Stderr.Write([]byte(err.Error()))
		done <- 1
	}

	if err := command.Wait(); err != nil {
		// process kill; after refresh ratio accomplisment
		command.Stderr.Write([]byte(err.Error()))
		command.Stderr.Write([]byte("\x1b[1m[INFO] \x1b[0mRefreshing proxies and attack!\r\n"))
		done <- 1
	}
}
