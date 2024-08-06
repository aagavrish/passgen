package generator

import "math/rand/v2"

type Standard interface {
	GetTemplate() string
	GetLength() uint
}

func Generate(standard Standard) string {
	template := standard.GetTemplate()

	password := make([]byte, standard.GetLength())
	for i := range password {
		index := rand.Int() % len(template)
		password[i] = template[index]
	}

	return string(password)
}
