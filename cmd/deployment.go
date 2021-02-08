package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

var deploymentName string

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getDeployCmd)
	rootCmd.AddCommand(deleteDeployCmd)
	deleteDeployCmd.PersistentFlags().StringVarP(&deploymentName, "name", "n", "go-api-server", "This flag sets the name of the deployment to be deleted")
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
