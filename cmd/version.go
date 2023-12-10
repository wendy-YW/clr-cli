package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of clr-cli",
	Long:    `All software has versions. This is clr-cli' v`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Color CLI Tools v0.1.0 -- github@wendy-YW")
	},
}
