// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

// Package orderedmap provides OrderedMap type which is a map presering the
// order of key insertions.
package orderedmap

// OrderedMap is a struct which represents a map similar with Go standard map,
// or sync.Map, but preserves the order in which keys were inserted.
//
// This map has same methods with sync.Map except CompareAndDelete and 
// CompareAndSwap.
// Its Range method processes a key and a value of each map entry, and the
// processing order is same with the order of key insertions.
// And this map also has a method Iter, which iterates this map entries, and
// the order is also same with the order of key insertions.
//
// Moreover, this map also supported generics, but does not support
// synchronization.
type OrderedMap[K comparable, V any] struct {
	m    map[K](*Entry[K, V])
	head *Entry[K, V]
	last *Entry[K, V]
}

// Entry is a struct which is a map element and holds a pair of key and value.
// This struct also has methods: Next and Prev which moves next or previous entties
// sequencially.
type Entry[K comparable, V any] struct {
	key   K
	value V
	prev  *Entry[K, V]
	next  *Entry[K, V]
}

// New is a function which creates a new ordered map, which is ampty.
func New[K comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{m: make(map[K](*Entry[K, V]))}
}

// Len is a method which returns the number of entries in this map.
func (om *OrderedMap[K, V]) Len() int {
	return len(om.m)
}

// Store is a method which sets a value for a key
func (om *OrderedMap[K, V]) Store(key K, value V) {
	ent, exists := om.m[key]
	if exists {
		ent.value = value
		return
	}

	ent = &Entry[K, V]{key: key, value: value}

	if len(om.m) == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	return
}

// Swap is a method which sets a value for a key. If the key was present, this
// map returns the previous value and the loaded flag which is set to true.
func (om *OrderedMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	ent, exists := om.m[key]
	if exists {
		loaded = true
		previous = ent.value
		ent.value = value
		return
	}

	ent = &Entry[K, V]{key: key, value: value}

	if len(om.m) == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent
	return
}

// Load is a method which returns a value stored in this map for a key.
// If no value was found for a key, the ok result is false.
func (om *OrderedMap[K, V]) Load(key K) (value V, ok bool) {
	ent, ok := om.m[key]
	if ok {
		value = ent.value
	}
	return
}

// LoadOrStore is a method which returns a value for a key if presents,
// otherwise stores and returns a given value.
// The loaded flag is true if the value was loaded, false if stored.
func (om *OrderedMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	ent, exists := om.m[key]
	if exists {
		actual = ent.value
		loaded = true
		return
	}

	actual = value
	ent = &Entry[K, V]{key: key, value: value}

	if len(om.m) == 0 {
		om.head = ent
		om.last = ent
		om.m[key] = ent
		return
	}

	ent.prev = om.last
	om.last.next = ent
	om.last = ent
	om.m[key] = ent

	return
}

// Delete is a method which deletes a value for a key.
func (om *OrderedMap[K, V]) Delete(key K) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	delete(om.m, key)

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

// LoadAndDelete is a method which deletes a value for a key, and returns the
// previous value if any.
// The loaded flag is true if the key was present.
func (om *OrderedMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	ent, exists := om.m[key]
	if !exists {
		return
	}

	delete(om.m, key)

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
func (om *OrderedMap[K, V]) Range(fn func(key K, value V) bool) {
	for entry := om.head; entry != nil; entry = entry.next {
		if !fn(entry.key, entry.value) {
			break
		}
	}
}

// Front is a method which returns the head entry of this map.
func (om *OrderedMap[K, V]) Front() *Entry[K, V] {
	return om.head
}

// Back is a method which returns the last entry of this map.
func (om *OrderedMap[K, V]) Back() *Entry[K, V] {
	return om.last
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
