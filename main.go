package main

import (
	"fmt"
	"foresightLicenseFileParserForZabbix/config"
	"foresightLicenseFileParserForZabbix/flf"
	"log"
	"os/exec"

	"flag"

	"runtime"

	"bytes"

	"time"

	"git.zabbix.com/ap/plugin-support/plugin"
	"git.zabbix.com/ap/plugin-support/plugin/container"
)
type Plugin struct {
	plugin.Base
}

var impl Plugin

var pluginConfig config.Config

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {

	var activeUsersOfAnalyticsPlatform int

	var flf flf.ForesightLicenseFile

	// Static file mode

	if (!pluginConfig.CommandMode) && (len(pluginConfig.LocalFile) > 0) {

		flf.NewForesightLicenseFile(pluginConfig.LocalFile)

		activeUsersOfAnalyticsPlatform = flf.CountActiveUsersOfFeature("AnalyticsPlatform")

	}

	// Command line mode

	if (pluginConfig.CommandMode) && (len(pluginConfig.CommandLine) > 0) {

		cmd := exec.Command(pluginConfig.CommandLine)

		buf := bytes.Buffer{}
		cmd.Stdout = &buf
		err := cmd.Run()
	
		//fmt.Println(buf.String(), err)

		// cmd.Stdout = os.Stdout

		// if err := cmd.Run(); err != nil {
		// 	fmt.Println("could not run command: ", err)
		//   }

		//  cmdOutput, err := cmd.CombinedOutput()

		if err != nil {
			log.Printf("cmd.Run() for %s failed with %s\n", pluginConfig.CommandLine, err)
		}

		// time.Sleep(20 * time.Second)


		//stderr, _ := cmd.StderrPipe()
		//stderr, _ := cmd.StdoutPipe()
		// cmd.Start()
	
		 time.Sleep(20 * time.Second)

	//	 var cmdOutput string

		// scanner := bufio.NewScanner(cmd.Stdout)
		// scanner.Split(bufio.ScanWords)
		// for scanner.Scan() {
		// 	cmdOutput = scanner.Text()			
		// }
		// cmd.Wait()

		//flf.NewForesightLicenseFileByContent([]byte(cmdOutput))

		//activeUsersOfAnalyticsPlatform = flf.CountActiveUsersOfFeature("AnalyticsPlatform")

		// ZZZZZZZZZZZZZZZZFORDEBUG PURPOSE
		activeUsersOfAnalyticsPlatform = len(buf.String())

	}

	return activeUsersOfAnalyticsPlatform, nil


}

// Mandatory functions

func init() {

	plugin.RegisterMetrics(&impl, "ForesightLicenseFileParser", "analyticsplatformonlineuserscount", "Returns an amount of AnalyticsPLatform feature online users.")

}

func main() {

	// Getting current directory as a configuration file default location
	// currentDirectory, err := os.Getwd()
	// if err != nil {
	//     fmt.Println(err)
	// }

	// When running in Zabbix Agent 2 mode there is no change to execute a plugin with command line parameters
	// In addition it is not clear which is a current directory for Zabbix Agent 2
	// That is why for Windows configuration file to be put by default in C:\Program Files\Zabbix Agent 2\config.yml

	var defaultConfigFilePath string

	if runtime.GOOS == "windows" {

		defaultConfigFilePath = "C:\\Program Files\\Zabbix Agent 2\\config.yml"

	}

	// Getting flags
	configFilePath := flag.String("configfile", defaultConfigFilePath, "Full path to configuration file")
	//configFilePath := flag.String("configfile", currentDirectory + "\\config.yml", "Full path to configuration file")

	flag.Parse()

	// Getting configuration

	var appConfig config.AppConfig

	appConfig.NewConfig(*configFilePath)

	pluginConfig = appConfig.Get()

	// Creating handler

	h, err := container.NewHandler(impl.Name())

	if err != nil {
		panic(fmt.Sprintf("failed to create plugin handler %s", err.Error()))
	}
	impl.Logger = &h
	err = h.Execute()
	if err != nil {
		panic(fmt.Sprintf("failed to execute plugin handler %s", err.Error()))
	}

}
