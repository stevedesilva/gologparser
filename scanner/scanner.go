package scanner

import (
	"bufio"
	"fmt"
	"io"
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

// Process method
func (s *MyScanner) newscanner() *bufio.Scanner {
	return bufio.NewScanner(s.r)
}
