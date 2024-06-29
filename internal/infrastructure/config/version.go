package config

import (
	"io/ioutil"
	"log"
	"os"
)

func LoadVersion() string {
	versionFile := "version"
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		log.Fatalf("Version file does not exist: %v", err)
	}

	version, err := ioutil.ReadFile(versionFile)
	if err != nil {
		log.Fatalf("Failed to read version file: %v", err)
	}

	return string(version)
}
