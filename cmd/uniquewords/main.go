package main

import (
	"os"

	"github.com/stevedesilva/gologparser/scanner"
)

func main() {

	s := scanner.New(os.Stdin)
	s.UniqueWords()

}
