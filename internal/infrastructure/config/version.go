package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadVersion() string {
	file, err := os.Open("version")
	if err != nil {
		log.Printf("Error reading version file: %v", err)
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning version file: %v", err)
	}
	return "unknown"
}
