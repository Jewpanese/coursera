package main

import (
	"fmt"
	"strings"
)

// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
//
// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

func main() {

	fmt.Println("Please enter a string.")

	string1 := "ian"
	string2 := "Ian"
	string3 := "iuiygaygn"
	string4 := "I d skd a efju N"
	string5 := "ihhhhhn"
	string6 := "ina"
	string7 := "xian"

	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(string3)
	fmt.Println(string4)
	fmt.Println(string5)
	fmt.Println(string6)
	fmt.Println(string7)

	findian1 := strings.ContainsAny(string1, "i, a, n")
	findian2 := strings.ContainsAny(string2, "i, a, n")
	findian3 := strings.ContainsAny(string3, "i, a, n")
	findian4 := strings.ContainsAny(string4, "i, a, n")
	findian5 := strings.ContainsAny(string5, "i, a, n")
	findian6 := strings.ContainsAny(string6, "i, a, n")
	findian7 := strings.ContainsAny(string7, "i, a, n")

	fmt.Println(findian1, findian2, findian3, findian4, findian5, findian6, findian7)

	fmt.Println("FOUND!")

	fmt.Println("Not Found!")
}
