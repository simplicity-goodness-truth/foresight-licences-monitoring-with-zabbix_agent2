package config

import (
	"fmt"
	"foresightLicenseFileParserStandalone/file"

	"gopkg.in/yaml.v3"
)

// Interfaces

type Configuration interface {
	Print()
	Get() Config
}

// Classes definition

type AppConfig struct {
	config Config
}

type OnlineUsersCounterConfigItem struct {
    MetricName string `yaml:"metricName"`
    FeatureName string `yaml:"featureName"`
}

type Config struct {

		CommandMode bool `yaml:"commandMode"`
		CommandLine string `yaml:"commandLine"`
		LocalFile   string `yaml:"localFile"`
		OnlineUsersCounterConfig []OnlineUsersCounterConfigItem `yaml:"onlineUsersCounterConfig"`
	
}

// Constructor

func (ac *AppConfig) NewConfig(filePath string) {

	configFile := file.NewFile(filePath)

	var config Config

	err := yaml.Unmarshal(configFile.FileContent, &config)

	if err != nil {

		fmt.Printf("Error")

	}

	 ac.setConfig(config)

}

// Interfaces implementation

func (ac AppConfig) Print() {

	fmt.Print(ac.Get())
}

func (ac *AppConfig) Get() Config {

	return ac.config

}

// Private methods

func (ac *AppConfig) setConfig(config Config) {

	ac.config = config

}
