package orderedmap_test

import (
	"fmt"

	"github.com/sttk/orderedmap"
)

func ExampleMap_MarshalJSON() {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	b, e := om.MarshalJSON()
	fmt.Printf("json = %s\n", string(b))
	fmt.Printf("e = %v\n", e)
	// Output:
	// json = {"foo":"bar","baz":"qux"}
	// e = <nil>
}

func ExampleMap_UnmarshalJSON() {
	om := orderedmap.New[string, string]()
	b := []byte(`{"foo":"bar","baz":"qux"}`)

	e := om.UnmarshalJSON(b)
	fmt.Printf("om = %v\n", om)
	fmt.Printf("e = %v\n", e)
	// Output:
	// om = Map[foo:bar baz:qux]
	// e = <nil>
}
