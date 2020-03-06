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
