package standard

import "strings"

type Template string

type Standard struct {
	length  uint
	formats []Template
}

func (s Standard) GetLength() uint {
	return s.length
}

func (s Standard) GetTemplate() string {
	var f strings.Builder
	for _, template := range s.formats {
		f.WriteString(string(template))
	}

	return f.String()
}

func CreateStandard(passLength uint, formats ...Template) Standard {
	return Standard{
		length:  passLength,
		formats: append([]Template(nil), formats...),
	}
}
