package context

import (
	"context"
	"errors"
)

type (
	ContextKey string
)

var (
	// SomeCustomKey is a key that can be used to store a value in a context.Context.
	SomeCustomKey ContextKey = "a_key_to_store_in_ctx"
)

// IsCancelledError determines if the specified error is an instance of context.Cancelled.
func IsCanceledError(err error) bool {
	return errors.Is(err, context.Canceled)
}
