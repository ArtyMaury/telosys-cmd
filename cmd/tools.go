package cmd

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Ask the user to choose in a list, and redirecting to the right function in case of a success or a failure
// It matches the user input with the possibilities
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

// Checks if the input matches only one possibility and returns the success boolean and full name
func isUniquePossibility(input string, list []string) (bool, string) {
	hasOccured := false
	finalName := ""
	for _, item := range list {
		if strings.Contains(item, input) {
			if hasOccured {
				return false, ""
			}
			hasOccured = true
			finalName = item
		}
	}
	return hasOccured, finalName
}

// Tests if an item is included in a list
func contains(item string, list []string) bool {
	for _, elt := range list {
		if item == elt {
			return true
		}
	}
	return false
}

// Asks a question to the user annd returns the input
func askUser(question string) string {
	fmt.Print(question)
	var choice string
	fmt.Scanln(&choice)
	return choice
}

// Ask if user confirms the action
// Can be overriden by the -y flag
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

// Returns a list of maps matching a json returned by a url request
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
			var jsonMap map[string]interface{}
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

// Returns a list of maps matching a json returned by a url request filtered by keys
// Keys can define a path like in javascript ex: "parent.child.name"
func getHttpJsonValues(url string, keys ...string) []map[string]interface{} {
	jsonMaps := getHttpJsonMap(url)
	var newMaps []map[string]interface{}
	for _, jsonMap := range jsonMaps {
		newMap := make(map[string]interface{})
		for _, key := range keys {
			keyCuts := strings.Split(key, ".")
			jsonMapSub := jsonMap
			for _, path := range keyCuts[:len(keyCuts)-1] {
				jsonMapSub = jsonMapSub[path].(map[string]interface{})
			}
			newMap[key] = jsonMapSub[keyCuts[len(keyCuts)-1]]
		}
		newMaps = append(newMaps, newMap)
	}
	return newMaps
}

func printList(list []string) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// Unzip function, credit https://golangcode.com/unzip-files-in-go/
func unzip(src, dest string) error {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
