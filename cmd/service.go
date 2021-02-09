package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createSvcCmd)
}

var createSvcCmd = &cobra.Command{
	Use:   "create-svc",
	Short: "Create headless service",
	Long:  "Create a headless service that maintain stateful set",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateHeadlessService()
	},
}
