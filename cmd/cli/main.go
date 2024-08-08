package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aagavrish/passgen/pkg/generator"
	"github.com/aagavrish/passgen/pkg/generator/examples"
	"github.com/aagavrish/passgen/pkg/generator/standard"
)

const (
	defaultMinLength = 10
	defaultMaxLength = 20
)

func main() {
	defaultStandard := standard.CreateStandard(
		standard.WithRange(defaultMinLength, defaultMaxLength),
		examples.Digits, examples.LowerLetters, examples.UpperLetters, examples.Special)

	var (
		minLength uint
		maxLength uint
		format    string
	)
	flag.StringVar(&format, "format", "", "--format=abcde123klaj90")
	flag.UintVar(&minLength, "min_length", 0, "--min_length=15")
	flag.UintVar(&maxLength, "max_length", 0, "--max_length=30")
	flag.Parse()

	var std standard.Standard
	if format == "" || minLength == 0 || maxLength == 0 {
		std = defaultStandard
	} else {
		std = standard.CreateStandard(
			standard.WithRange(minLength, maxLength),
			standard.Template(format))
	}

	password, err := generator.Generate(std)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)
}
