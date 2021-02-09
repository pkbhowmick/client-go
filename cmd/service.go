package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createSvcCmd)
	rootCmd.AddCommand(deleteSvcCmd)
}

var createSvcCmd = &cobra.Command{
	Use:   "create-svc",
	Short: "Create headless service",
	Long:  "Create a headless service that maintain stateful set",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateHeadlessService()
	},
}

var deleteSvcCmd = &cobra.Command{
	Use:   "delete-svc",
	Short: "Delete the given service",
	Long:  "Delete the services",
	Run: func(cmd *cobra.Command, args []string) {
		api.DeleteService(args)
	},
}
