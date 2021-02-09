package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createSecretCmd)
}

var createSecretCmd = &cobra.Command{
	Use:   "create-secret",
	Short: "Create Opaque type secret",
	Long:  "Create a secret of type Opaque in default namespace. It creates username & password secret data",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateSecret()
	},
}
