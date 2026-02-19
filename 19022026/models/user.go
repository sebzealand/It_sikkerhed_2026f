package models

type User struct {
	ID				int		`json:"person_id"`
	FirstName		string	`json:"first_name"`
	LastName		string	`json:"last_name"`
	Address			string	`json:"address"`
	StreetNumber 	int		`json:"street_number"`
	Password		string	`json:"password"`
	Enabled			bool	`json:"enabled"`
}

type UserList struct {
	Users []User `json:"users"`
}


