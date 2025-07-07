package mini_project

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

func SaveUsersToFile(users []User, filename string) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadUsersFromFile(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	return users, err
}

func MiniProject() {
	// Sample data
	users := []User{
		{"Prince", "prince@example.com", 25, "New York"},
		{"Arjun", "arjun@example.com", 30, "Delhi"},
	}

	filename := "users.json"

	// Save to file
	err := SaveUsersToFile(users, filename)
	if err != nil {
		fmt.Println("Error saving users:", err)
		return
	}
	fmt.Println("✅ Users saved to file!")

	// Load from file
	loadedUsers, err := LoadUsersFromFile(filename)
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	// Print loaded data
	fmt.Println("\n📂 Loaded Users:")
	for _, user := range loadedUsers {
		fmt.Printf("- %s (%d) from %s\n", user.Name, user.Age, user.Location)
	}
}
