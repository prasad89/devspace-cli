/*
Copyright © 2025 Prasad Bhalerao
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// devspaceCmd represents the devspace command
var devspaceCmd = &cobra.Command{
	Use:   "devspace",
	Short: "CLI for managing DevSpace instances.",
	Long:  "The DevSpace CLI serves as the client-side interface for interacting with the DevSpace API server. It allows users to authenticate, configure settings, and manage DevSpace instances.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := devspaceCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
