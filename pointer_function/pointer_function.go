package pointer_function

import "fmt"

func updateValue(x *int) {
	*x = *x + 10
}
func PointerFunction() {
	var a int = 10
	var p *int = &a
	fmt.Println(a)
	fmt.Println(p)  //print memory address
	fmt.Println(*p) //print 10

	//Modify value with pointer
	val := 5
	updateValue(&val) //passed the address of val
	fmt.Println(val, "valllll")
}

// Struct+pointer
type Book struct {
	name   string
	author string
	year   int
}

func (b Book) Info() string {
	return fmt.Sprintf("%s by %s (%d)", b.name, b.author, b.year)
}

// Go automatically takes the address of book for Promote()
func (b *Book) Promote() {
	b.year += 10
}

/*
*

	func  Promote(b *Book) {
		b.year += 10
	}
*/
func DualPointer() {
	book := Book{
		name:   "H1",
		author: "a1",
		year:   2000,
	}

	book.Promote()
	// Promote(&book) //For commented promote method
	fmt.Println(book.Info())
}

/**
 Why use a pointer receiver (*Book) instead of value receiver (Book)?
Because with a value receiver, any modification happens on a copy — so the original stays unchanged.

With a pointer receiver, your method can modify the original struct.
*/
