package main

//#include "square.c"
import "C"
import "fmt"

func main() {
	fmt.Print(makeSquare(9))
}

func makeSquare(a int) int {
	//Convert Go ints to C ints
	aC := C.int(a)

	sum, err := C.square(aC)
	if err != nil {
		return 0
	}

	//Convert C.int result to Go int
	res := int(sum)

	return res
}
