package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var availableModels []string

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
	availableModels = getModels()
	rootCmd.AddCommand(mCmd)
}

func setModel(name string) {
	availableModels = getModels()
	if isUnique, modelName := isUniquePossibility(name, availableModels); isUnique {
		setConfValue(cfgModel, modelName)
		fmt.Println("Model successfully set to", name)
	} else {
		fmt.Println("Model doesn't exist")
	}
}

func getModels() []string {
	//TODO init bug

	models := getMatching("*.model")
	newList := []string{}
	for _, model := range models {
		modelName := rmExt(model)
		modelFolder := getMatching(modelName + "_model")
		if len(modelFolder) > 0 {
			newList = append(newList, modelName)
		} else {
			fmt.Println("Model", modelName, "is missing the model folder")
		}
	}
	return newList
}

func selectModel() {
	fmt.Println("Here are the available models:")
	listSelector(availableModels, setModel, func() {
		fmt.Println("You didn't pick a correct model, please retry")
		selectModel()
	})
}
