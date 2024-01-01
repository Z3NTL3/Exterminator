package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"
        "path"
	"z3ntl3/exterminator/cmd"
	"z3ntl3/exterminator/globals"
	"z3ntl3/exterminator/macro"
)

func main() {
	cmd.Init()
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	if err := os.Chdir(path.Dir(*globals.ClientArgs.Bin)); err != nil {
		log.Fatal(err)
	}

	var c macro.Exterminator
	c = &macro.Client{}

	for {
		wait := make(chan int)
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(*globals.ClientArgs.Refresh)))

		go c.Run(macro.CmdCtx{
			Name: *globals.ClientArgs.Bin,
			Args: strings.Split("start "+*globals.ClientArgs.CommandString, " "),
		}, os.Stdout, ctx, wait)

		<-wait
		cancel()
	}
}
