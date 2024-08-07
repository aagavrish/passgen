package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aagavrish/passgen/pkg/generator"
	"github.com/aagavrish/passgen/pkg/generator/examples"
	"github.com/aagavrish/passgen/pkg/generator/standard"
)

func main() {
	defaultStandard := standard.CreateStandard(standard.Range{Min: 10, Max: 15},
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
		std = standard.CreateStandard(
			standard.Range{Min: l, Max: l},
			standard.Template(t))
	}

	password, err := generator.Generate(std)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)
}
