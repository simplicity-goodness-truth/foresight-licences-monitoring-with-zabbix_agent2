package file

import (
	"foresightLicenseFileParserForZabbix/check"
	"os"
)

// Class definition

type File struct {
	FilePath    string
	FileContent []byte
}

// Constructor

func NewFile(filePath string) *File {

	return &File{

		FilePath:    filePath,
		FileContent: readFile(filePath),
	}

}

// Private methods

func readFile(filePath string) []byte {

	var fileContent []byte

	fileContent, err := os.ReadFile(filePath)

	check.Check(err)

	return fileContent

}
