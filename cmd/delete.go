package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an object",
	Long:  `Delete an object from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		delObj()
	},
}

func delObj() {
	keyPrompt := promptui.Select{
		Label: "Select the key you want to delete.",
		Items: DataStore.GetKeys(),
	}

	_, key, err := keyPrompt.Run()
	if err != nil {
		fmt.Println("error while running ", err)
		return
	}
	DataStore.DeleteValue(key)
	DataStore.Persist()
	fmt.Println("Deleted object: ", key)
}
