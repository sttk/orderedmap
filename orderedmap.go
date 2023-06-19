// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

// Package orderedmap provides Map type which is a map presering the
// order of key insertions.
//
// # Usage
//
// To create an ordered map is as follows:
//
//	om := orderedmap.New[string, string]()
//
// To add a map entry is as follows:
//
//	om.Store("foo", "hoge")
//	prev, swapped := om.Swap("bar", "fuga")
//	actual, loaded := om.LoadOrStore("baz", "fuga")
//
// To get a value for a key is as follows:
//
//	om.Load("foo")
//
// To delete a map entry is as follows:
//
//	om.Delete("bar")
//	v,, deleted := om.LoadAndDelete("baz")
//
// To delete a map entry logically is as follows:
//
//	om.Ldelete("bar")
//	v,, deleted := om.LoadAndLdelete("baz")
//
// To iterate map entries is as folLows. The order is same with key insertions:
//
//	om.Range(func(k, v) bool { ... })
//	for ent := om.Front(); ent != nil; ent = ent.Next() {
//	    k := ent.Key(); v : = ent.Value(); ...
//	}
//	for ent := om.Back(); ent != nil; ent = ent.Prev() {
//	    k := ent.Key(); v : = ent.Value(); ...
//	}
package orderedmap

import (
	"fmt"
	"strings"
)

// Map is a struct which represents a map similar with Go standard map,
// or sync.Map, but preserves the order in which keys were inserted.
//
// This map has same methods with sync.Map except CompareAndDelete and
// CompareAndSwap. (But not support concurrent use.)
// Its Range method processes a key and a value of each map entry, and the
// processing order is same with the order of key insertions.
// And this map also has methods: Front and Back, which iterate this map
// entries in the order of key insertions and in that reverse order.
type Map[K comparable, V any] struct {
	m    map[K](*Entry[K, V])
	head *Entry[K, V]
	last *Entry[K, V]
	len  int
}

// Entry is a struct which is a map element and holds a pair of key and value.
// This struct also has methods: Next and Prev which moves next or previous entties
// sequencially.
type Entry[K comparable, V any] struct {
	key     K
	value   V
	prev    *Entry[K, V]
	next    *Entry[K, V]
	deleted bool
}

// New is a function which creates a new ordered map, which is ampty.
func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{m: make(map[K](*Entry[K, V]))}
}

// Len is a method which returns the number of entries in this map.
func (om *Map[K, V]) Len() int {
	return om.len
}

// Store is a method which sets a value for a key
func (om *Map[K, V]) Store(key K, value V) {
	ent, exists := om.m[key]
	if exists {
		if !ent.deleted {
			ent.value = value
			return
		}
		ent.value = value
		ent.deleted = false
	} else {
		ent = &Entry[K, V]{key: key, value: value}
	}

	if om.len == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		om.len = 1
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	om.len++
	return
}

// Swap is a method which sets a value for a key. If the key was present, this
// map returns the previous value and the loaded flag which is set to true.
func (om *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	ent, exists := om.m[key]
	if exists {
		if !ent.deleted {
			loaded = true
			previous = ent.value
			ent.value = value
			return
		}
		ent.deleted = false
		ent.value = value
	} else {
		ent = &Entry[K, V]{key: key, value: value}
	}

	if om.len == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		om.len = 1
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	om.len++
	return
}

// Load is a method which returns a value stored in this map for a key.
// If no value was found for a key, the ok result is false.
func (om *Map[K, V]) Load(key K) (value V, ok bool) {
	ent, exists := om.m[key]
	if exists {
		if !ent.deleted {
			value = ent.value
			ok = true
		}
	}
	return
}

// LoadOrStore is a method which returns a value for a key if presents,
// otherwise stores and returns a given value.
// The loaded flag is true if the value was loaded, false if stored.
func (om *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	ent, exists := om.m[key]
	if exists {
		if !ent.deleted {
			actual = ent.value
			loaded = true
			return
		}
		ent.deleted = false
		ent.value = value
	} else {
		ent = &Entry[K, V]{key: key, value: value}
	}

	actual = value

	if om.len == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		om.len = 1
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	om.len++

	return
}

