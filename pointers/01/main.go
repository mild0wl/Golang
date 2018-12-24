package main

import (
	"fmt"
)

func main()  {

	// var a = 42

	// fmt.Println(a)

	// var b *int
	// *b = &a
	
	// fmt.Printf("%d\n",&a)	 //assigning b gives the address of "a" where it is stored

	// fmt.Println(b)   //assigning b gives the address of "a" where it is stored
	// fmt.Println(*b) //assigning *b gives the value of "a" where it is stored

	// *b = 56

	// fmt.Println(a)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

