package helpers

import (
	"log"
	"os"
)

func LogError(msg string, err error) {
	if err != nil {
		log.Fatal(msg+"\n", err)
		os.Exit(2)
	}
}
