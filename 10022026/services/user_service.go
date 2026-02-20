package services

import (
	"encoding/json"
	"os"
	"flat_file_cryptography/models"
)

// Read funktionalitet
func GetUsers(filename string) (models.UserList, error) {
	var list models.UserList

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(fileBytes, &list)
	if err != nil {
		return list, err
	}

	return list, nil
}

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
