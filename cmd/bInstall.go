package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
