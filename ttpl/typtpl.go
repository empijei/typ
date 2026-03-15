// Package ttpl connects a template to the type it can render.
package ttpl

import (
	"bytes"
	"html/template"
	"io"

	"github.com/empijei/typ/tsync"
)

var pool = tsync.NewPool(func() *bytes.Buffer { return &bytes.Buffer{} })

// Template is a type-safe wrapper around [template.Template].
type Template[T any] struct {
	tpl *template.Template
}

// New creates a new [Template] with the given name.
func New[T any](name string) *Template[T] {
	return &Template[T]{template.New(name)}
}

// Execute applies a parsed template to the specified data object of type T,
// and writes the output to w.
func (t *Template[T]) Execute(w io.Writer, data T) error {
	buf := pool.Get()
	buf.Reset()
	if err := t.tpl.Execute(buf, data); err != nil {
		return err
	}
	_, _ = io.Copy(w, buf)
	return nil
}

// Delims sets the action delimiters to the specified strings, to be used in
// subsequent calls to Parse.
func (t *Template[T]) Delims(left, right string) *Template[T] {
	t.tpl = t.tpl.Delims(left, right)
	return t
}

// Funcs adds the elements of the argument map to the template's function map.
// It must be called before the template is parsed.
func (t *Template[T]) Funcs(funcMap template.FuncMap) *Template[T] {
	t.tpl = t.tpl.Funcs(funcMap)
	return t
}

// Option sets options for the template. Options are described by strings,
// either a simple name or "key=value".
func (t *Template[T]) Option(opt ...string) *Template[T] {
	t.tpl = t.tpl.Option(opt...)
	return t
}

// Parse parses text as a template body for t.
func (t *Template[T]) Parse(text string) (*Template[T], error) {
	tpl, err := t.tpl.Parse(text)
	if err != nil {
		return t, err
	}
	t.tpl = tpl
	return t, err
}
