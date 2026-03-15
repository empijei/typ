package tsync

import (
	"testing"

	"github.com/empijei/tst"
)

func TestMap(t *testing.T) {
	tst.Go(t)

	var m Map[string, int]

	t.Run("StoreLoad", func(t *testing.T) {
		m.Store("foo", 42)
		v, ok := m.Load("foo")
		tst.Be(ok, t)
		tst.Is(42, v, t)

		v = m.Get("foo")
		tst.Is(42, v, t)
	})

	t.Run("LoadNonExistent", func(t *testing.T) {
		v, ok := m.Load("bar")
		tst.Be(!ok, t)
		tst.Is(0, v, t)

		v = m.Get("bar")
		tst.Is(0, v, t)
	})

	t.Run("Delete", func(t *testing.T) {
		m.Store("to_delete", 100)
		m.Delete("to_delete")
		_, ok := m.Load("to_delete")
		tst.Be(!ok, t)
	})

	t.Run("LoadAndDelete", func(t *testing.T) {
		m.Store("load_delete", 200)
		v, loaded := m.LoadAndDelete("load_delete")
		tst.Be(loaded, t)
		tst.Is(200, v, t)
		_, ok := m.Load("load_delete")
		tst.Be(!ok, t)

		v, loaded = m.LoadAndDelete("non_existent")
		tst.Be(!loaded, t)
		tst.Is(0, v, t)
	})

	t.Run("LoadOrStore", func(t *testing.T) {
		v, loaded := m.LoadOrStore("load_or_store", 300)
		tst.Be(!loaded, t)
		tst.Is(300, v, t)

		v, loaded = m.LoadOrStore("load_or_store", 400)
		tst.Be(loaded, t)
		tst.Is(300, v, t)
	})

	t.Run("Swap", func(t *testing.T) {
		m.Store("swap", 500)
		prev, loaded := m.Swap("swap", 600)
		tst.Be(loaded, t)
		tst.Is(500, prev, t)
		v, ok := m.Load("swap")
		tst.Be(ok, t)
		tst.Is(600, v, t)

		prev, loaded = m.Swap("swap_new", 700)
		tst.Be(!loaded, t)
		tst.Is(0, prev, t)
	})

	t.Run("CompareAndSwap", func(t *testing.T) {
		m.Store("cas", 800)
		swapped := m.CompareAndSwap("cas", 800, 900)
		tst.Be(swapped, t)
		v, _ := m.Load("cas")
		tst.Is(900, v, t)

		swapped = m.CompareAndSwap("cas", 800, 1000)
		tst.Be(!swapped, t)
		v, _ = m.Load("cas")
		tst.Is(900, v, t)
	})

	t.Run("CompareAndDelete", func(t *testing.T) {
		m.Store("cad", 1100)
		deleted := m.CompareAndDelete("cad", 1100)
		tst.Be(deleted, t)
		_, ok := m.Load("cad")
		tst.Be(!ok, t)

		m.Store("cad", 1200)
		deleted = m.CompareAndDelete("cad", 1100)
		tst.Be(!deleted, t)
		v, _ := m.Load("cad")
		tst.Is(1200, v, t)
	})

	t.Run("Clear", func(t *testing.T) {
		m.Store("a", 1)
		m.Store("b", 2)
		m.Clear()
		_, ok := m.Load("a")
		tst.Be(!ok, t)
		_, ok = m.Load("b")
		tst.Be(!ok, t)
	})

	t.Run("All", func(t *testing.T) {
		m.Clear()
		m.Store("a", 1)
		m.Store("b", 2)

		found := make(map[string]int)
		for k, v := range m.All() {
			found[k] = v
		}

		tst.Is(2, len(found), t)
		tst.Is(1, found["a"], t)
		tst.Is(2, found["b"], t)
	})
}
