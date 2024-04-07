package command

import "context"

type Command[input, output any] interface {
	Execute(context.Context, input) (output, error)
}
