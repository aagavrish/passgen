package generator

import "math/rand/v2"

type Standard interface {
	GetTemplate() string
	GetLength() uint
}

func Generate(standard Standard) (string, error) {
	template := standard.GetTemplate()
	if len(template) == 0 {
		return "", ErrEmptyTemplate
	}

	length := standard.GetLength()

	password := make([]byte, length)
	for i := range password {
		index := rand.Int() % len(template)
		password[i] = template[index]
	}

	return string(password), nil
}
