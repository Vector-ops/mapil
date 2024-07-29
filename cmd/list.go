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
		data := DataStore.GetAllData()
		if len(data) == 0 {
			fmt.Println("Data store empty.")
		} else {
			for _, v := range data {
				fmt.Println(v.Key, ": ", v.Value)
			}
		}
	},
}
