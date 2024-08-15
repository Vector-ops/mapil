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
	if *delAll {
		DataStore.DeleteAll()
		DataStore.Persist()
		fmt.Printf("Deleted all data.\n")
		return
	}
	keys := DataStore.GetKeys()
	if len(keys) == 0 {
		fmt.Println("Data store empty.")
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "{{ . | red | cyan }}",
	}

	keyPrompt := promptui.Select{
		Label:     "Select the key you want to delete.",
		Items:     keys,
		Templates: templates,
	}

	_, key, err := keyPrompt.Run()
	if err != nil {
		fmt.Println("error while running ", err)
		return
	}
	DataStore.DeleteValue(key)
	DataStore.Persist()
	fmt.Printf("'%s' deleted.\n", key)
}
