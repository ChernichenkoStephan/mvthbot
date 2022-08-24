package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   `version`,
	Short: `Print version number of mvthbot service`,
	Long:  `This command can be used get the version number of mvthbot service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`mvthbot v0.0.1`)
	},
}
