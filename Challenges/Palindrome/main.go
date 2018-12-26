package main

import ( "fmt"
)

func main(){
	var palindrome string
	fmt.Println("Enter the word which is palindrome")
	fmt.Scan(&palindrome)
	input := isPalindrome(palindrome)
	fmt.Println(input)
}

func isPalindrome(word string) bool{
	for i := 0 ; i < len(word)/2 ; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}