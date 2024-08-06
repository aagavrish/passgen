package generator

import (
	"math/rand/v2"

	"github.com/aagavrish/passgen/pkg/generator/standard"
)

func Generate(standard standard.Standard) string {
	template := standard.GetTemplate()

	password := make([]byte, standard.GetLength())
	for i, _ := range password {
		index := rand.Int() % len(template)
		password[i] = template[index]
	}

	return string(password)
}
