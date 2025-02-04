package uid

import gonanoid "github.com/matoous/go-nanoid/v2"

func NewNanoID(fns ...NanoIDOptionFunc) string {
	option := NanoIDOption{
		Length: 21,
	}

	for _, fn := range fns {
		fn(&option)
	}

	id := gonanoid.Must(option.Length)
	return id
}

type NanoIDOption struct {
	Length int
}

type NanoIDOptionFunc func(o *NanoIDOption)

func NanoIDWithLength(length int) NanoIDOptionFunc {
	return func(o *NanoIDOption) {
		o.Length = length
	}
}
