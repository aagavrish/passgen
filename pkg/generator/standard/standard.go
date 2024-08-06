package standard

import "strings"

type Template string
type Length uint

type Standard struct {
	name    string
	length  Length
	formats []Template
}

func (s Standard) GetName() string {
	return s.name
}

func (s Standard) GetLength() Length {
	return s.length
}

func (s Standard) GetTemplate() string {
	var f strings.Builder
	for _, template := range s.formats {
		f.WriteString(string(template))
	}

	return f.String()
}

func CreateStandard(name string, passLength Length, formats ...Template) Standard {
	return Standard{
		name:    name,
		length:  passLength,
		formats: append([]Template(nil), formats...),
	}
}
