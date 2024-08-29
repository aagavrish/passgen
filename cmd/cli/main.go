package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aagavrish/passgen"
	"github.com/aagavrish/passgen/examples"
	"github.com/aagavrish/passgen/standard"
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

	std := defaultStandard
	if format != "" && minLength != 0 && maxLength != 0 {
		std = standard.CreateStandard(
			standard.WithRange(minLength, maxLength),
			standard.Template(format))
	}

	password, err := passgen.Generate(std)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)
}
