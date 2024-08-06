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

	t.Run("generate password with zero length and empty template", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(0, "")
		password, err := Generate(testStandard)

		require.ErrorIs(t, err, ErrEmptyTemplate)
		require.Equal(t, password, "")
	})

	t.Run("generate password with zero length", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(0, "abc")
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, password, "")
	})

	t.Run("generate password with empty template", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(3, "")
		_, err := Generate(testStandard)

		require.ErrorIs(t, err, ErrEmptyTemplate)
	})

	t.Run("generate password with correct args", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(3, examples.Digits)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
	})

	t.Run("generate password from lowercase letters", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(3, examples.LowerLetters)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, false, isDigit(password))
		require.Equal(t, strings.ToLower(password), password)
	})

	t.Run("generate password from uppercase letters", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(3, examples.UpperLetters)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, false, isDigit(password))
		require.Equal(t, strings.ToUpper(password), password)
	})

	t.Run("generate password from digits", func(t *testing.T) {
		t.Parallel()

		testStandard := standard.CreateStandard(3, examples.Digits)
		password, err := Generate(testStandard)

		require.NoError(t, err)
		require.Equal(t, len([]rune(password)), 3)
		require.Equal(t, true, isDigit(password))
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
