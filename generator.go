package passgen

import (
	"fmt"
	"math/rand/v2"
)

const pkg = "passgen.generator"
const emptyPassword = ""

type Standard interface {
	Template() string
	Length() uint
}

func Generate(standard Standard) (string, error) {
	template := standard.Template()
	if err := validateTemplate(template); err != nil {
		return emptyPassword, err
	}

	length := standard.Length()
	if err := validateLength(length); err != nil {
		return emptyPassword, err
	}

	return generate(length, template), nil
}

func validateTemplate(template string) error {
	if len(template) == 0 {
		return fmt.Errorf("%s: %w", pkg, ErrEmptyTemplate)
	}

	return nil
}

func validateLength(length uint) error {
	if length == 0 {
		return fmt.Errorf("%s: %w", pkg, ErrZeroLength)
	}

	return nil
}

func generate(length uint, template string) string {
	password := make([]rune, length)
	symbols := []rune(template)

	for i := range password {
		index := rand.Int() % len(template)
		password[i] = symbols[index]
	}

	return string(password)
}
