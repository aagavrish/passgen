package generator

import (
	"strings"
	"testing"
	"unicode"

	"github.com/aagavrish/passgen/pkg/generator/examples"
	"github.com/aagavrish/passgen/pkg/generator/standard"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	t.Parallel()
	testRange := standard.Range{Min: 3, Max: 3}

	t.Run("generate password with zero length and empty template", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(standard.Range{Min: 0, Max: 0}, "")
		password, err := Generate(testStandard)

		require.ErrorIs(t, err, ErrEmptyTemplate)
		require.Equal(t, password, "")
	})

	t.Run("generate password with zero length", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(standard.Range{Min: 0, Max: 0}, "abc")
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, password, "")
	})

	t.Run("generate password with empty template", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(testRange, "")
		_, err := Generate(testStandard)

		require.ErrorIs(t, err, ErrEmptyTemplate)
	})

	t.Run("generate password with correct args", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(testRange, examples.Digits)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
	})

	t.Run("generate password from lowercase letters", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(testRange, examples.LowerLetters)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, false, isDigit(password))
		require.Equal(t, strings.ToLower(password), password)
	})

	t.Run("generate password from uppercase letters", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(testRange, examples.UpperLetters)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, false, isDigit(password))
		require.Equal(t, strings.ToUpper(password), password)
	})

	t.Run("generate password from digits", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(testRange, examples.Digits)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, true, isDigit(password))
	})

	t.Run("generate password with min and max length", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(standard.Range{
			Min: 5,
			Max: 10,
		}, examples.Digits)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.GreaterOrEqual(t, len([]rune(password)), 5)
		require.LessOrEqual(t, len([]rune(password)), 10)
	})
}

func isDigit(s string) bool {
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}

	return true
}
