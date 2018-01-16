package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validModels []string

// nmCmd represents the nm command
var mCmd = &cobra.Command{
	Use:     "m",
	Aliases: []string{"model"},
	Short:   "Set current model",
	Long:    "Set the model used in the telosys project",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			selectModel()
		} else {
			setModel(args[0])
		}
	},
}

func init() {
	validModels = getModels()
	rootCmd.AddCommand(mCmd)
}

func setModel(name string) {
	validModels = getModels()
	if isUnique, modelName := isUniquePossibility(name, validModels); isUnique {
		config.ReadInConfig()
		config.Set(cfgModel, modelName)
		config.WriteConfig()
		fmt.Println("Model successfully set to", name)
	} else {
		fmt.Println("Model doesn't exist")
	}
}

func getModels() []string {
	models := getMatching("*.model")
	newList := []string{}
	for _, model := range models {
		newList = append(newList, rmExt(model))
	}
	return newList
}

func selectModel() {
	fmt.Println("Here are the available models:")
	listSelector(validModels, setModel, func() {
		fmt.Println("You didn't pick a correct model, please retry")
		selectModel()
	})
}
