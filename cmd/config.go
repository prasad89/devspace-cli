/*
Copyright © 2025 Prasad Bhalerao
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/prasad89/devspace-cli/internal"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the DevSpace CLI",
	Long:  "The config command is used to configure the DevSpace CLI. It allows you to set the API endpoint and other necessary settings.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, configPath := internal.GetConfig()

		endpoint, _ := cmd.Flags().GetString("endpoint")

		cfg.Section("server").Key("endpoint").SetValue(endpoint)
		if err := cfg.SaveTo(configPath); err != nil {
			fmt.Printf("Failed to update config file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Updated config file at", configPath)
	},
}

func init() {
	devspaceCmd.AddCommand(configCmd)

	configCmd.Flags().StringP("endpoint", "e", "", "Set the API endpoint for DevSpace CLI")
	configCmd.MarkFlagRequired("endpoint")
}
