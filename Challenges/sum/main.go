package main

import ( "fmt"
)

func main(){
	var val1 int 
	var val2 int
	fmt.Println("Enter the value of param1 and param2")
	fmt.Scan(&val1)
	fmt.Scan(&val2)
	var sum  = add(val1, val2)
	fmt.Println(sum)
}

func add(param1 int , param2 int) int{
	return param1 + param2
}