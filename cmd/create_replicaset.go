package cmd

import (
	"github.com/pkbhowmick/client-go/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createReplicaSetCmd)
}

var createReplicaSetCmd = &cobra.Command{
	Use:   "create-replicaset",
	Short: "This command is for creating replica set",
	Long:  "This command is for creating replica set using kubernetes API",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateReplicaSet()
	},
}
