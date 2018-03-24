package main

import "fmt"

var a = 5

func main(){
	//a:=5 if we declar x here func main1 doesn't have an access of x. It is type block scope
	fmt.Println(a)
	main1()

}
func main1(){
	
	fmt.Println(a)
}