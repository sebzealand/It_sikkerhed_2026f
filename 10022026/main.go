package main

import (
	"flat_file_cryptography/services"
	"fmt"
)

func main() {
	filePath := "data/flat_file_db.json"
	users, err := services.GetUsers(filePath)
	if err != nil {

	}
	/*
		var newUser = models.User{
			ID:           10,
			FirstName:    "Sebastian",
			LastName:     "Brinch",
			Address:      "hovedvejen",
			StreetNumber: 2,
			Password:     []byte("Test123"),
			Enabled:      true,
		}

		users.Users = append(users.Users, newUser)

		services.CreateUser(filePath, users)

		users1, err := services.GetUsers(filePath)
		print("Result: ", users1.Users[0].FirstName)
	*/
	check := services.CheckPasswordHash([]byte("Test123"), users.Users[0].Password)

	if check {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Password is incorrect")
	}
}
