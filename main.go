package main

import (
	"fmt"
	"github.com/imwill/go-challenge/drum"
)

func main() {
	fmt.Print(drum.DecodeFile("drum/fixtures/pattern_4.splice"))
}
