package main

import "fmt"

func main(){
	x := 4
	fmt.Println(x) 
	{
		fmt.Println(x) //here x can be used because it is inside of the total scope
		y := "Let's celebrate the LOVE! :-)"
		fmt.Println(y)
	}
	//  fmt.Println(y)  // here y is not usable which  is outside of scope
}