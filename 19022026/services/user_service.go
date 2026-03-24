package services

import (
	"encoding/json"
	"fmt"
	"go-rest-api/internal/crypto"
	"go-rest-api/models"
	"os"
)

// Read funktionalitet
func GetUsers() (models.UserList, error) {
	var list models.UserList

	fileBytes, err := os.ReadFile("data.json")
	if err != nil {
		return list, err
	}

	// Converting from raw bytes to user objects and adds to UserList
	err = json.Unmarshal(fileBytes, &list)
	if err != nil {
		return list, err
	}

	// Decryption of user objects
	for i := range list.Users {
		err := DecryptData(&list.Users[i])
		if err != nil {
			return models.UserList{}, err
		}
	}

	return list, nil
}

// Create funktionalitet
func CreateUser(newUser models.User) error {
	users, err := GetUsers()
	if err != nil {
		return err
	}

	users.Users = append(users.Users, newUser)
	for i := range users.Users {
		if err := EncryptData(&users.Users[i]); err != nil {
			return err
		}
	}

	// List to raw bytes
	fileBytes, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("data.json", fileBytes, 0644)
}

// Update funktionalitet
func UpdateUser(updatedUser models.User) bool {
	list, err := GetUsers()
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	for i, user := range list.Users {
		if user.ID == updatedUser.ID {
			list.Users[i] = updatedUser

			for j := range list.Users {
				if err := EncryptData(&list.Users[j]); err != nil {
					return false
				}
			}

			fileBytes, _ := json.MarshalIndent(list, "", "    ")
			err := os.WriteFile("data.json", fileBytes, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return false
			}

			return true
		}
	}
	return false
}

// Delete funktionalitet
func DeleteUser(id int) bool {
	list, err := GetUsers()
	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	for i, user := range list.Users {
		if user.ID == id {
			list.Users = append(list.Users[:i], list.Users[i+1:]...)

			// Encrypting the data before saving to file
			for j := range list.Users {
				if err := EncryptData(&list.Users[i]); err != nil {
					return false
				}
			}

			// Saves the encrypted data
			fileBytes, _ := json.MarshalIndent(list, "", "    ")
			err := os.WriteFile("data.json", fileBytes, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
			}

			return true
		}
	}

	return false
}

func EncryptData(u *models.User) error {
	var err error
	u.FirstName, err = crypto.Encrypt([]byte(u.FirstName))
	u.LastName, err = crypto.Encrypt([]byte(u.LastName))
	u.Email, err = crypto.Encrypt([]byte(u.Email))
	u.Address, err = crypto.Encrypt([]byte(u.Address))

	// Hash passwordet hvis det ikke allerede er gjort (fx længde-tjek)
	if len(u.Password) < 30 {
		hashed, err := crypto.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = []byte(hashed)
	}
	return err
}

// Decrypt gør dataen læsbar igen
func DecryptData(u *models.User) error {
	firstName, err := crypto.Decrypt(u.FirstName)
	if err == nil {
		u.FirstName = string(firstName)
	}
	lastName, err := crypto.Decrypt(u.LastName)
	if err == nil {
		u.LastName = string(lastName)
	}
	email, err := crypto.Decrypt(u.Email)
	if err == nil {
		u.Email = string(email)
	}
	address, err := crypto.Decrypt(u.Address)
	if err == nil {
		u.Address = string(address)
	}
	return nil // Du kan vælge at returnere fejlen hvis nødvendigt
}
