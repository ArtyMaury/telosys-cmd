package cmd

import (
	"errors"
	"fmt"
	"strings"
)

func listSelector(list []string, actionSuccess func(string), actionFailure func()) {
	for _, item := range list {
		fmt.Println(item)
	}
	choice := askUser("Choose here: ")
	if isUnique, name := isUniquePossibility(choice, list); isUnique {
		actionSuccess(name)
	} else {
		actionFailure()
	}
}

func isUniquePossibility(prefix string, list []string) (bool, string) {
	hasOccured := false
	finalName := ""
	for _, item := range list {
		if strings.HasPrefix(item, prefix) {
			if hasOccured {
				return false, ""
			}
			hasOccured = true
			finalName = item
		}
	}
	return hasOccured, finalName
}

func contains(item string, list []string) bool {
	for _, elt := range list {
		if item == elt {
			return true
		}
	}
	return false
}

func askUser(question string) string {
	fmt.Print(question)
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func askConfirmation() error {
	skip, _ := rootCmd.PersistentFlags().GetBool("y")
	if !skip {
		choice := askUser("Are you sure? [Y/n] :")
		if strings.HasPrefix(choice, "n") {
			return errors.New("Cancelled action")
		}
	}
	return nil
}
