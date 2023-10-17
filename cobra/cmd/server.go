package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "study-server",
		Short:   "start the study-server",
		Example: "go-study study-server",
		Args: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Received arguments: %v\n", args)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Received arguments: %v\n", args)
			setup()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	fmt.Println("setup")
}
