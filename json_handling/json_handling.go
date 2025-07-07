package json_handling

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func JsonHandling() {
	//Marshal
	/*
		json.Marshal() takes a Go struct and returns a JSON-encoded []byte
		json:"name" tells Go to use "name" as the key in JSON (not Name)
		If a field is capitalized (exported), it's included. If lowercase, it's ignored.
	*/
	u := User{Name: "Prince", Email: "p@yopmail.com", Age: 25}
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Formate Json:", string(jsonData))
	//unMarshal
	/*
			| Function            | Description                               |
		| ------------------- | ----------------------------------------- |
		| `json.Unmarshal`    | Converts `[]byte` JSON → Go struct        |
		| `[]byte(jsonInput)` | Input string must be in `[]byte` form     |
		| `var user User`     | Target struct must be a pointer (`&user`) |

	*/
	jsonString := `{"name": "Bob","email": "b@yopmail.com","age": 30}`
	var u1 User
	err1 := json.Unmarshal([]byte(jsonString), &u1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Printf("Parsed struct: %+v\n", u1)
	/**
	In Go, %+v is a formatting verb used with the fmt.Printf function.
	It's particularly useful when working with structs.
	 When used with a struct, %+v will print the struct's field names along with their corresponding values
	*/

}
