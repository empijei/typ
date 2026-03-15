package tsync

import (
	"fmt"
	"reflect"
	"sync"
)

// Pool is a type safe sync.Pool.
type Pool[T any] struct {
	p *sync.Pool
}

// NewPool constructs a new Pool.
func NewPool[T any](factory func() T) Pool[T] {
	if reflect.TypeFor[T]().Kind() != reflect.Pointer {
		panic(fmt.Sprintf("%T used to create new Pool is not a pointer", *new(T)))
	}
	return Pool[T]{
		p: &sync.Pool{
			New: func() any {
				return factory()
			},
		},
	}
}

// Put returns the value to the pool.
func (p *Pool[T]) Put(t T) { p.p.Put(t) }

// Get returns a value from the pool.
//
// If the value needs to be reset, it's the responsibility of the caller to do so.
func (p *Pool[T]) Get() T { return p.p.Get().(T) } //nolint: errcheck,forcetypeassert // Known to be correct.
