package services

import (
	"encoding/json"
	"os"
	"go-rest-api/models"
)


// Read funktionalitet
func LoadUsers() (models.UserList, error) {
	var list models.UserList

	fileBytes, err := os.ReadFile("data.json")
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
func SaveUsers(list models.UserList) error {
	fileBytes, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Update funktionalitet
func UpdateUser(list *models.UserList, updatedUser models.User) bool {
	for i, user := range list.Users {
		if user.ID == updatedUser.ID {
			list.Users[i] = updatedUser
			return true
		}
	}

	return false
}

// Delete funktionalitet
func DeleteUser(list *models.UserList, id int) bool {
	for i, user := range list.Users {
		if user.ID == id {
			list.Users = append(list.Users[:i], list.Users[i + 1:]...)
			return true
		}
	}

	return false
}
