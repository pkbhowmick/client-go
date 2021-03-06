package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

var deploymentName string
var replicas int32

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getDeployCmd)
	rootCmd.AddCommand(deleteDeployCmd)
	rootCmd.AddCommand(updateDeployCmd)
	deleteDeployCmd.PersistentFlags().StringVarP(&deploymentName, "name", "n", "go-api-server", "This flag sets the name of the deployment to be deleted")
	updateDeployCmd.PersistentFlags().Int32VarP(&replicas, "replicas", "r", 1, "This flag sets the number of replicas")
	updateDeployCmd.PersistentFlags().StringVarP(&deploymentName, "name", "n", "go-api-server", "This flag sets the name of the deployment to be deleted")

}

var createCmd = &cobra.Command{
	Use:   "create-deploy",
	Short: "This command is for creating deployment",
	Long:  "This command is used for creating deployment object using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDeployment()
	},
}

var getDeployCmd = &cobra.Command{
	Use:   "get-deploy",
	Short: "This command will list all deployment resources running",
	Long:  "This command will list all deployment resources that are running on",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetDeployment()
	},
}

var deleteDeployCmd = &cobra.Command{
	Use:   "delete-deploy",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetDeploymentName(deploymentName)
		api.DeleteDeployment()
	},
}

var updateDeployCmd = &cobra.Command{
	Use:   "update-deploy",
	Short: "This command updates the number of replicas of the given deployment",
	Long:  "This command updates the number of replicas of the given deployment using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetDeploymentName(deploymentName)
		api.SetReplicas(replicas)
		api.UpdateDeployment()
	},
}
