package handler

import (
	"devboxctl/utils"
	"log"
)

func Fatal(message string, err error) {
	log.Fatalf("%s \n\n %s\n", utils.Alert.Sprint(message + "\n"), err)
}