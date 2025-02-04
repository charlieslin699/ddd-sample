package util

import "sync"

func NewLazyInitalize[T any](initFunc func() T) func() T {
	once := new(sync.Once)
	var instance T
	return func() T {
		once.Do(func() {
			instance = initFunc()
		})

		return instance
	}
}
