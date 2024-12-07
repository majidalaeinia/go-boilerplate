package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var Version = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of boilerplate",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("boilerplate version %s\n", Version)
		return nil
	},
}
