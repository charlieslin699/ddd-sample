package random

import "time"

type RandomOptionFunc func(o *randomOption)

type randomOption struct {
	seed int64
}

func newOption(fns ...RandomOptionFunc) randomOption {
	option := randomOption{
		seed: time.Now().UnixNano(),
	}

	for _, fn := range fns {
		fn(&option)
	}

	return option
}

func WithSeed(seed int64) RandomOptionFunc {
	return func(o *randomOption) {
		o.seed = seed
	}
}
