package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

var stsName string
var image string

func init() {
	rootCmd.AddCommand(createStsCmd)
	rootCmd.AddCommand(listStsCmd)
	rootCmd.AddCommand(deleteStsCmd)
	rootCmd.AddCommand(updateStsCmd)
	deleteStsCmd.PersistentFlags().StringVarP(&stsName, "name", "n", "", "This flag sets StatefulSet name to be deleted")
	updateStsCmd.PersistentFlags().StringVarP(&stsName, "name", "n", "", "This flag sets StatefulSet name to be updated")
	updateStsCmd.PersistentFlags().Int32VarP(&replicas, "replicas", "r", 1, "This flag sets the number of replicas of the StatefulSet")
	updateStsCmd.PersistentFlags().StringVarP(&image, "image", "i", "mongo", "This flag sets image name to be updated")
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
	Short: "Return the list of StatefulSet",
	Long:  "list-sts command return the list of all StatefulSet running on the default namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		api.ListStatefulSet()
	},
}

var deleteStsCmd = &cobra.Command{
	Use:   "delete-sts",
	Short: "Delete the given StatefulSet",
	Long:  "delete-sts command deletes the given StatefulSet",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetStsName(stsName)
		api.DeleteStatefulSet()
	},
}

var updateStsCmd = &cobra.Command{
	Use:   "update-sts",
	Short: "Update the replica number",
	Long:  "Update-sts updates the number of replicas of the given StatefulSet",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetStsName(stsName)
		api.SetReplicas(replicas)
		api.SetImage(image)
		api.UpdateStatefulSet()
	},
}
