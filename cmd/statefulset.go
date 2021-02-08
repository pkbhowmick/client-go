package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createStsCmd)
	rootCmd.AddCommand(listStsCmd)
}

var createStsCmd = &cobra.Command{
	Use:   "create-sts",
	Short: "This command is for updating current deployment",
	Long:  "This command is used for updating current deployment using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateStatefulSet()
	},
}

var listStsCmd = &cobra.Command{
	Use:   "list-sts",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		api.ListStatefulSet()
	},
}
