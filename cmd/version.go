package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Mapil",
	Long:  `All software has versions. This is Mapil's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(info.Main.Version)
	},
}
