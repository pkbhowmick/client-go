package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createJobCmd)
}

var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "This is command for creating Job",
	Long:  "This is the command for creating Kubernetes Job using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateJob()
	},
}
