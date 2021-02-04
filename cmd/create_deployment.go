package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pkbhowmick/client-go/api"
)

func init()  {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "This command is for creating deployment",
	Long: "This command is used for creating deployment object using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDeployment()
	},
}
