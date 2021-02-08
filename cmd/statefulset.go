package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

var stsName string

func init() {
	rootCmd.AddCommand(createStsCmd)
	rootCmd.AddCommand(listStsCmd)
	rootCmd.AddCommand(deleteStsCmd)
	deleteStsCmd.PersistentFlags().StringVarP(&stsName, "name", "n", "", "This flag sets StatefulSet name to be deleted")
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

var deleteStsCmd = &cobra.Command{
	Use:   "delete-sts",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetStsName(stsName)
		api.DeleteStatefulSet()
	},
}
