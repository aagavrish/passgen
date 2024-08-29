package standard

type Range struct {
	Min uint
	Max uint
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
