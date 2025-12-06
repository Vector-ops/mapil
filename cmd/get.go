package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a key",
	Long:  `Use key to get the value`,
	Run: func(cmd *cobra.Command, args []string) {
		key := ""
		if len(args) > 0 {
			key = args[0]
		} else {
			fmt.Println("get command requires a key")
			return
		}

		values, err := DataStore.GetValue(key)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(values) == 0 {
			fmt.Println("Key not found.")
		} else {
			fmt.Println(key, ": ", strings.Join(values, ", "))
		}
	},
}
