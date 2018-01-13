package cmd

import (
	"fmt"
	"strings"

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
			if isUnique, modelName := isUniquePossibility(args[0]); isUnique {
				setModel(modelName)
			} else {
				fmt.Println("You didn't pick a correct model, please retry")
				selectModel()
			}
		}
	},
}

func init() {
	validModels = getModels()
	rootCmd.AddCommand(mCmd)
}

func setModel(name string) {
	modelExists(name)
	config.ReadInConfig()
	config.Set(cfgModel, name)
	config.WriteConfig()
	fmt.Println("Model successfully set to", name)
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
	for _, model := range validModels {
		fmt.Println(model)
	}
	fmt.Print("Choose here: ")
	var choice string
	fmt.Scanln(&choice)
	if isUnique, modelName := isUniquePossibility(choice); isUnique {
		setModel(modelName)
	} else {
		fmt.Println("You didn't pick a correct model, please retry")
		selectModel()
	}
}

func isUniquePossibility(name string) (bool, string) {
	hasOccured := false
	modelName := ""
	for _, model := range validModels {
		if strings.HasPrefix(model, name) {
			if hasOccured {
				return false, ""
			}
			hasOccured = true
			modelName = model
		}
	}
	return hasOccured, modelName
}

func modelExists(name string) bool {
	for _, model := range validModels {
		if model == name {
			return true
		}
	}
	return false
}
