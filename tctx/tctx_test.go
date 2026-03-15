package tctx_test

import (
	"testing"

	"github.com/empijei/tst"
	"github.com/empijei/typ/tctx"
)

func TestKey(t *testing.T) {
	ctx := tst.Go(t)

	key1 := tctx.NewKey[string]("key1")
	key2 := tctx.NewKey[int]("key2")

	// Initial state: keys should not be found
	_, ok := key1.Get(ctx)
	tst.Is(false, ok, t)
	_, ok = key2.Get(ctx)
	tst.Is(false, ok, t)

	// Set key1
	ctx1 := key1.Set(ctx, "hello")
	val1, ok := key1.Get(ctx1)
	tst.Be(ok, t)
	tst.Is("hello", val1, t)

	// key2 still not found in ctx1
	_, ok = key2.Get(ctx1)
	tst.Is(false, ok, t)

	// Set key2 in ctx1
	ctx2 := key2.Set(ctx1, 42)
	val2, ok := key2.Get(ctx2)
	tst.Is(true, ok, t)
	tst.Is(42, val2, t)

	// key1 still found in ctx2
	val1, ok = key1.Get(ctx2)
	tst.Is(true, ok, t)
	tst.Is("hello", val1, t)

	// Overwrite key1
	ctx3 := key1.Set(ctx2, "world")
	val1, ok = key1.Get(ctx3)
	tst.Is(true, ok, t)
	tst.Is("world", val1, t)

	// key2 still found in ctx3
	val2, ok = key2.Get(ctx3)
	tst.Is(true, ok, t)
	tst.Is(42, val2, t)
}
