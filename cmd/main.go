package main

import (
	"fmt"
	"os"

	"github.com/stevedesilva/gologparser/scanner"
)

// ---------------------------------------------------------
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func main() {

	s := scanner.New(os.Stdin)
	//s.Uppercaser()
	// s.UniqueWords()

	// if args := os.Args[1:]; len(args) != 1 {
	// 	fmt.Println("Only one pattern should be supplied!")
	// } else {
	// 	pattern := args[0]

	// 	fmt.Println("pattern = ", pattern)
	// 	sp := scanner.NewPattern(os.Stdin, pattern)

	// 	r := sp.GrepClone()
	// 	printMatchedGrepClone(r)
	// }


	s.Quit()

}

func printMatchedGrepClone(r []string) {
	fmt.Println("Matched :")
	for _, v := range r {
		fmt.Printf("%v \n", v)
	}
}
