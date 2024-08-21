package main

import (
	"bytes"
	"flag"
	"fmt"
	"foresightLicenseFileParserStandalone/config"
	"foresightLicenseFileParserStandalone/flf"
	"log"
	"os"
	"os/exec"
)

// Mandatory functions

func init() {

}

func main() {

	// Getting current directory as a configuration file default location

	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Getting flags

	configFilePath := flag.String("configfile", currentDirectory+"\\config.yml", "Full path to configuration file")

	flag.Parse()

	// Getting configuration

	var appConfig config.AppConfig

	appConfig.NewConfig(*configFilePath)

	var flf flf.ForesightLicenseFile

	// Static file mode

	if (!appConfig.Get().CommandMode) && (len(appConfig.Get().LocalFile) > 0) {

		flf.NewForesightLicenseFile(appConfig.Get().LocalFile)

		for i, _ := range appConfig.Get().OnlineUsersCounterConfig {			

			fmt.Printf("metric: %s value %d \n", appConfig.Get().OnlineUsersCounterConfig[i].MetricName,
				flf.CountActiveUsersOfFeature(appConfig.Get().OnlineUsersCounterConfig[i].FeatureName) )
	
		}
	}

	// Command line mode

	if (appConfig.Get().CommandMode) && (len(appConfig.Get().CommandLine) > 0) {

		fmt.Printf("Executing command mode\n")

		var out bytes.Buffer
		var stderr bytes.Buffer

		cmd := exec.Command(appConfig.Get().CommandLine)

		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()

	}

}
