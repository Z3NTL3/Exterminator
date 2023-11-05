package cmd

import (
	"errors"

	"z3ntl3/exterminator/globals"

	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command = &cobra.Command{
	Use:   "run",
	Short: "Run CursedSpirits with fancy modifications or utilities",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdString, err := cmd.Flags().GetString("cmd")
		if cmdString == "" || err != nil {
			return errors.New("Empty 'cmd' flag or error occurance")
		}

		return nil
	},
}

func Init() {
	globals.ClientArgs.CommandString = RootCmd.Flags().String("cmd", "", "Flag arguments to pass to CursedSpirits (example: --url https://google.com --concurrency 4000)")
	globals.ClientArgs.Refresh = RootCmd.Flags().Int("refresh", 20, "Sets the refresh ratio in to given seconds")
	globals.ClientArgs.ProxyRefresh = RootCmd.Flags().Bool("proxy-refresh", false, "Whether to also refresh proxy list")
	globals.ClientArgs.Bin = RootCmd.Flags().String("bin", "$HOME/CursedSpirits", "If you have a custom installation folder, provide the location to the CursedSpirits binary executable file")
}
