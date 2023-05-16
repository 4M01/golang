//Go supports pointers, allowing you to pass references to values and records within your program.

package main

import "fmt"

// passed to it by value.
func zeroval(ival int) {
	ival = 2
}

//pass references
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initials: ", i)

	zeroval(i)
	//zeroval will get a copy of ival distinct from the one in the calling function.
	fmt.Println("zeroval: ", i)

	//The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	//Assigning a value to a dereferenced pointer changes the value at the referenced address.
	zeroptr(&i)
	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("zeropt: ", i)
	// Pointers can be printed too.
	fmt.Println("pointer", &i)

	//zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
}
