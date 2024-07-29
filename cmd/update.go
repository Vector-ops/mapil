package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var updCmd = &cobra.Command{
	Use:   "upd",
	Short: "Update an object",
	Long:  `Update an object from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		updObj()
	},
}

func updObj() {
	selectPrompt := promptui.Select{
		Label: "Select the key you want to update.",
		Items: DataStore.GetKeys(),
	}

	_, key, err := selectPrompt.Run()
	if err != nil {
		fmt.Println("error while running ", err)
		return
	}
	validate := func(input string) error {
		if input == "" {
			return fmt.Errorf("name should not be empty")
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | green }} ",
	}

	valuePrompt := promptui.Prompt{
		Label:     "Enter the new value: ",
		Templates: templates,
		Validate:  validate,
	}

	value, err := valuePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	DataStore.UpdateValue(key, value)
	DataStore.Persist()
	fmt.Println("Object updated")
}
