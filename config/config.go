package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type configFile struct {
	Discord struct {
		Token string `yaml:"token"`
	} `yaml:"discord"`
	Log struct {
		Lvl      string `yaml:"lvl"`
		ToStdout bool   `yaml:"to-stdout"`
	} `yaml:"log"`
}

var Config configFile

func Init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configFileFound, conf := reader(dir + "/config.yaml")
	if !configFileFound {
		panic("Cant find config file")
	}

	Config = conf
}

// Parses the config.yaml file to the Config variable
// Panics on uncaught errors
func reader(configPath string) (bool, configFile) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil && os.IsNotExist(err) {
		return false, configFile{}
	} else if err != nil {
		panic(err)
	}

	var ret configFile
	err = yaml.Unmarshal(yamlFile, &ret)
	if err != nil {
		panic(err)
	}

	return true, ret
}
