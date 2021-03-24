package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This function is from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559
// You have to go run . command under the same directory where the example.txt is located
// Go search file based on the directory the command run instead of directory where the .go file is
// There is one caveat: Scanner does not deal well with lines longer than 65536 characters.
// If that is an issue for you then then you should probably roll your own on top of Reader.Read().
func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
