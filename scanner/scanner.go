package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Scanner type
type Scanner struct{}

// New construct
func New() *Scanner {
	return &Scanner{}
}

// Process method
func (s *Scanner) Process(r io.Reader) []string {
	in := bufio.NewScanner(r)
	result := []string{}
	for in.Scan() {
		line := strings.ToUpper(in.Text())
		fmt.Println(line)
		result = append(result, line)

	}

	return result
}
