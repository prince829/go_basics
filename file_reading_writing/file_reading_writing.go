package file_reading_writing

import (
	"bufio"
	"fmt"
	"os"
)

func FileReadingWritingFunc() {
	//write File
	data := []byte("Hello Prince\n How are you\n and seems your are great")
	err := os.WriteFile("output.txt", data, 0644) // prem 0644 means owner can read/write, group can read only and others can read only
	if err != nil {
		fmt.Println(err)
		return
	}
	//Read Full File
	readInfo, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OutputFile Info: ", string(readInfo))

	//Line by Line Reading with bufio
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("\nLine-by-line output:")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}
	//Check For Error
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner Error:", err)
	}
}
