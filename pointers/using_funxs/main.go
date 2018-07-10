package main

import (
	_"fmt"
	"fmt"
)

func zero(a int){
	a=0
	fmt.Println(&a)	//0xc0420080d8

}

func main(){
	var z = 5
	zero(z)
	fmt.Println(z)	// still z is 5
	fmt.Println(&z)	//0xc042004030

	zero1(&z)
	fmt.Println(z)	// now z is 0
	fmt.Println(&z)	//0xc042004030

}

func zero1(b *int) {
	*b =0
	fmt.Println(b) //0xc042004030

}
