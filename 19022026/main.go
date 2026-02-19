package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Read funktionalitet
func loadUsers() (UserList, error) {
	var list UserList

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
func saveUsers(list UserList) error {
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
func updateUser(list *UserList, updatedUser User) bool {
	for i, user := range list.Users {
		if user.ID == updatedUser.ID {
			list.Users[i] = updatedUser
			return true
		}
	}

	return false
}

// Delete funktionalitet
func deleteUser(list *UserList, id int) bool {
	for i, user := range list.Users {
		if user.ID == id {
			list.Users = append(list.Users[:i], list.Users[i + 1:]...)
			return true
		}
	}

	return false
}

func main() {
	users, err := loadUsers()
	if err != nil {
		fmt.Println("Fejl ved indlæsning:", err)
		return
	}

	fmt.Printf("Indlæste %d brugere fra filen.\n", len(users.Users))
	fmt.Println("Første bruger er: ", users.Users[0].FirstName)

	nyInfo := User{
		ID: 			1,
		FirstName:		"John (Opdateret)",
		LastName:		"Doe",
		Address: 		"somewhere",
		StreetNumber:	69,
		Password:		"test",
		Enabled:		true,
	}

	if success := updateUser(&users, nyInfo); success {
		fmt.Println("Bruger 1 blev opdateret!")
	}

	if success := deleteUser(&users, 2); success {
		fmt.Println("Bruger 2 blev slettet!")
	}

	err = saveUsers(users)
	if err != nil {
		fmt.Println("Fejl ved gemning: ", err)
	} else {
		fmt.Println("Filen er nu opdateret med ændringerne!")
	}

	fmt.Println("Sidste bruger er: ", users.Users[len(users.Users) - 1].FirstName)

}
