// Package tctx implements statically strongly typed wrappers for context keys.
package tctx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ctxKey string

// Key is a key for a context.
type Key[T any] struct{ k ctxKey }

// NewKey returns a new unique key.
//
// Names don't have to be unique, but they should be for debugging clarity.
func NewKey[T any](name string) Key[T] {
	return Key[T]{ctxKey(name + "(" + uuid.New().String() + ")")}
}

// Set returns a child context with the specified value for k.
func (k *Key[T]) Set(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, k.k, value)
}

// Get returns the context value, if any.
func (k *Key[T]) Get(ctx context.Context) (value T, ok bool) {
	got := ctx.Value(k.k)
	if got == nil {
		return value, false
	}
	value, ok = got.(T)
	if !ok {
		fmt.Printf("🚨 BUG IN CODE 🚨: Context Key %s maps to type %T but got type %T", k.k, value, got)
		return value, false
	}
	return value, true
}
