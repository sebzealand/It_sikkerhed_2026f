package models

type User struct {
	ID           int    `json:"person_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	StreetNumber int    `json:"street_number"`
	Password     []byte `json:"password"`
	Enabled      bool   `json:"enabled"`
}

// Creates a list of user objects
type UserList struct {
	Users []User `json:"users"`
}