// LoadOrStoreFunc is a method which returns a value for a key if presents,
// otherwise executes a give function, then stores and returns the result
// value.
// The loaded flag is true if the value was loaded, false if stored.
func (om *Map[K, V]) LoadOrStoreFunc(
	key K,
	fn func() (V, error),
) (actual V, loaded bool, err error) {
	ent, exists := om.m[key]
	if exists {
		if !ent.deleted {
			actual = ent.value
			loaded = true
			return
		}
		ent.deleted = false
		v, e := fn()
		if e != nil {
			err = e
			return
		}
		actual = v
		ent.value = actual
	} else {
		v, e := fn()
		if e != nil {
			err = e
			return
		}
		actual = v
		ent = &Entry[K, V]{key: key, value: actual}
	}

	if om.len == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		om.len = 1
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	om.len++

	return
}

// Delete is a method which deletes a value for a key.
func (om *Map[K, V]) Delete(key K) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	delete(om.m, key)

	if ent.deleted {
		return
	}
	om.len--

	if ent.prev != nil {
		ent.prev.next = ent.next
	} else {
		om.head = ent.next
	}

	if ent.next != nil {
		ent.next.prev = ent.prev
	} else {
		om.last = ent.prev
	}
}

// Ldelete is a method which logically deletes a value for a key.
func (om *Map[K, V]) Ldelete(key K) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	if ent.deleted {
		return
	}
	ent.deleted = true
	om.len--

	if ent.prev != nil {
		ent.prev.next = ent.next
	} else {
		om.head = ent.next
	}

	if ent.next != nil {
		ent.next.prev = ent.prev
	} else {
		om.last = ent.prev
	}

	ent.next = nil
	ent.prev = nil
}

// LoadAndDelete is a method which deletes a value for a key, and returns the
// previous value if any.
// The loaded flag is true if the key was present.
func (om *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	delete(om.m, key)

	if ent.deleted {
		return
	}
	om.len--

	if ent.prev != nil {
		ent.prev.next = ent.next
	} else {
		om.head = ent.next
	}

	if ent.next != nil {
		ent.next.prev = ent.prev
	} else {
		om.last = ent.prev
	}

	value = ent.value
	loaded = true
	return
}

// LoadAndLdelete is a method which logically deletes a value for a key, and
// returns the previous value if any.
// The loaded flag is true if the key was present.
func (om *Map[K, V]) LoadAndLdelete(key K) (value V, loaded bool) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	if ent.deleted {
		return
	}
	ent.deleted = true
	om.len--

	if ent.prev != nil {
		ent.prev.next = ent.next
	} else {
		om.head = ent.next
	}

	if ent.next != nil {
		ent.next.prev = ent.prev
	} else {
		om.last = ent.prev
	}

	value = ent.value
	loaded = true
	return
}

// Range is a method which calls the specified function: fn sequentially for
// each key and value in this map.
// If fn returns false, this method stops the iteration.
func (om *Map[K, V]) Range(fn func(key K, value V) bool) {
	for entry := om.head; entry != nil; entry = entry.next {
		if !fn(entry.key, entry.value) {
			break
		}
	}
}

// Front is a method which returns the head entry of this map.
func (om *Map[K, V]) Front() *Entry[K, V] {
	return om.head
}

// Back is a method which returns the last entry of this map.
func (om *Map[K, V]) Back() *Entry[K, V] {
	return om.last
}

// String is a method which returns a string of the content of this map.
func (om Map[K, V]) String() string {
	var b strings.Builder
	b.WriteString("Map[")
	ent := om.Front()
	if ent != nil {
		b.WriteString(fmt.Sprintf("%v:%v", ent.Key(), ent.Value()))
	}
	for ent = ent.Next(); ent != nil; ent = ent.Next() {
		b.WriteString(fmt.Sprintf(" %v:%v", ent.Key(), ent.Value()))
	}
	b.WriteString("]")
	return b.String()
}

// Prev is a method which returns the previous entry of this entry.
// If this entry is a head entry of an ordered map, the returned value is nil.
func (ent *Entry[K, V]) Prev() *Entry[K, V] {
	return ent.prev
}

// Next is a method which returns the next entry of this entry.
// If this entry is a last entry of an ordered map, the returned value is nil.
func (ent *Entry[K, V]) Next() *Entry[K, V] {
	return ent.next
}

// Key is a method which returns the key of this entry.
func (ent *Entry[K, V]) Key() K {
	return ent.key
}

// Value is a method which returns the value of this entry.
func (ent *Entry[K, V]) Value() V {
	return ent.value
}
