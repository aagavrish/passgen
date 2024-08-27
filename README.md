# Password Generator

```go
package main

import (
	"fmt"
	"log"

	"github.com/aagavrish/passgen"
	"github.com/aagavrish/passgen/examples"
	"github.com/aagavrish/passgen/standard"
)

func main() {
	// Creating a password standard with a fixed length and a specified character set.
	std := standard.CreateStandard(standard.WithoutRange(10),
		examples.Digits, examples.LowerLetters, examples.UpperLetters, examples.Special)

	// You can also set a standard with a floating length: standard.WithRange(minLength, maxLength), where
	// minLength - minimum password length,
	// maxLength - maximum password length (inclusive)
	// std := standard.CreateStandard(standard.WithRange(5, 10),
	//     examples.Digits, examples.LowerLetters, examples.UpperLetters, examples.Special)

	// Generating a password of a given standard
	password, err := passgen.Generate(std)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)
}
```