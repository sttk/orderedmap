package orderedmap_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/orderedmap"
	"testing"
)

type foo struct {
	S string
	N int
}

func TestNew(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)
}

func TestStore_newEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	_, ok := om.Load("foo-0")
	assert.False(t, ok)

	om.Store("foo-0", foo{S: "ABC", N: 123})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	om.Store("foo-1", foo{S: "DEF", N: 456})
	assert.Equal(t, om.Len(), 2)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)
}

func TestStore_rewriteEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	_, ok := om.Load("foo-0")
	assert.False(t, ok)

	om.Store("foo-0", foo{S: "ABC", N: 123})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	om.Store("foo-0", foo{S: "DEF", N: 456})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)
}

func TestSwap_newEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	_, ok := om.Load("foo-0")
	assert.False(t, ok)

	_, loaded := om.Swap("foo-0", foo{S: "ABC", N: 123})
	assert.Equal(t, om.Len(), 1)
	assert.False(t, loaded)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	_, loaded = om.Swap("foo-1", foo{S: "DEF", N: 456})
	assert.Equal(t, om.Len(), 2)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)
}

func TestSwap_rewriteEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	_, ok := om.Load("foo-0")
	assert.False(t, ok)

	prev, loaded := om.Swap("foo-0", foo{S: "ABC", N: 123})
	assert.Equal(t, om.Len(), 1)
	assert.False(t, loaded)
	_ = prev

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	prev, loaded = om.Swap("foo-0", foo{S: "DEF", N: 456})
	assert.Equal(t, om.Len(), 1)
	assert.True(t, loaded)
	assert.Equal(t, prev.S, "ABC")
	assert.Equal(t, prev.N, 123)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)
}

func TestLoadOrStore(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	v, loaded := om.LoadOrStore("foo-0", foo{S: "ABC", N: 123})
	assert.False(t, loaded)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	var ok bool
	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	v, ok = om.Load("foo-1")
	assert.False(t, ok)

	v, loaded = om.LoadOrStore("foo-0", foo{S: "DEF", N: 456})
	assert.True(t, loaded)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	v, ok = om.Load("foo-1")
	assert.False(t, ok)

	v, loaded = om.LoadOrStore("foo-1", foo{S: "DEF", N: 456})
	assert.False(t, loaded)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v.S, "ABC")
	assert.Equal(t, v.N, 123)

	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v.S, "DEF")
	assert.Equal(t, v.N, 456)
}

func TestIterateWithRange_zeroEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	n := 0
	om.Range(func(k string, v foo) bool {
		n++
		return false
	})
	assert.Equal(t, n, 0)
}

func TestIterateWithRange_oneEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	assert.Equal(t, om.Len(), 1)

	n := 0
	om.Range(func(k string, v foo) bool {
		n++
		assert.Equal(t, k, "foo-0")
		assert.Equal(t, v.S, "AAA")
		assert.Equal(t, v.N, 123)
		return false
	})
	assert.Equal(t, n, 1)
}

func TestIterateWithRange_multipleEntries(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	om.Store("foo-1", foo{S: "BBB", N: 456})
	assert.Equal(t, om.Len(), 2)

	n := 0
	om.Range(func(k string, v foo) bool {
		if n == 0 {
			assert.Equal(t, k, "foo-0")
			assert.Equal(t, v.S, "AAA")
			assert.Equal(t, v.N, 123)
		} else if n == 1 {
			assert.Equal(t, k, "foo-1")
			assert.Equal(t, v.S, "BBB")
			assert.Equal(t, v.N, 456)
		} else {
			assert.Fail(t, "n > 1")
		}
		n++
		return false
	})
	assert.Equal(t, n, 1)
}

func TestIterateWithFront_zeroEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	ent := om.Front()
	assert.Nil(t, ent)
}

func TestIterateWithFront_oneEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	assert.Equal(t, om.Len(), 1)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Next()
	assert.Nil(t, ent)
}

