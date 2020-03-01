package scanner_test

import (
	"strings"
	"testing"

	"github.com/stevedesilva/gologparser/scanner"
	"github.com/stretchr/testify/assert"
)

func TestScannerReturnsNew(t *testing.T) {
	s := scanner.New()
	assert.NotNil(t, s)
}

func TestScannerShouldConvertLinesToUpperCase(t *testing.T) {
	s := scanner.New()
	reader := strings.NewReader("test sentence one\ntest sentence two\ntest sentence three")
	result := s.Process(reader)

	expected := []string{
		"TEST SENTENCE ONE",
		"TEST SENTENCE TWO",
		"TEST SENTENCE THREE",
	}

	assert.NotNil(t, s)
	assert.Equal(t, expected, result)
}
