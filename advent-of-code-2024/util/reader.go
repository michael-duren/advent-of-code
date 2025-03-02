package util

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

func ReadFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	return string(content)
}
