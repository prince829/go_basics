package buffered_channel

import "fmt"

/*
A buffered channel lets you send values without needing immediate receiving — up to a limit.
BUFFERED: Do not requires receiver imedetaily && hold values for limited
UNBUFFERED: Requires receiver && Could not hold values
*/
func BufferedChannel() {
	//Syntax
	//	ch:=make(chan int 3)//You can now send up to 3 values before it blocks.
	// ch := make(chan int, 2) // buffer size 2
	// ch <- 2
	// ch <- 3
	// // ch <- 4           //it will block(buffer full)
	// fmt.Println(<-ch) // 10
	// fmt.Println(<-ch) // 20
	ch := make(chan int, 3)

	// send values
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch) // ✅ This tells range to stop when done
	// receive values in loop
	for value := range ch {
		fmt.Println(value)
	}
}
