package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var cfgHomedir = "homeDir"
var cfgModel = "model"
var cfgBundle = "bundle"
var cfgGithub = "github"
var cfgList = []string{cfgHomedir, cfgModel, cfgBundle, cfgGithub}

// envCmd represents the conf command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Display current environment",
	Long:  "Display current environment",
	Run: func(cmd *cobra.Command, args []string) {
		if mapConf, err := getConf(); err == nil {
			for key, value := range mapConf {
				fmt.Println(key, ":", value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func getConf() (map[string]string, error) {
	confMap := make(map[string]string)
	err := config.ReadInConfig()
	if err == nil {
		for _, conf := range cfgList {
			confMap[conf] = config.GetString(conf)
		}
	}
	return confMap, err
}

func checkConfig() error {
	mapConf, err := getConf()
	if err == nil {
		if mapConf[cfgModel] != "" {
			if modelFolder := getMatching(mapConf[cfgModel] + "_model"); len(modelFolder) < 0 {
				err = errors.New("Model non disponible")
			}
			if modelFile := getMatching(mapConf[cfgModel] + ".model"); len(modelFile) < 0 {
				err = errors.New("Model non disponible")
			}
		}
	}
	return err
}

func setConfValue(entry, value string) error {
	if err := config.ReadInConfig(); err == nil {
		config.Set(entry, value)
		config.WriteConfig()
	} else {
		return err
	}
	return nil
}

func getConfValue(entry string) string {
	if err := config.ReadInConfig(); err == nil {
		return config.GetString(entry)
	}
	return ""
}
