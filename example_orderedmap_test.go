package orderedmap_test

import (
	"fmt"

	"github.com/sttk/orderedmap"
)

func ExampleNew() {
	om := orderedmap.New[string, string]()
	fmt.Printf("om = %v\n", om)
	// Output:
	// om = Map[]
}

func ExampleMap_Len() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	fmt.Printf("orderedmap's length = %v\n", om.Len())
	om.Store("baz", "qux")
	fmt.Printf("orderedmap's length = %v\n", om.Len())
	// Output:
	// orderedmap's length = 1
	// orderedmap's length = 2
}

func ExampleMap_Store() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	fmt.Printf("om = %v\n", om)
	om.Store("baz", "qux")
	fmt.Printf("om = %v\n", om)
	// Output:
	// om = Map[foo:bar]
	// om = Map[foo:bar baz:qux]
}

func ExampleMap_Swap() {
	om := orderedmap.New[string, string]()
	prev, loaded := om.Swap("foo", "bar")
	fmt.Printf("prev = %v, loaded = %t\n", prev, loaded)
	prev, loaded = om.Swap("foo", "BAZ")
	fmt.Printf("prev = %v, loaded = %t\n", prev, loaded)
	// Output:
	// prev = , loaded = false
	// prev = bar, loaded = true
}

func ExampleMap_Load() {
	om := orderedmap.New[string, string]()
	value, ok := om.Load("foo")
	fmt.Printf("value = %v, ok = %t\n", value, ok)
	om.Store("foo", "bar")
	value, ok = om.Load("foo")
	fmt.Printf("value = %v, ok = %t\n", value, ok)
	// Output:
	// value = , ok = false
	// value = bar, ok = true
}

func ExampleMap_LoadOrStore() {
	om := orderedmap.New[string, string]()
	actual, loaded := om.LoadOrStore("foo", "bar")
	fmt.Printf("actual = %v, loaded = %t\n", actual, loaded)
	actual, loaded = om.LoadOrStore("foo", "BAZ")
	fmt.Printf("actual = %v, loaded = %t\n", actual, loaded)
	// Output:
	// actual = bar, loaded = false
	// actual = bar, loaded = true
}

func ExampleMap_LoadOrStoreFunc() {
	om := orderedmap.New[string, string]()
	actual, loaded, err := om.LoadOrStoreFunc("foo", func() (string, error) {
		return "bar", nil
	})
	fmt.Printf("actual = %v, loaded = %t, err = %v\n", actual, loaded, err)
	actual, loaded, err = om.LoadOrStoreFunc("foo", func() (string, error) {
		return "BAZ", nil
	})
	fmt.Printf("actual = %v, loaded = %t, err = %v\n", actual, loaded, err)
	// Output:
	// actual = bar, loaded = false, err = <nil>
	// actual = bar, loaded = true, err = <nil>
}

func ExampleMap_Delete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	om.Delete("foo")
	fmt.Printf("om = %v\n", om)
	// Output:
	// om = Map[baz:qux]
}

func ExampleMap_Ldelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	om.Ldelete("foo")
	fmt.Printf("om = %v\n", om)
	// Output:
	// om = Map[baz:qux]
}

func ExampleMap_LoadAndDelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	value, loaded := om.LoadAndDelete("foo")
	fmt.Printf("value = %v, loaded = %t\n", value, loaded)
	fmt.Printf("om = %v\n", om)
	// Output:
	// value = bar, loaded = true
	// om = Map[baz:qux]
}

func ExampleMap_LoadAndLdelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	value, loaded := om.LoadAndLdelete("foo")
	fmt.Printf("value = %v, loaded = %t\n", value, loaded)
	fmt.Printf("om = %v\n", om)
	// Output:
	// value = bar, loaded = true
	// om = Map[baz:qux]
}

func ExampleMap_FrontAndDelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.FrontAndDelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.FrontAndDelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.FrontAndDelete()
	fmt.Printf("entry = %v\n", entry)
	fmt.Printf("om = %v\n", om)
	// Output:
	// key = foo, value = bar
	// key = baz, value = qux
	// entry = <nil>
	// om = Map[]
}

func ExampleMap_FrontAndLdelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.FrontAndLdelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.FrontAndLdelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.FrontAndLdelete()
	fmt.Printf("entry = %v\n", entry)
	fmt.Printf("om = %v\n", om)
	// Output:
	// key = foo, value = bar
	// key = baz, value = qux
	// entry = <nil>
	// om = Map[]
}

func ExampleMap_BackAndDelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.BackAndDelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.BackAndDelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.BackAndDelete()
	fmt.Printf("entry = %v\n", entry)
	fmt.Printf("om = %v\n", om)
	// Output:
	// key = baz, value = qux
	// key = foo, value = bar
	// entry = <nil>
	// om = Map[]
}

func ExampleMap_BackAndLdelete() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.BackAndLdelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.BackAndLdelete()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = om.BackAndLdelete()
	fmt.Printf("entry = %v\n", entry)
	fmt.Printf("om = %v\n", om)
	// Output:
	// key = baz, value = qux
	// key = foo, value = bar
	// entry = <nil>
	// om = Map[]
}

func ExampleMap_Range() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")
	om.Store("quux", "corge")

	om.Range(func(key, value string) bool {
		fmt.Printf("key = %v, value = %v\n", key, value)
		if key == "baz" {
			return false
		}
		return true
	})
	// Output:
	// key = foo, value = bar
	// key = baz, value = qux
}

func ExampleMap_Front() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.Front()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	// Output:
	// key = foo, value = bar
}

func ExampleMap_Back() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.Back()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	// Output:
	// key = baz, value = qux
}

func ExampleMap_String() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")
	fmt.Printf("om = %v\n", om)
	// Output:
	// om = Map[foo:bar baz:qux]
}

func ExampleEntry_Prev() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.Back()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = entry.Prev()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = entry.Prev()
	fmt.Printf("entry = %v\n", entry)
	// Output:
	// key = baz, value = qux
	// key = foo, value = bar
	// entry = <nil>
}

func ExampleEntry_Next() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	entry := om.Front()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = entry.Next()
	fmt.Printf("key = %v, value = %v\n", entry.Key(), entry.Value())
	entry = entry.Next()
	fmt.Printf("entry = %v\n", entry)
	// Output:
	// key = foo, value = bar
	// key = baz, value = qux
	// entry = <nil>
}

func ExampleEntry_Key() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")

	entry := om.Front()
	fmt.Printf("key = %v\n", entry.Key())
	// Output:
	// key = foo
}

func ExampleEntry_Value() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")

	entry := om.Front()
	fmt.Printf("value = %v\n", entry.Value())
	// Output:
	// value = bar
}
