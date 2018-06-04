package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {						//i = init ; loop syntax for showing first 200 character UTF-8
		fmt.Printf("%d \t %q \n", i, i)	//%d decimal ; %q showing single quote character
	}
}
