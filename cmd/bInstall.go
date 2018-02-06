package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var githubMasterZipPattern = "https://github.com/${USER}/${REPO}/archive/master.zip"

// bInstallCmd represents the b install command
var bInstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Install a bundle from the github repository",
	Long:    "Install a bundle from the github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			selectGithubRepo()
		} else {
			getGithubRepoList()
		}
	},
}

func init() {
	bCmd.AddCommand(bInstallCmd)
}

func selectGithubRepo() {
	listRepo := getGithubRepoList()
	listSelector(listRepo, installGithubRepo, func() {
		fmt.Println("Aucun repo github ne correspond")
		selectGithubRepo()
	})
}

// Downloads and unzip the github repo in the templates folder
func installGithubRepo(repo string) {
	newDir("templates")
	if err := downloadGithubRepo(repo, toAbsPath("templates", repo+".zip")); err != nil {
		panic(err)
	}
	if err := unzip(toAbsPath("templates", repo+".zip"), toAbsPath("templates")); err == nil {
		fmt.Println("Installation completed")
	} else {
		fmt.Println(err)
	}
	os.Remove(toAbsPath("templates", repo+".zip"))
}

// Downloads the github repo zip in the given folder
func downloadGithubRepo(repo, filepath string) error {
	url := "https://github.com/" + getGithubUser(getConfValue(cfgGithub)) + "/" + repo + "/archive/master.zip"

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
