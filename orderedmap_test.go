package orderedmap_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/orderedmap"
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

func TestLdelete_and_Load(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	assert.Equal(t, om.Len(), 0)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	assert.Equal(t, om.Len(), 0)
}

func TestLdelete_and_Store(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, ok = om.Load("foo-0")
	assert.False(t, ok)

	assert.Nil(t, om.Front())

	om.Store("foo-0", foo{S: "AA", N: 11})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	om.Store("foo-1", foo{S: "B", N: 2})
	assert.Equal(t, om.Len(), 2)
	om.Store("foo-2", foo{S: "C", N: 3})
	assert.Equal(t, om.Len(), 3)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-1")
	assert.Equal(t, om.Len(), 2)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-2")
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	om.Store("foo-2", foo{S: "CC", N: 33})
	assert.Equal(t, om.Len(), 2)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	assert.Nil(t, ent.Next())

	om.Store("foo-1", foo{S: "BB", N: 22})
	assert.Equal(t, om.Len(), 3)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "BB", N: 22})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "BB", N: 22})
	assert.Nil(t, ent.Next())
}

func TestLdelete_and_Swap(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	prev, loaded := om.Swap("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)
	assert.False(t, loaded)
	_ = prev

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	assert.Nil(t, om.Front())

	prev, loaded = om.Swap("foo-0", foo{S: "AA", N: 11})
	assert.Equal(t, om.Len(), 1)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	prev, loaded = om.Swap("foo-1", foo{S: "B", N: 2})
	assert.Equal(t, om.Len(), 2)
	assert.False(t, loaded)
	prev, loaded = om.Swap("foo-2", foo{S: "C", N: 3})
	assert.Equal(t, om.Len(), 3)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-1")
	assert.Equal(t, om.Len(), 2)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-2")
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	prev, loaded = om.Swap("foo-2", foo{S: "CC", N: 33})
	assert.Equal(t, om.Len(), 2)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	assert.Nil(t, ent.Next())

	prev, loaded = om.Swap("foo-1", foo{S: "BB", N: 22})
	assert.Equal(t, om.Len(), 3)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "BB", N: 22})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "BB", N: 22})
	assert.Nil(t, ent.Next())

	prev, loaded = om.Swap("foo-0", foo{S: "AAA", N: 111})
	assert.True(t, loaded)
	assert.Equal(t, prev, foo{S: "AA", N: 11})
	prev, loaded = om.Swap("foo-1", foo{S: "BBB", N: 222})
	assert.True(t, loaded)
	assert.Equal(t, prev, foo{S: "BB", N: 22})
	prev, loaded = om.Swap("foo-2", foo{S: "CCC", N: 333})
	assert.True(t, loaded)
	assert.Equal(t, prev, foo{S: "CC", N: 33})
}

func TestLdelete_and_LoadOrStore(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	v, loaded := om.LoadOrStore("foo-0", foo{S: "A", N: 1})
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	assert.Nil(t, om.Front())

	v, loaded = om.LoadOrStore("foo-0", foo{S: "AA", N: 11})
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadOrStore("foo-1", foo{S: "B", N: 2})
	assert.Equal(t, om.Len(), 2)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, loaded = om.LoadOrStore("foo-2", foo{S: "C", N: 3})
	assert.Equal(t, om.Len(), 3)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "C", N: 3})

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-1")
	assert.Equal(t, om.Len(), 2)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-2")
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadOrStore("foo-2", foo{S: "CC", N: 33})
	assert.Equal(t, om.Len(), 2)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadOrStore("foo-1", foo{S: "BB", N: 22})
	assert.Equal(t, om.Len(), 3)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "BB", N: 22})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "CC", N: 33})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "CC", N: 33})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "BB", N: 22})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadOrStore("foo-0", foo{S: "AAA", N: 111})
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AA", N: 11})
	v, loaded = om.LoadOrStore("foo-1", foo{S: "BBB", N: 222})
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "BB", N: 22})
	v, loaded = om.LoadOrStore("foo-2", foo{S: "CCC", N: 333})
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "CC", N: 33})
}

func TestLdelete_and_Delete(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	om.Delete("foo-0")
	assert.Equal(t, om.Len(), 0)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)
}

func TestLdelete_and_LoadAndDelete(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, loaded := om.LoadAndDelete("foo-0")
	assert.False(t, loaded)
	assert.Equal(t, om.Len(), 0)
	_ = v

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok := om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	v, loaded = om.LoadAndDelete("foo-0")
	assert.False(t, loaded)
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())
	_ = v

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndDelete("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())
	_ = v

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())
}

