/*
Copyright © 2025 Prasad Bhalerao
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/prasad89/devspace-cli/internal"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new DevSpace instance",
	Long:  "The create command is used to create a new DevSpace instance with the specified configurations. It provisions the necessary resources and applies your settings to initialize the environment.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		cfg, _ := internal.GetConfig()

		endpoint := cfg.Section("server").Key("endpoint").String()
		token := cfg.Section("auth").Key("token").String()

		if endpoint == "" {
			fmt.Println("Error: API endpoint is not configured. Check 'devspace config --help' for more information.")
			os.Exit(1)
		}

		if token == "" {
			fmt.Println("Error: Authentication token is missing. Check 'devspace login --help' for more information.")
			os.Exit(1)
		}

		payload := map[string]string{
			"name": name,
		}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error creating request payload: %v\n", err)
			os.Exit(1)
		}

		req, err := http.NewRequest("POST", endpoint+"/devspace", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Failed to create DevSpace: %s\n", body)
			os.Exit(1)
		}

		fmt.Println("✅ DevSpace created successfully!")
		fmt.Println(string(body))
	},
}

func init() {
	devspaceCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name for the DevSpace")
	createCmd.MarkFlagRequired("name")
}
