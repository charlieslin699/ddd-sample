package query

import "context"

type Query[input, output any] interface {
	Execute(context.Context, input) (output, error)
}
