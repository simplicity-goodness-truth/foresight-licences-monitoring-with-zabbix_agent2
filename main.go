package main

import (
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

	// appConfig.Print()

	var flf flf.ForesightLicenseFile

	// Static file mode

	if (!appConfig.Get().CommandMode) && (len(appConfig.Get().LocalFile) > 0) {


		flf.NewForesightLicenseFile(appConfig.Get().LocalFile)

		activeUsersOfAnalyticsPlatform := flf.CountActiveUsersOfFeature("AnalyticsPlatform")

		fmt.Printf("Users of AnalyticsPlatform: %d \n", activeUsersOfAnalyticsPlatform)

		activeUsersOfPP_Adm := flf.CountActiveUsersOfFeature("PP_Adm")

		fmt.Printf("Users of PP_Adm: %d \n", activeUsersOfPP_Adm)

	}

	// Command line mode

	if (appConfig.Get().CommandMode) && (len(appConfig.Get().CommandLine) > 0) {

		cmd := exec.Command(appConfig.Get().CommandLine)

		cmdOutput, err := cmd.CombinedOutput()
		
		if err != nil {
			log.Printf("cmd.Run() for %s failed with %s\n", appConfig.Get().CommandLine, err)
		}

		flf.NewForesightLicenseFileByContent(cmdOutput)

		activeUsersOfAnalyticsPlatform := flf.CountActiveUsersOfFeature("AnalyticsPlatform")

		fmt.Printf("Users of AnalyticsPlatform: %d \n", activeUsersOfAnalyticsPlatform)

	}

}
