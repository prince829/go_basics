package struct_function

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func (b User) Info() string {
	return fmt.Sprintf("%s has %d (%s)", b.name, b.age, b.email)
}

func StructFunc() {
	/**
		Structs let you group related values into a single composite type.
	    They’re like JavaScript objects or classes without methods.
	*/
	//A struct is a collection of fields.

	users := []User{
		{
			name:  "Prince",
			age:   25,
			email: "prince@yopmail.com"},
	}
	/**
	For single
	user1:=User{
	       name:  "Prince",
			age:   25,
			email: "prince@yopmail.com"
	      }

	// fmt.Println(user1.name, "user name")
	// user1.age = 26 // modify the age
	*/
	for index, user := range users {
		fmt.Println(user.name, "has age", user.age, "at position", index)
		fmt.Println(user.Info(), "Info")
	}
}
