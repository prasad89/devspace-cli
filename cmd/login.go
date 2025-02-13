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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with the DevSpace API server",
	Long:  "The login command is used to authenticate a user with the DevSpace API server. It validates user credentials and retrieves an authentication token for future requests.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, configPath := internal.GetConfig()

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		endpoint := cfg.Section("server").Key("endpoint").String()
		if endpoint == "" {
			fmt.Println("Error: API endpoint is not configured. Check 'devspace config --help' for more information.")
			os.Exit(1)
		}

		payload := map[string]string{
			"username": username,
			"password": password,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error creating JSON payload: %v\n", err)
			os.Exit(1)
		}

		resp, err := http.Post(endpoint+"/login", "application/json", bytes.NewBuffer(jsonData))
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

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Login failed: %s\n", body)
			os.Exit(1)
		}

		var responseData struct {
			Token string `json:"token"`
		}
		if err := json.Unmarshal(body, &responseData); err != nil {
			fmt.Printf("Error parsing response JSON: %v\n", err)
			os.Exit(1)
		}

		cfg.Section("auth").Key("username").SetValue(username)
		if err := cfg.SaveTo(configPath); err != nil {
			fmt.Printf("Failed to save username: %v\n", err)
			os.Exit(1)
		}

		cfg.Section("auth").Key("token").SetValue(responseData.Token)
		if err := cfg.SaveTo(configPath); err != nil {
			fmt.Printf("Failed to save authentication token: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Login successful!")
	},
}

func init() {
	devspaceCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "Username for authentication")
	loginCmd.MarkFlagRequired("username")
	loginCmd.Flags().StringP("password", "p", "", "Password for authentication")
	loginCmd.MarkFlagRequired("password")
}
