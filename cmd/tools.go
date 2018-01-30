package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
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
		if strings.Contains(item, prefix) {
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

type mapObject map[string]interface{}

func getHttpJsonMap(url string) []map[string]interface{} {
	var jsonMaps []map[string]interface{}
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		dec := json.NewDecoder(resp.Body)
		_, error1 := dec.Token()
		if error1 != nil {
			log.Fatal(err)
		}
		for dec.More() {
			var jsonMap mapObject
			err := dec.Decode(&jsonMap)
			if err != nil {
				log.Fatal(err)
			}
			jsonMaps = append(jsonMaps, jsonMap)
		}
		_, error1 = dec.Token()
		if error1 != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("Error getting json from url")
	}
	return jsonMaps
}

func getHttpJsonValues(url string, keys ...string) []map[string]interface{} {
	jsonMaps := getHttpJsonMap(url)
	var newMaps []map[string]interface{}
	for _, jsonMap := range jsonMaps {
		newMap := make(map[string]interface{})
		for _, key := range keys {
			keyCuts := strings.Split(key, ".")
			for _, path := range keyCuts[:len(keyCuts)-1] {
				jsonMap = jsonMap[path].(map[string]interface{})
			}
			newMap[key] = jsonMap[keyCuts[len(keyCuts)-1]]
		}
		newMaps = append(newMaps, newMap)
	}
	return newMaps
}
