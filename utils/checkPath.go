package utils

import "os"

type filePath = string

func FileExists(file filePath) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true // File exists
	}
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return false // Error occurred while checking file existence
}