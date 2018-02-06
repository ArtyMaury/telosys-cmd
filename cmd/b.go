package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var availableBundles []string

// bCmd represents the b command
var bCmd = &cobra.Command{
	Use:     "b",
	Aliases: []string{"bundle"},
	Short:   "Set current bundle",
	Long:    "Set the bundle used in the telosys project",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			selectBundle()
		} else {
			setBundle(args[0])
		}
	},
}

func init() {
	availableBundles = getBundles()
	rootCmd.AddCommand(bCmd)
}

func setBundle(name string) {
	//Check if the name passed matches a unique bundle name
	if isUnique, bundleName := isUniquePossibility(name, availableBundles); isUnique {
		setConfValue(cfgBundle, bundleName)
		fmt.Println("Bundle successfully set to", bundleName)
	} else {
		fmt.Println("Bundle doesn't exist")
	}
}

// Gets the installed bundles
func getBundles() []string {
	bundles := getMatching("templates/*")
	newList := []string{}
	for _, bundle := range bundles {
		newList = append(newList, rmPath(bundle))
	}
	return newList
}

// Allows the user to select a bundle in the list of available bundles
func selectBundle() {
	fmt.Println("Here are the available bundles:")
	listSelector(availableBundles, setBundle, func() {
		fmt.Println("You didn't pick a correct bundle, please retry")
		selectBundle()
	})
}