func TestLoadAndLdelete(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	v, loaded := om.LoadAndLdelete("foo-0")
	assert.False(t, loaded)
	_ = v
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("foo-0", foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	var ok bool
	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "A", N: 1})

	ent := om.Front()
	assert.Equal(t, om.Back(), ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndLdelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	_ = v

	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("foo-0", foo{S: "AA", N: 11})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AA", N: 11})

	ent = om.Front()
	assert.Equal(t, om.Back(), ent)
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AA", N: 11})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndLdelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AA", N: 11})

	v, loaded = om.LoadAndLdelete("foo-0")
	assert.Equal(t, om.Len(), 0)
	assert.False(t, loaded)

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	_ = v

	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("foo-1", foo{S: "B", N: 2})
	assert.Equal(t, om.Len(), 1)

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, om.Back(), ent)
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	assert.Nil(t, ent.Next())

	om.Store("foo-0", foo{S: "AAA", N: 111})
	assert.Equal(t, om.Len(), 2)
	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AAA", N: 111})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AAA", N: 111})
	assert.Nil(t, ent.Next())

	om.Store("foo-2", foo{S: "C", N: 3})
	assert.Equal(t, om.Len(), 3)

	v, ok = om.Load("foo-0")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "AAA", N: 111})
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AAA", N: 111})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndLdelete("foo-0")
	assert.Equal(t, om.Len(), 2)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AAA", N: 111})

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "C", N: 3})

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-2")
	assert.Equal(t, ent.Value(), foo{S: "C", N: 3})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndLdelete("foo-2")
	assert.Equal(t, om.Len(), 1)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "C", N: 3})

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	v, ok = om.Load("foo-1")
	assert.True(t, ok)
	assert.Equal(t, v, foo{S: "B", N: 2})
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	assert.Nil(t, ent.Next())

	v, loaded = om.LoadAndLdelete("foo-1")
	assert.Equal(t, om.Len(), 0)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "B", N: 2})

	v, ok = om.Load("foo-0")
	assert.False(t, ok)
	v, ok = om.Load("foo-1")
	assert.False(t, ok)
	v, ok = om.Load("foo-2")
	assert.False(t, ok)

	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())
}

func TestLoadOrStoreFunc(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	v, loaded, err := om.LoadOrStoreFunc("foo-0", func() (foo, error) {
		return foo{S: "A", N: 1}, nil
	})
	assert.Nil(t, err)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, loaded = om.Load("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	ent = ent.Next()
	assert.Nil(t, ent)

	v, loaded, err = om.LoadOrStoreFunc("foo-0", func() (foo, error) {
		return foo{S: "AA", N: 11}, nil
	})
	assert.Nil(t, err)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	assert.Equal(t, om.Len(), 1)

	v, loaded = om.Load("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "A", N: 1})
	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "A", N: 1})
	ent = ent.Next()
	assert.Nil(t, ent)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, loaded, err = om.LoadOrStoreFunc("foo-0", func() (foo, error) {
		return foo{S: "AAA", N: 111}, nil
	})
	assert.Nil(t, err)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "AAA", N: 111})
	assert.Equal(t, om.Len(), 1)

	v, loaded = om.Load("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AAA", N: 111})
	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AAA", N: 111})
	ent = ent.Next()
	assert.Nil(t, ent)

	om.Ldelete("foo-0")
	assert.Equal(t, om.Len(), 0)

	v, loaded, err = om.LoadOrStoreFunc("foo-0", func() (foo, error) {
		return foo{S: "AAAA", N: 1111}, nil
	})
	assert.Nil(t, err)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "AAAA", N: 1111})
	assert.Equal(t, om.Len(), 1)

	v, loaded = om.Load("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AAAA", N: 1111})
	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AAAA", N: 1111})
	ent = ent.Next()
	assert.Nil(t, ent)

	v, loaded, err = om.LoadOrStoreFunc("foo-1", func() (foo, error) {
		return foo{S: "B", N: 2}, nil
	})
	assert.Nil(t, err)
	assert.False(t, loaded)
	assert.Equal(t, v, foo{S: "B", N: 2})
	assert.Equal(t, om.Len(), 2)

	v, loaded = om.Load("foo-0")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "AAAA", N: 1111})
	v, loaded = om.Load("foo-1")
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "B", N: 2})
	ent = om.Front()
	assert.Equal(t, ent.Key(), "foo-0")
	assert.Equal(t, ent.Value(), foo{S: "AAAA", N: 1111})
	ent = ent.Next()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	ent = ent.Next()
	assert.Nil(t, ent)
}

func TestLoadOrStoreFunc_StoreFuncCauseError(t *testing.T) {
	om := orderedmap.New[string, foo]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	v, loaded, err := om.LoadOrStoreFunc("foo-1", func() (foo, error) {
		return foo{}, fmt.Errorf("error")
	})
	assert.Equal(t, err.Error(), "error")
	assert.False(t, loaded)
	assert.Equal(t, v, foo{})
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("foo-1", foo{S: "B", N: 2})

	v, loaded, err = om.LoadOrStoreFunc("foo-1", func() (foo, error) {
		return foo{}, fmt.Errorf("error")
	})
	assert.Nil(t, err)
	assert.True(t, loaded)
	assert.Equal(t, v, foo{S: "B", N: 2})
	ent := om.Front()
	assert.Equal(t, ent.Key(), "foo-1")
	assert.Equal(t, ent.Value(), foo{S: "B", N: 2})
	assert.Nil(t, ent.Next())

	om.Ldelete("foo-1")

	v, loaded, err = om.LoadOrStoreFunc("foo-1", func() (foo, error) {
		return foo{}, fmt.Errorf("error")
	})
	assert.Equal(t, err.Error(), "error")
	assert.False(t, loaded)
	assert.Equal(t, v, foo{})
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())
}

