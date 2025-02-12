/*
Copyright © 2025 Prasad Bhalerao
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// devspaceCmd represents the devspace command
var devspaceCmd = &cobra.Command{
	Use:   "devspace",
	Short: "CLI for managing DevSpace(s)",
	Long:  `The CLI, DevSpace, serves as the client-side interface for interacting with the DevSpace API server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("devspace called")
	},
}

func Execute() {
	err := devspaceCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
