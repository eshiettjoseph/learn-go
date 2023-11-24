package main 

import "fmt"
	

func testPointer(x *int) {
	*x = 2
}

func main() {
	test := 5
	testPointer(&test)
	fmt.Println(test)
}


