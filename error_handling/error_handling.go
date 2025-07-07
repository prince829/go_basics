package error_handling

import (
	"errors"
	"fmt"
)

func getAge(name string) (int, error) {
	switch name {
	case "prince":
		return 25, nil
	case "":
		return 0, errors.New("name cannot be empty")
	default:
		return 0, fmt.Errorf("user %q not found", name)
	}
}

func ErrorHandling() {
	age, err := getAge("alex")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("AGE:", age)
	}

}
