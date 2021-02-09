package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

var secretName string

func init() {
	rootCmd.AddCommand(createSecretCmd)
	rootCmd.AddCommand(deleteSecretCmd)
}

var createSecretCmd = &cobra.Command{
	Use:   "create-secret",
	Short: "Create Opaque type secret",
	Long:  "Create a secret of type Opaque in default namespace. It creates username & password secret data",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateSecret()
	},
}

var deleteSecretCmd = &cobra.Command{
	Use:   "delete-secret",
	Short: "Delete the given secret object",
	Long:  "Delete the given secret object of",
	Run: func(cmd *cobra.Command, args []string) {
		api.DeleteSecret(args)
	},
}
