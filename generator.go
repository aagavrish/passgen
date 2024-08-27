package passgen

import (
	"fmt"
	"math/rand/v2"
)

const pkg = "passgen.generator"
const emptyPassword = ""

type Standard interface {
	GetTemplate() string
	GetLength() uint
}

func Generate(standard Standard) (string, error) {
	template := standard.GetTemplate()
	if len(template) == 0 {
		return emptyPassword, fmt.Errorf("%s: %w", pkg, ErrEmptyTemplate)
	}

	length := standard.GetLength()
	if length == 0 {
		return emptyPassword, fmt.Errorf("%s: %w", pkg, ErrZeroLength)
	}

	return generate(length, template), nil
}

func generate(length uint, template string) string {
	password := make([]byte, length)
	for i := range password {
		index := rand.Int() % len(template)
		password[i] = template[index]
	}

	return string(password)
}
