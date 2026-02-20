package services

import (
	"encoding/json"
	"os"
	"flat_file_cryptography/models"
)

// Create funktionalitet
func CreateUser(filename string, list models.UserList) error {
	fileBytes, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
