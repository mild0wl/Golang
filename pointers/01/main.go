package main

import (
	"fmt"
)

func main()  {

	var a = 42

	fmt.Println(a)

	var b *int = &a

	fmt.Printf("%d\n",&a)	 //assigning b gives the address of "a" where it is stored

	fmt.Println(b)   //assigning b gives the address of "a" where it is stored
	fmt.Println(*b) //assigning *b gives the value of "a" where it is stored

	*b = 56

	fmt.Println(a)
}

