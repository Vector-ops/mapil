package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all objects",
	Long:  `All objects stored are listed`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range DataStore.GetAllData() {
			fmt.Println(v.Key, ": ", v.Value)
		}
	},
}
