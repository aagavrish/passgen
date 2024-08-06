package examples

import "github.com/aagavrish/passgen/pkg/generator/standard"

const (
	Digits       standard.Template = "0123456789"
	LowerLetters standard.Template = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters standard.Template = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Special      standard.Template = "~`!@#$%^&*()-+={}[]."
)

var (
	// NIST ...
	NIST = standard.CreateStandard(8,
		Digits, LowerLetters)

	// ISOIEC27001 ...
	ISOIEC27001 = standard.CreateStandard(12,
		Digits, LowerLetters, UpperLetters, Special)

	// PSIDSS ...
	PSIDSS = standard.CreateStandard(7,
		Digits, LowerLetters, UpperLetters)

	// CIS ...
	CIS = standard.CreateStandard(14,
		Digits, LowerLetters, UpperLetters, Special)
)
