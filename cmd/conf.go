package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "Display current config",
	Long:  "Display current config",
	Run: func(cmd *cobra.Command, args []string) {
		if mapConf, err := getConf(); err == nil {
			for key, value := range mapConf {
				fmt.Println(key, ":", value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}

func getConf() (map[string]string, error) {
	confMap := make(map[string]string)
	err := config.ReadInConfig()
	if err == nil {
		for _, conf := range []string{cfgHomedir, cfgBundle, cfgModel} {
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
