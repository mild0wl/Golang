package main


import ( "fmt"
		"math"
)
func main(){
	var year float64
	println("Enter the value of year")
	fmt.Scan(&year)
	sol := yearCentury(year)
	fmt.Println(sol)
}

func yearCentury(year float64) float64{
	return math.Ceil(year/100.0)
}