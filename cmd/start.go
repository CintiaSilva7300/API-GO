package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the API-GO application",
	Long:  `Start the API-GO application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the API-GO application...")

		startCmd := exec.Command("go", "run", "main.go")

		err := startCmd.Run()
		if err != nil {
			fmt.Printf("Failed to start the API-GO application: %s", err)

		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
