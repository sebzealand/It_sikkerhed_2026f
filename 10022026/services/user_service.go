package services

import (
	"encoding/json"
	"flat_file_cryptography/models"
	"os"
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

	for i := range list.Users {
		decryptedFirst, err := Decrypt(list.Users[i].FirstName)
		if err == nil {
			list.Users[i].FirstName = string(decryptedFirst)
		}

		decryptedLast, err := Decrypt(list.Users[i].LastName)
		if err == nil {
			list.Users[i].LastName = string(decryptedLast)
		}

		decryptedAddress, err := Decrypt(list.Users[i].Address)
		if err == nil {
			list.Users[i].Address = string(decryptedAddress)
		}
	}

	return list, nil
}

// Create funktionalitet
func CreateUser(filename string, list models.UserList) error {
	encryptedList := list

	for i := range encryptedList.Users {
		encryptedList.Users[i].FirstName, _ = Encrypt([]byte(encryptedList.Users[i].FirstName))
		encryptedList.Users[i].LastName, _ = Encrypt([]byte(encryptedList.Users[i].LastName))
		encryptedList.Users[i].Address, _ = Encrypt([]byte(encryptedList.Users[i].Address))

		if len(encryptedList.Users[i].Password) < 30 {
			hashedPw, _ := HashPassword(encryptedList.Users[i].Password)
			encryptedList.Users[i].Password = []byte(hashedPw)
		}
	}

	fileBytes, err := json.MarshalIndent(encryptedList, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
