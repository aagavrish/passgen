package main

import (
	"flag"
	"fmt"

	"github.com/aagavrish/passgen/pkg/generator"
	"github.com/aagavrish/passgen/pkg/generator/examples"
	"github.com/aagavrish/passgen/pkg/generator/standard"
)

func main() {
	defaultStandard := standard.CreateStandard(20,
		examples.Digits, examples.LowerLetters, examples.UpperLetters, examples.Special)

	var (
		l uint
		t string
	)
	flag.StringVar(&t, "template", "", "USAGE")
	flag.UintVar(&l, "length", 0, "USAGE")
	flag.Parse()

	var std standard.Standard
	if t == "" || l == 0 {
		std = defaultStandard
	} else {
		std = standard.CreateStandard(l, standard.Template(t))
	}

	password := generator.Generate(std)
	fmt.Println(password)
}
