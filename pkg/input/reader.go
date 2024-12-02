package input

import (
	"log"
	"os"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file %s due to: %v", path, err)
		panic(err)
	}
	return string(content)
}
