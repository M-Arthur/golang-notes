package main

import (
	"fmt"
)

// This shows how the rune works in Go
// Rune would work as the same as the byte if the string only contains ACSII (such as 32 - 127 => a-zA-Z0-9) valid characters.
// You can try to replace the string to test it by yourself.
func main() {
	s := "这是个测试"
	r := []rune(s) // Represent string with rune slice
	b := []byte(s) // Represent string with byte slice

	// Output both rune and byte slices
	// It is clear that the values and length of two slices are different
	fmt.Printf("Rune slice: %v\n", r)
	fmt.Printf("Byte slice: %v\n", b)

	// Output the first value and convert it to string
	// We can see only one Chinese character is composed of three byte values
	fmt.Printf("First value in rune slice: %v\n", string(r[0]))
	fmt.Printf("First value in byte slice: %v\n", string(b[0]))
	fmt.Printf("Combination of first three values of byte slice: %v\n", string(b[0:3]))

	// The length of string is same as the length of the byte slice, therefore we validate that the string are stored in byte slice format
	// The length of rune slice is equal to the sentence length in that string
	fmt.Printf("Length of string: %v\n", len(r))
	fmt.Printf("Length of string: %v\n", len(s))
	fmt.Printf("Length of byte slice: %v\n", len(b))
}
