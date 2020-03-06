package scanner_test

import (
	"strings"
	"testing"

	"github.com/stevedesilva/gologparser/scanner"
	"github.com/stretchr/testify/assert"
)

func TestScannerReturnsNew(t *testing.T) {
	s := scanner.New(strings.NewReader("test"))
	assert.NotNil(t, s)
}

func TestScannerShouldConvertLinesToUpperCase(t *testing.T) {

	reader := strings.NewReader("test sentence one\ntest sentence two\ntest sentence three")
	s := scanner.New(reader)

	result := s.Uppercaser()

	expected := []string{
		"TEST SENTENCE ONE",
		"TEST SENTENCE TWO",
		"TEST SENTENCE THREE",
	}

	assert.NotNil(t, s)
	assert.Equal(t, expected, result)
}

// There are 99 words, 70 of them are unique.
func TestScannerShouldReturnTotalAndUniqueWords(t *testing.T) {

	reader := strings.NewReader("test sentence one\ntest sentence two\ntest sentence three")
	s := scanner.New(reader)

	total, unique := s.UniqueWords()

	expectedTotal := 9
	expectedUnique := 5

	assert.NotNil(t, s)
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, expectedUnique, unique)
}

func TestScannerGrepOnlyShouldReturnAllLinesWhenNoInput(t *testing.T) {

	reader := strings.NewReader("test sentence one\ntest sentence two\ntest sentence three")
	s := scanner.New(reader)

	result := s.GrepClone()

	expected := []string{
		"test sentence one",
		"test sentence two",
		"test sentence three",
	}

	assert.NotNil(t, s)
	assert.Equal(t, expected, result)
}

func TestScannerGrepOnlyShouldReturnLinesWhichMatchPattern(t *testing.T) {

	reader := strings.NewReader("test sentence one\ntest sentence two\ntest sentence three")
	s := scanner.NewPattern(reader, "one")

	result := s.GrepClone()

	expected := []string{
		"test sentence one",
	}

	assert.NotNil(t, s)
	assert.Equal(t, expected, result)
}

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
func TestScannerQuitShouldExitWhenSameWordIsTypeTwice(t *testing.T) {


	reader := strings.NewReader("one\ntwo\none")
	s := scanner.New(reader)

	result := s.Quit()

	expected := "TWICE"
	assert.NotNil(t, s)
	assert.Equal(t, expected, result)
}
