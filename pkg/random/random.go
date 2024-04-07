package random

import (
	"math/rand"
)

type Random interface {
	Int() int
}

type random struct {
	r      *rand.Rand
	option randomOption
}

func NewRandom(fns ...RandomOptionFunc) Random {
	option := newOption(fns...)
	r := rand.New(rand.NewSource(option.seed))

	return &random{r, option}
}

func (r *random) Int() int {
	return r.r.Int() // TODO: 併發不安全，待調整
}
