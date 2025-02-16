/*
Copyright © 2025 Prasad Bhalerao
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/markkurossi/tabulate"
	"github.com/prasad89/devspace-cli/internal"
	"github.com/prasad89/devspace-cli/models"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available DevSpace instances",
	Long:    "The list command displays all available DevSpace instances within your namespace. It provides an overview of active deployments and their configurations.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := internal.GetConfig()

		endpoint := cfg.Section("server").Key("endpoint").String()
		if endpoint == "" {
			fmt.Println("Error: API endpoint is not configured. Check 'devspace config --help' for more information.")
			os.Exit(1)
		}

		token := cfg.Section("auth").Key("token").String()
		if token == "" {
			fmt.Println("Error: Authentication token is missing. Check 'devspace login --help' for more information.")
			os.Exit(1)
		}

		req, err := http.NewRequest("GET", endpoint+"/devspaces", nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}

		req.Header.Set("Authorization", "Bearer "+token)

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

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Failed to list DevSpaces: %s\n", body)
			os.Exit(1)
		}

		var data struct {
			Devspaces []models.Devspace `json:"devspaces"`
		}

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			os.Exit(1)
		}

		devspaces := data.Devspaces

		if len(devspaces) == 0 {
			fmt.Println("No DevSpaces found.")
			return
		}

		tab := tabulate.New(tabulate.Unicode)
		tab.Header("Name").SetAlign(tabulate.ML)

		for _, devspace := range devspaces {
			row := tab.Row()
			row.Column(devspace.Name)
		}

		tab.Print(os.Stdout)
	},
}

func init() {
	devspaceCmd.AddCommand(listCmd)
}
