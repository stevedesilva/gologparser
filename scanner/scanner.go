package scanner

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// MyScanner type
type MyScanner struct {
	r io.Reader
	p string
	a []string
}

// New construct
func New(reader io.Reader) *MyScanner {
	return &MyScanner{
		r: reader,
	}
}

// NewPattern construct with pattern
func NewPattern(reader io.Reader, pattern string) *MyScanner {
	return &MyScanner{
		r: reader,
		p: pattern,
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

// GrepClone struct
// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------
func (s *MyScanner) GrepClone() (results []string) {

	in := s.newscanner()
	rxPattern := regexp.MustCompile(s.p)
	for in.Scan() {

		input := in.Text()
		if rxPattern.MatchString(input) {
			results = append(results, input)
		}

	}

	return
}

// Quit exe
// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------
func (s *MyScanner) Quit() (result string) {
	const twice = "TWICE"
	found := make(map[string]bool)
	in := s.newscanner()

	for in.Scan() {

		input := strings.ToLower(in.Text())

		if found[input] {
			fmt.Println(twice, input)
			return twice
		}
		fmt.Println("Adding", input)
		found[input] = true

	}
	return
}

// LogParser exe
// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------
func (s *MyScanner) LogParser() (result string, err error) {
	var (
		sum     = make(map[string]int, 0)
		domains []string
		total   int
		lines   int
	)

	in := s.newscanner()
	for in.Scan() {
		lines++
		// convert input to lower
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			err := fmt.Errorf("wrong input: %v (line #%d)", fields, lines)
			return result, err
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			err := fmt.Errorf("wrong input: %v (line #%d)", fields[1], lines)
			return result, err
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		sum[domain] += visits
		total += visits
	}

	fmt.Printf("%-30s %-10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %-10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %-10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		return result, err
	}
	return
}
