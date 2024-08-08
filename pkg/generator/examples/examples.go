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
		standard.WithRange(8, 15),
		Digits, LowerLetters)

	// ISOIEC27001 ...
	ISOIEC27001 = standard.CreateStandard(
		standard.WithRange(12, 20),
		Digits, LowerLetters, UpperLetters, Special)

	// PSIDSS ...
	PSIDSS = standard.CreateStandard(
		standard.WithRange(7, 15),
		Digits, LowerLetters, UpperLetters)

	// CIS ...
	CIS = standard.CreateStandard(
		standard.WithRange(14, 20),
		Digits, LowerLetters, UpperLetters, Special)
)