func TestString(t *testing.T) {
	om := orderedmap.New[string, string]()
	om.Store("key-0", "value-0")
	om.Store("key-1", "value-1")

	assert.Equal(t, fmt.Sprintf("%v", om), "Map[key-0:value-0 key-1:value-1]")
}

func TestFrontAndDelete(t *testing.T) {
	om := orderedmap.New[string, string]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.FrontAndDelete())

	om.Store("key-0", "value-0")
	om.Store("key-1", "value-1")
	om.Store("key-2", "value-2")
	assert.Equal(t, om.Len(), 3)

	ent := om.FrontAndDelete()
	assert.Equal(t, om.Len(), 2)
	assert.Equal(t, ent.Key(), "key-0")
	assert.Equal(t, ent.Value(), "value-0")

	v, b := om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "value-1")
	assert.True(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "value-2")
	assert.True(t, b)

	ent = om.FrontAndDelete()
	assert.Equal(t, om.Len(), 1)
	assert.Equal(t, ent.Key(), "key-1")
	assert.Equal(t, ent.Value(), "value-1")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "value-2")
	assert.True(t, b)

	ent = om.FrontAndDelete()
	assert.Equal(t, om.Len(), 0)
	assert.Equal(t, ent.Key(), "key-2")
	assert.Equal(t, ent.Value(), "value-2")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.FrontAndDelete()
	assert.Nil(t, ent)
}

func TestFrontAndLdelete(t *testing.T) {
	om := orderedmap.New[string, string]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("key-0", "value-0")
	om.Store("key-1", "value-1")
	om.Store("key-2", "value-2")
	assert.Equal(t, om.Len(), 3)

	ent := om.FrontAndLdelete()
	assert.Equal(t, om.Len(), 2)
	assert.Equal(t, ent.Key(), "key-0")
	assert.Equal(t, ent.Value(), "value-0")

	v, b := om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "value-1")
	assert.True(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "value-2")
	assert.True(t, b)

	ent = om.FrontAndLdelete()
	assert.Equal(t, om.Len(), 1)
	assert.Equal(t, ent.Key(), "key-1")
	assert.Equal(t, ent.Value(), "value-1")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "value-2")
	assert.True(t, b)

	ent = om.FrontAndLdelete()
	assert.Equal(t, om.Len(), 0)
	assert.Equal(t, ent.Key(), "key-2")
	assert.Equal(t, ent.Value(), "value-2")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.FrontAndLdelete()
	assert.Nil(t, ent)
}

func TestBackAndDelete(t *testing.T) {
	om := orderedmap.New[string, string]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("key-0", "value-0")
	om.Store("key-1", "value-1")
	om.Store("key-2", "value-2")
	assert.Equal(t, om.Len(), 3)

	ent := om.BackAndDelete()
	assert.Equal(t, om.Len(), 2)
	assert.Equal(t, ent.Key(), "key-2")
	assert.Equal(t, ent.Value(), "value-2")

	v, b := om.Load("key-0")
	assert.Equal(t, v, "value-0")
	assert.True(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "value-1")
	assert.True(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndDelete()
	assert.Equal(t, om.Len(), 1)
	assert.Equal(t, ent.Key(), "key-1")
	assert.Equal(t, ent.Value(), "value-1")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "value-0")
	assert.True(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndDelete()
	assert.Equal(t, om.Len(), 0)
	assert.Equal(t, ent.Key(), "key-0")
	assert.Equal(t, ent.Value(), "value-0")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndDelete()
	assert.Nil(t, ent)
}

func TestBackAndLdelete(t *testing.T) {
	om := orderedmap.New[string, string]()
	assert.Equal(t, om.Len(), 0)
	assert.Nil(t, om.Front())
	assert.Nil(t, om.Back())

	om.Store("key-0", "value-0")
	om.Store("key-1", "value-1")
	om.Store("key-2", "value-2")
	assert.Equal(t, om.Len(), 3)

	ent := om.BackAndLdelete()
	assert.Equal(t, om.Len(), 2)
	assert.Equal(t, ent.Key(), "key-2")
	assert.Equal(t, ent.Value(), "value-2")

	v, b := om.Load("key-0")
	assert.Equal(t, v, "value-0")
	assert.True(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "value-1")
	assert.True(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndLdelete()
	assert.Equal(t, om.Len(), 1)
	assert.Equal(t, ent.Key(), "key-1")
	assert.Equal(t, ent.Value(), "value-1")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "value-0")
	assert.True(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndLdelete()
	assert.Equal(t, om.Len(), 0)
	assert.Equal(t, ent.Key(), "key-0")
	assert.Equal(t, ent.Value(), "value-0")

	v, b = om.Load("key-0")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-1")
	assert.Equal(t, v, "")
	assert.False(t, b)
	v, b = om.Load("key-2")
	assert.Equal(t, v, "")
	assert.False(t, b)

	ent = om.BackAndLdelete()
	assert.Nil(t, ent)
}
