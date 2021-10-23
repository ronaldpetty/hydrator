package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hydrator v0.0.6")
	},
}

func getVersionCmd() *cobra.Command {
	return versionCmd
}