func TestIterateWithFront_multipleEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	om.Store("foo-1", foo{S: "BBB", N: 456})
	assert.Equal(t, om.Len(), 2)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value().S, "BBB")
	assert.Equal(t, ent.Value().N, 456)

	ent = ent.Next()
	assert.Nil(t, ent)
}

func TestIterateWithBack_zeroEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	ent := om.Back()
	assert.Nil(t, ent)
}

func TestIterateWithBack_oneEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	assert.Equal(t, om.Len(), 1)

	ent := om.Back()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Prev()
	assert.Nil(t, ent)
}

func TestIterateWithBack_multipleEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	om.Store("foo-0", foo{S: "AAA", N: 123})
	om.Store("foo-1", foo{S: "BBB", N: 456})
	assert.Equal(t, om.Len(), 2)

	ent := om.Back()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value().S, "BBB")
	assert.Equal(t, ent.Value().N, 456)

	ent = ent.Prev()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Prev()
	assert.Nil(t, ent)
}

func TestDelete_zeroEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	ent := om.Front()
	assert.Nil(t, ent)

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	ent = om.Front()
	assert.Nil(t, ent)
}

func TestDelete_oneEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "AAA", N: 123})
	assert.Equal(t, om.Len(), 1)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Next()
	assert.Nil(t, ent)

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	ent = om.Front()
	assert.Nil(t, ent)
}

func TestDelete_multipleEntries(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "AAA", N: 123})
	om.Store("foo-1", foo{S: "BBB", N: 456})
	om.Store("foo-2", foo{S: "CCC", N: 789})
	om.Store("foo-3", foo{S: "DDD", N: 321})
	assert.Equal(t, om.Len(), 4)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-1")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a middle entry
	om.Delete("foo-1")
	assert.Equal(t, om.Len(), 3)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a head entry
	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 2)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a last entry
	om.Delete("foo-3")
	assert.Equal(t, om.Len(), 1)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.Nil(t, ent)
}

func TestLoadAndDelete_zeroEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	ent := om.Front()
	assert.Nil(t, ent)

	_, loaded := om.LoadAndDelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.False(t, loaded)

	ent = om.Front()
	assert.Nil(t, ent)
}

func TestLoadAndDelete_oneEntry(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "AAA", N: 123})
	assert.Equal(t, om.Len(), 1)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value().S, "AAA")
	assert.Equal(t, ent.Value().N, 123)

	ent = ent.Next()
	assert.Nil(t, ent)

	v, loaded := om.LoadAndDelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.True(t, loaded)
	assert.Equal(t, v.S, "AAA")
	assert.Equal(t, v.N, 123)

	ent = om.Front()
	assert.Nil(t, ent)
}

func TestLoadAndDelete_multipleEntries(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "AAA", N: 123})
	om.Store("foo-1", foo{S: "BBB", N: 456})
	om.Store("foo-2", foo{S: "CCC", N: 789})
	om.Store("foo-3", foo{S: "DDD", N: 321})
	assert.Equal(t, om.Len(), 4)

	ent := om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-1")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a middle entry
	v, loaded := om.LoadAndDelete("foo-1")
	assert.Equal(t, om.Len(), 3)
	assert.True(t, loaded)
	assert.Equal(t, v.S, "BBB")
	assert.Equal(t, v.N, 456)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-0")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a head entry
	v, loaded = om.LoadAndDelete("foo-0")
	assert.Equal(t, om.Len(), 2)
	assert.True(t, loaded)
	assert.Equal(t, v.S, "AAA")
	assert.Equal(t, v.N, 123)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-3")

	ent = ent.Next()
	assert.Nil(t, ent)

	// delete a last entry
	v, loaded = om.LoadAndDelete("foo-3")
	assert.Equal(t, om.Len(), 1)
	assert.True(t, loaded)
	assert.Equal(t, v.S, "DDD")
	assert.Equal(t, v.N, 321)

	ent = om.Front()
	assert.NotNil(t, ent)
	assert.Equal(t, ent.Key(), "foo-2")

	ent = ent.Next()
	assert.Nil(t, ent)
}
