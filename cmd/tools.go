package cmd

import (
	"fmt"
	"strings"
)

func listSelector(list []string, actionSuccess func(string), actionFailure func()) {
	for _, item := range list {
		fmt.Println(item)
	}
	fmt.Print("Choose here: ")
	var choice string
	fmt.Scanln(&choice)
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
