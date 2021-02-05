package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createDaemonSetCmd)
}

var createDaemonSetCmd = &cobra.Command{
	Use:   "create-daemonset",
	Short: "This command is used for creating DaemonSet workload",
	Long:  "This command is for creating DaemonSet using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDaemonSet()
	},
}
