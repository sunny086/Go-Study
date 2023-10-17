package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "go-admin",
	Short:        "go-admin",
	SilenceUsage: true,
	Long:         `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println("--------cobra rootCmd----------")
			fmt.Println("requires at least one arg")
			fmt.Println("--------cobra rootCmd----------")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--------cobra rootCmd----------")
		fmt.Println("Received arguments: %v\n", args)
		fmt.Println("--------cobra rootCmd----------")
	},
}

func init() {
	rootCmd.AddCommand(StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
