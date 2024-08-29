package standard

import (
	"math/rand/v2"
	"strings"
)

type Template string

type Standard struct {
	r       Range
	formats []Template
}

func CreateStandard(r Range, formats ...Template) Standard {
	return Standard{
		r:       r,
		formats: append([]Template(nil), formats...),
	}
}

func (s Standard) Length() uint {
	return rand.UintN(s.r.Max-s.r.Min+1) + s.r.Min
}

func (s Standard) Template() string {
	var f strings.Builder
	for _, template := range s.formats {
		f.WriteString(string(template))
	}

	return f.String()
}
