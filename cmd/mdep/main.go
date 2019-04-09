package main

import (
	"fmt"

	"github.com/daved/multidep/x"
	"github.com/daved/multidep/y"
)

func main() {
	mx := x.Make()
	for k, v := range mx {
		fmt.Println(k, v())
	}

	my := y.Make()
	for k, v := range my {
		fmt.Println(k, v())
	}
}
