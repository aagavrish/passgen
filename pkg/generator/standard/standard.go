package standard

import (
	"math/rand/v2"
	"strings"
)

type Template string
type Range struct {
	Min uint
	Max uint
}

type Standard struct {
	r       Range
	formats []Template
}

func (s Standard) GetLength() uint {
	return rand.UintN(s.r.Max-s.r.Min+1) + s.r.Min
}

func (s Standard) GetTemplate() string {
	var f strings.Builder
	for _, template := range s.formats {
		f.WriteString(string(template))
	}

	return f.String()
}

func CreateStandard(r Range, formats ...Template) Standard {
	return Standard{
		r:       r,
		formats: append([]Template(nil), formats...),
	}
}

func WithoutRange(length uint) Range {
	return Range{
		Min: length,
		Max: length,
	}
}

func WithRange(minLength, maxLength uint) Range {
	return Range{
		Min: minLength,
		Max: maxLength,
	}
}
