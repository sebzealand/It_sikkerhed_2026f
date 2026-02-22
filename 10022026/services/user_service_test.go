package services

import (
	"encoding/json"
	"flat_file_cryptography/models"
	"os"
	"reflect"
	"testing"
)

// TEST AF GET FUNKTIONALITET

func TestGetUsers(t *testing.T) {
	// Arrange
	testFile := "../data/test_data.json"
	defer os.Remove(testFile)

	mockData := `{
		"users": [
			{"person_id": 1, "first_name": "John"},
			{"person_id": 2, "first_name": "Jane"}
		]
	}`
	os.WriteFile(testFile, []byte(mockData), 0644)

	// Act
	got, err := GetUsers(testFile)

	// Assert
	if err != nil {
		t.Fatalf("GetUsers fejlede uventet: %v", err)
	}

	wantCount := 2
	gotCount := len(got.Users)

	if gotCount != wantCount {
		t.Errorf("Antal brugere forkert: fik %d, men ville have %d", gotCount, wantCount)
	}

	if got.Users[0].FirstName != "John" || got.Users[1].FirstName != "Jane" {
		t.Error("Navnene i de indlæste objekter matcher ikke mock-data")
	}
}

func TestGetUsers_FileNotFound(t *testing.T) {
	_, err := GetUsers("")

	if err == nil {
		t.Error("Forventede en fejl for en ikke-eksisterende fil, men fik nil")
	}
}

// TEST AF CREATE FUNKTIONALITET
func TestCreateUser(t *testing.T) {
	// Arrange
	testFile := "../data/test_data.json"
	defer os.Remove(testFile)

	want := models.UserList{
		Users: []models.User{
			{
				ID:           1,
				FirstName:    "Test",
				LastName:     "Bruger",
				Address:      "Testvej",
				StreetNumber: 1,
				Password:     []byte("Test123"),
				Enabled:      true,
			},
		},
	}

	// Act
	err := CreateUser(testFile, want)
	if err != nil {
		t.Fatalf("Kan ikke gemme bruger: %v", err)
	}

	// Assert
	data, _ := os.ReadFile(testFile)
	var got models.UserList
	json.Unmarshal(data, &got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("De to objekter matcher ikke: \ngot: %+v \nwant: %+v", got, want)
	}
}

func TestCreateUser_InvalidPath(t *testing.T) {
	// Arrange
	invalidFile := ""
	defer os.Remove(invalidFile)
	data := models.UserList{Users: []models.User{{ID: 1}}}

	// Act
	err := CreateUser(invalidFile, data)

	// Assert
	if err == nil {
		t.Error("Forventede en fejl pga. ugyldig sti, men fik ingen.")
	}
}

func TestCreateUser_ReadOnlyFile(t *testing.T) {
	// Arrange
	fileName := "readonly_test.json"
	os.WriteFile(fileName, []byte("start"), 0444)
	defer os.Remove(fileName)

	data := models.UserList{Users: []models.User{{ID: 1}}}

	// Act
	err := CreateUser(fileName, data)

	// Assert
	if err == nil {
		t.Error("Forventede en fejl ved skrivning til en skriebeskyttet fil, men fik inge")
	}
}
