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
	keys := DataStore.GetKeys()
	if len(keys) == 0 {
		fmt.Println("Data store empty.")
		return
	}

	selectTemplates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "{{ . | red | cyan }}",
	}

	selectPrompt := promptui.Select{
		Label:     "? Choose a key to update:",
		Items:     keys,
		Templates: selectTemplates,
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
		Label:     "? Enter the new value: ",
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
	fmt.Printf("'%s' updated.\n", key)
}
