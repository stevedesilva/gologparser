package scanner_test

import (
	"testing"

	"github.com/stevedesilva/gologparser/scanner"
	"github.com/stretchr/testify/assert"
)

func TestScannerReturnsNew(t *testing.T) {
	s := scanner.New()
	assert.NotNil(t, s)
}
