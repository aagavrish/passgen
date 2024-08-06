# Password Generator

```go
Digits       standard.Template = "0123456789"
LowerLetters standard.Template = "abcdefghijklmnopqrstuvwxyz"
UpperLetters standard.Template = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
Special      standard.Template = "~`!@#$%^&*()-+={}[]."
```

```go
defaultStandard := standard.CreateStandard(20,
    examples.Digits, examples.LowerLetters, examples.UpperLetters, examples.Special)
password := generator.Generate(defaultStandard)
```

```go
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
```