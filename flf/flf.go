package flf

import (
	"bufio"
	"fmt"
	"foresightLicenseFileParserForZabbix/file"
	"regexp"
	"strings"
)

// Constants

const featureName string = "Feature name"
const userName string = "User name"

// Interfaces

type ForesightLicenseСounter interface {
	CountActiveUsersOfFeature(featureName string) (activeUsersOfFeature int)
	GetFile() file.File
	GetFeatureUsersOnline() usersOnline
	PrintFeatureUsersOnline()
}

// Class definition

type ForesightLicenseFile struct {
	file               file.File
	featureUsersOnline usersOnline
}

// Private types

type featureUsers struct {
	name  string
	users []string
}

type usersOnline []featureUsers

// Constructor

func (f *ForesightLicenseFile) NewForesightLicenseFile(filePath string) {

	f.setFile(*file.NewFile(filePath))

}

func (f *ForesightLicenseFile) NewForesightLicenseFileByContent(content []byte) {

	file := &file.File{

		FilePath:    "dummy",
		FileContent: content,
	}

	f.setFile(*file)

}

// Interfaces implementations

func (f *ForesightLicenseFile) CountActiveUsersOfFeature(featureName string) (activeUsersOfFeature int) {

	if len(f.featureUsersOnline) == 0 {

		f.buildTree()

	}

	// f.PrintFeatureUsersOnline()

	activeUsersOfFeature = f.countOnlineUsersByFeatureName(featureName)

	return activeUsersOfFeature
}

func (f ForesightLicenseFile) GetFile() file.File {

	return f.file

}

func (f *ForesightLicenseFile) GetFeatureUsersOnline() usersOnline {

	return f.featureUsersOnline

}

func (f *ForesightLicenseFile) PrintFeatureUsersOnline() {

	for i, item := range f.GetFeatureUsersOnline() {
		fmt.Println(i, ":", item)
	}
}

// Private methods

func (f *ForesightLicenseFile) countOnlineUsersByFeatureName(featureName string) int {

	var activeUsersOfFeature int

	for _, item := range f.GetFeatureUsersOnline() {

		if item.name == featureName {

			activeUsersOfFeature = len(item.users)
		}
	}

	return activeUsersOfFeature
}

func (f *ForesightLicenseFile) setFile(file file.File) {

	f.file = file

}

func (f *ForesightLicenseFile) buildTree() {

	var featureUsers featureUsers
	var usersOnline usersOnline
	var isFeatureInAnalysis bool

	licenseFile := string(f.GetFile().FileContent)

	if len(licenseFile) > 0 {

		scanner := bufio.NewScanner(strings.NewReader(licenseFile))

		for scanner.Scan() {

			// Licence file line

			licenseFileLine := scanner.Text()

			// Picking feature name

			if strings.Contains(licenseFileLine, featureName) {

				if isFeatureInAnalysis {

					// Close a feature that was in processing on previous cycle iterations

					usersOnline = append(usersOnline, featureUsers)
					featureUsers.name = ""
					featureUsers.users = nil

				}

				// Regexp to take a feature name in brackets
				featureNamePattern := regexp.MustCompile("\"([^\"]*)\"")

				// Taking all matching groups
				featureNameMatches := featureNamePattern.FindAllStringSubmatch(licenseFileLine, -1)

				// A feature name without quotes will be stored in a first matching group
				featureUsers.name = featureNameMatches[0][1]

				//	fmt.Printf("UNAME==>%s\n",featureUsers.name )

				isFeatureInAnalysis = true
			}

			if strings.Contains(licenseFileLine, userName) && isFeatureInAnalysis {

				// Regexp to take a user name after colon
				userNamePattern := regexp.MustCompile(":\\s(.+)")

				// Taking all matching groups
				userNameMatches := userNamePattern.FindAllStringSubmatch(licenseFileLine, -1)

				// A user name without colon will be stored in a first matching group
				featureUsers.users = append(featureUsers.users, userNameMatches[0][1])

			}

		}

		if isFeatureInAnalysis {

			// Close a feature that was in processing on a last cycle iteration

			usersOnline = append(usersOnline, featureUsers)

		}

		f.featureUsersOnline = usersOnline

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		}

	}

}
