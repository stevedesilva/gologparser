package scanner

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// MyScanner type
type MyScanner struct {
	r io.Reader
}

// New construct
func New(reader io.Reader) *MyScanner {
	return &MyScanner{
		r: reader,
	}
}

// Uppercaser converts input to uppercase
func (s *MyScanner) Uppercaser() (result []string) {

	in := s.newscanner()
	for in.Scan() {
		line := strings.ToUpper(in.Text())
		fmt.Println(line)
		result = append(result, line)
	}

	return
}

// UniqueWords method
// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------
func (s *MyScanner) UniqueWords() (total, unique int) {
	// find any character except letters 1 or more times
	rx := regexp.MustCompile(`[^A-Za-z]+`)

	in := s.newscanner()
	in.Split(bufio.ScanWords)

	words := make(map[string]int)

	for in.Scan() {
		total++

		word := strings.ToLower(in.Text())
		// remove everything except the letters
		word = rx.ReplaceAllString(word, "")

		words[word]++

		fmt.Println(word)

	}
	unique = len(words)
	fmt.Printf("Total words: %d\t Unique words:%d \n", total, unique)
	fmt.Printf("Map values: %v", words)
	return total, len(words)
}

// Process method
func (s *MyScanner) newscanner() *bufio.Scanner {
	return bufio.NewScanner(s.r)
}
