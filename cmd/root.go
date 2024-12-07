package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "boilerplate",
	Short: "Boilerplace is a simple REST API application",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Base command")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
