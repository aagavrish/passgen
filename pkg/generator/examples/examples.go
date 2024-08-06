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
	NIST = standard.CreateStandard("NIST", 8,
		Digits, LowerLetters)

	// ISOIEC27001 ...
	ISOIEC27001 = standard.CreateStandard("ISO/IEC 27001", 12,
		Digits, LowerLetters, UpperLetters, Special)

	// PSIDSS ...
	PSIDSS = standard.CreateStandard("PCI DSS", 7,
		Digits, LowerLetters, UpperLetters)

	// CIS ...
	CIS = standard.CreateStandard("CIS", 14,
		Digits, LowerLetters, UpperLetters, Special)
)
