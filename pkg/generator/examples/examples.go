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
	NIST = standard.CreateStandard(
		standard.Range{Min: 8, Max: 15},
		Digits, LowerLetters)

	// ISOIEC27001 ...
	ISOIEC27001 = standard.CreateStandard(
		standard.Range{Min: 12, Max: 20},
		Digits, LowerLetters, UpperLetters, Special)

	// PSIDSS ...
	PSIDSS = standard.CreateStandard(
		standard.Range{Min: 7, Max: 15},
		Digits, LowerLetters, UpperLetters)

	// CIS ...
	CIS = standard.CreateStandard(
		standard.Range{Min: 14, Max: 20},
		Digits, LowerLetters, UpperLetters, Special)
)
