package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create-deploy",
	Short: "This command is for creating deployment",
	Long:  "This command is used for creating deployment object using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDeployment()
	},
}
