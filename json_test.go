package orderedmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/orderedmap"
)

func TestMarshalJSON_mapIsEmpty(t *testing.T) {
	om := orderedmap.New[string, string]()

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte("{}"))
}

func TestMarshalJSON_mapHasOneEntry(t *testing.T) {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":"bar"}`))
}

func TestMarshalJSON_mapHasMultipleEntries(t *testing.T) {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("baz", "qux")

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":"bar","baz":"qux"}`))
}

func TestMarshalJSON_variousKeyAndValueTypes_string(t *testing.T) {
	om := orderedmap.New[string, string]()
	om.Store("foo", "bar")
	om.Store("", "qux")

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":"bar","":"qux"}`))
}

func TestMarshalJSON_variousKeyAndValueTypes_bool(t *testing.T) {
	om := orderedmap.New[bool, bool]()
	om.Store(true, true)
	om.Store(false, false)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"true":true,"false":false}`))
}

func TestMarshalJSON_variousKeyAndValueTypes_int(t *testing.T) {
	om := orderedmap.New[int, int]()
	om.Store(12, 34)
	om.Store(0, 0)
	om.Store(-56, -78)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0,"-56":-78}`))
}

func TestMarshalJSON_variousKeyAndValueTypes_int8(t *testing.T) {
	om := orderedmap.New[int8, int8]()
	om.Store(int8(12), int8(34))
	om.Store(int8(0), int8(0))
	om.Store(int8(-56), int8(-78))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0,"-56":-78}`))
}

func TestMarshalJSON_variousKeyAndValueTypes_int16(t *testing.T) {
	om := orderedmap.New[int16, int16]()
	om.Store(int16(12), int16(34))
	om.Store(int16(0), int16(0))
	om.Store(int16(-56), int16(-78))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0,"-56":-78}`))
}

func TestMarshalJSON_variousKeyTypes_int32(t *testing.T) {
	om := orderedmap.New[int32, int32]()
	om.Store(int32(12), int32(34))
	om.Store(int32(0), int32(0))
	om.Store(int32(-56), int32(-78))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0,"-56":-78}`))
}

func TestMarshalJSON_variousKeyTypes_int64(t *testing.T) {
	om := orderedmap.New[int64, int64]()
	om.Store(int64(12), int64(34))
	om.Store(int64(0), int64(0))
	om.Store(int64(-56), int64(-78))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0,"-56":-78}`))
}

func TestMarshalJSON_variousKeyTypes_uint(t *testing.T) {
	om := orderedmap.New[uint, uint]()
	om.Store(uint(12), uint(34))
	om.Store(uint(0), uint(0))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte("{\"12\":34,\"0\":0}"))
}

func TestMarshalJSON_variousKeyTypes_uint8(t *testing.T) {
	om := orderedmap.New[uint8, uint8]()
	om.Store(uint8(12), uint8(34))
	om.Store(uint8(0), uint8(0))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte("{\"12\":34,\"0\":0}"))
}

func TestMarshalJSON_variousKeyTypes_uint16(t *testing.T) {
	om := orderedmap.New[uint16, uint16]()
	om.Store(uint16(12), uint16(34))
	om.Store(uint16(0), uint16(0))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte("{\"12\":34,\"0\":0}"))
}

func TestMarshalJSON_variousKeyTypes_uint32(t *testing.T) {
	om := orderedmap.New[uint32, uint32]()
	om.Store(uint32(12), uint32(34))
	om.Store(uint32(0), uint32(0))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte("{\"12\":34,\"0\":0}"))
}

func TestMarshalJSON_variousKeyTypes_uint64(t *testing.T) {
	om := orderedmap.New[uint64, uint64]()
	om.Store(uint64(12), uint64(34))
	om.Store(uint64(0), uint64(0))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"12":34,"0":0}`))
}

func TestMarshalJSON_variousKeyTypes_float32(t *testing.T) {
	om := orderedmap.New[float32, float32]()
	om.Store(float32(1.23), float32(4.56))
	om.Store(float32(0.0), float32(0.0))
	om.Store(float32(-1.23), float32(-4.56))

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"1.23":4.56,"0":0,"-1.23":-4.56}`))
}

func TestMarshalJSON_variousKeyTypes_float64(t *testing.T) {
	om := orderedmap.New[float64, float64]()
	om.Store(1.23, 4.56)
	om.Store(0.0, 0.0)
	om.Store(-1.23, -4.56)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"1.23":4.56,"0":0,"-1.23":-4.56}`))
}

func TestMarshalJSON_variousKeyTypes_complex64_unsupportedError(t *testing.T) {
	om := orderedmap.New[complex64, string]()
	om.Store(complex(2, 3), "foo")
	om.Store(2+3i, "bar")
	om.Store(2, "baz")
	om.Store(3i, "qux")
	om.Store(0, "quux")

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported key type: complex64")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousValueTypes_complex64_unsupportedError(t *testing.T) {
	om := orderedmap.New[string, complex64]()
	om.Store("foo", complex(2, 3))
	om.Store("bar", 2+3i)
	om.Store("baz", 2)
	om.Store("qux", 3i)
	om.Store("quux", 0)

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported type: complex64")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousKeyTypes_complex128_unsupportedError(t *testing.T) {
	om := orderedmap.New[complex128, string]()
	om.Store(complex(2, 3), "foo")
	om.Store(2+3i, "bar")
	om.Store(2, "baz")
	om.Store(3i, "qux")
	om.Store(0, "quux")

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported key type: complex128")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousValueTypes_complex128_unsupportedError(t *testing.T) {
	om := orderedmap.New[string, complex128]()
	om.Store("foo", complex(2, 3))
	om.Store("bar", 2+3i)
	om.Store("baz", 2)
	om.Store("qux", 3i)
	om.Store("quux", 0)

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported type: complex128")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousKeyTypes_channel_unsupportedError(t *testing.T) {
	ch := make(chan int)
	om := orderedmap.New[chan int, string]()
	om.Store(ch, "foo")

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported key type: chan int")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousValueTypes_channel_unsupportedError(t *testing.T) {
	ch := make(chan int)
	om := orderedmap.New[string, chan int]()
	om.Store("foo", ch)

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported type: chan int")
	assert.Equal(t, b, []byte(nil))
}

type A interface{ B() error }
type A1 struct {
	S, t string
	U    string `json:"u"`
}

func (a A1) B() error { return nil }

// COMPILE ERROR: A to satisfy comparable requires go1.20 or later (-lang was set to go1.18; check go.mod)
// func TestMarshalJSON_variousKeyTypes_interface_unsupportedError(t *testing.T) {
// 	k1 := A1{S: "aaa", t: "bbb", U: "ccc"}
// 	k2 := A1{S: "ggg", t: "hhh", U: "iii"}
//
// 	om := orderedmap.New[A, string]()
// 	om.Store(k1, "foo")
// 	om.Store(&k2, "bar")
//
// 	b, e := om.MarshalJSON()
// 	assert.Equal(t, e.Error(), "json: unsupported key type: orderedmap_test.A1")
// 	assert.Equal(t, b, []byte(nil))
// }

func TestMarshalJSON_variousValueTypes_interface(t *testing.T) {
	v1 := A1{S: "aaa", t: "bbb", U: "ccc"}
	v2 := A1{S: "ggg", t: "hhh", U: "iii"}

	om := orderedmap.New[string, A]()
	om.Store("foo", v1)
	om.Store("bar", &v2)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":{"S":"aaa","u":"ccc"},"bar":{"S":"ggg","u":"iii"}}`))
}

func TestMarshalJSON_variousKeyTypes_struct_unsupportedError(t *testing.T) {
	k1 := A1{S: "aaa", t: "bbb", U: "ccc"}
	k2 := A1{S: "ggg", t: "hhh", U: "iii"}

	om := orderedmap.New[A1, string]()
	om.Store(k1, "foo")
	om.Store(k2, "bar")

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported key type: orderedmap_test.A1")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousValueTypes_struct(t *testing.T) {
	v1 := A1{S: "aaa", t: "bbb", U: "ccc"}
	v2 := A1{S: "ggg", t: "hhh", U: "iii"}

	om := orderedmap.New[string, A1]()
	om.Store("foo", v1)
	om.Store("bar", v2)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":{"S":"aaa","u":"ccc"},"bar":{"S":"ggg","u":"iii"}}`))
}

// Cause a compile error
//func TestMarshalJSON_variousKeyTypes_array(t *testing.T) {
//	arr := []string{"a", "b"}
//
//	// []string does not satisfy comparable
//	om := orderedmap.New[[]string, string]()
//	om.Store(arr, "foo")
//}

func TestMarshalJSON_variousKeyTypes_stringPointer(t *testing.T) {
	om := orderedmap.New[*string, string]()
	s := "abc"
	om.Store(&s, "foo")
	om.Store(nil, "bar")

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"abc":"foo","null":"bar"}`))
}

func TestMarshalJSON_variousKeyTypes_intPointer(t *testing.T) {
	om := orderedmap.New[*int, string]()
	n := 123
	om.Store(&n, "foo")
	om.Store(nil, "bar")

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"123":"foo","null":"bar"}`))
}

func TestMarshalJSON_variousKeyTypes_structPointer(t *testing.T) {
	a1 := A1{S: "xxx", t: "yyy", U: "zzz"}

	om := orderedmap.New[*A1, string]()
	om.Store(&a1, "foo")

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported key type: *orderedmap_test.A1")
	assert.Equal(t, b, []byte(nil))
}

func TestMarshalJSON_variousValueTypes_structPointer(t *testing.T) {
	a1 := A1{S: "xxx", t: "yyy", U: "zzz"}

	om := orderedmap.New[string, *A1]()
	om.Store("foo", &a1)

	b, e := om.MarshalJSON()
	assert.Nil(t, e)
	assert.Equal(t, b, []byte(`{"foo":{"S":"xxx","u":"zzz"}}`))
}

// COMPILE ERROR: any to satisfy comparable requires go1.20 or later (-lang was set to go1.18; check go.mod)
// func TestMarshalJSON_keyTypeIsAny(t *testing.T) {
// 	om := orderedmap.New[any, string]()
// 	om.Store("foo", "bar")
//
// 	b, e := om.MarshalJSON()
// 	assert.Nil(t, e)
// 	assert.Equal(t, b, []byte(`{"foo":"bar"}`))
// }

// COMPILE ERROR: any to satisfy comparable requires go1.20 or later (-lang was set to go1.18; check go.mod)
// func TestMarshalJSON_keyTypeIsAny_butUnsuuportedType(t *testing.T) {
// 	bar := complex(3, 2)
//
// 	om := orderedmap.New[any, string]()
// 	om.Store("foo", "bar")
// 	om.Store(bar, "qux")
//
// 	b, e := om.MarshalJSON()
// 	assert.Equal(t, e.Error(), "json: unsupported key type: complex128")
// 	assert.Equal(t, b, []byte(nil))
// }

func TestMarshalJSON_valueTypeIsAny_butUnsuuportedType(t *testing.T) {
	qux := complex(3, 2)

	om := orderedmap.New[string, any]()
	om.Store("foo", "bar")
	om.Store("bar", qux)

	b, e := om.MarshalJSON()
	assert.Equal(t, e.Error(), "json: unsupported type: complex128")
	assert.Equal(t, b, []byte(nil))
}

func TestUnmarshalJSON_emptyObject(t *testing.T) {
	om0 := orderedmap.New[string, string]()

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte("{}"))

	om1 := orderedmap.New[string, string]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 0)
}

func TestUnmarshalJSON_notObject(t *testing.T) {
	bs := []byte("123")

	om1 := orderedmap.New[string, string]()
	err := om1.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "The input JSON does not start with '{' (offset:0)")
	assert.Equal(t, om1.Len(), 0)
}

func TestUnmarshalJSON_nil(t *testing.T) {
	om1 := orderedmap.New[string, string]()
	err := om1.UnmarshalJSON([]byte(nil))
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_string(t *testing.T) {
	om0 := orderedmap.New[string, string]()
	om0.Store("foo", "bar")

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"foo":"bar"}`))

	om1 := orderedmap.New[string, string]()
	err = om1.UnmarshalJSON(bs)

	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 1)

	v, b := om1.Load("foo")
	assert.True(t, b)
	assert.Equal(t, v, "bar")
}

func TestUnmarshalJSON_variousKeyAndValueTypes_stringPointer(t *testing.T) {
	k0 := "foo"
	v0 := "bar"
	om0 := orderedmap.New[*string, *string]()
	om0.Store(&k0, &v0)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"foo":"bar","null":null}`))

	om1 := orderedmap.New[*string, *string]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	ent := om1.Front()
	assert.Equal(t, *(ent.Key()), "foo")
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, "bar")

	ent = ent.Next()
	assert.Equal(t, ent.Key(), (*string)(nil))
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*string)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_bool(t *testing.T) {
	om0 := orderedmap.New[bool, bool]()
	om0.Store(true, true)
	om0.Store(false, false)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"true":true,"false":false}`))

	om1 := orderedmap.New[bool, bool]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(true)
	assert.True(t, b)
	assert.Equal(t, v, true)

	v, b = om1.Load(false)
	assert.True(t, b)
	assert.Equal(t, v, false)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_boolPointer(t *testing.T) {
	tt := true
	ff := false
	om0 := orderedmap.New[*bool, *bool]()
	om0.Store(&tt, &tt)
	om0.Store(&ff, &ff)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"true":true,"false":false,"null":null}`))

	om1 := orderedmap.New[*bool, *bool]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, true)

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, false)

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*bool)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int(t *testing.T) {
	om0 := orderedmap.New[int, int]()
	om0.Store(12, 34)
	om0.Store(0, 0)
	om0.Store(-56, -78)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)

	v, b = om1.Load(-56)
	assert.True(t, b)
	assert.Equal(t, v, -78)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_intPointer(t *testing.T) {
	k0, v0, k1, v1, k2, v2 := 12, 34, 0, 0, -56, -78
	om0 := orderedmap.New[*int, *int]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(&k2, &v2)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78,"null":null}`))

	om1 := orderedmap.New[*int, *int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 4)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, 34)

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, 0)

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, -78)

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*int)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int8(t *testing.T) {
	om0 := orderedmap.New[int8, int8]()
	om0.Store(int8(12), int8(34))
	om0.Store(int8(0), int8(0))
	om0.Store(int8(-56), int8(-78))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)

	v, b = om1.Load(-56)
	assert.True(t, b)
	assert.Equal(t, v, -78)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int8Pointer(t *testing.T) {
	k0, v0 := int8(12), int8(34)
	k1, v1 := int8(0), int8(0)
	k2, v2 := int8(-56), int8(-78)
	om0 := orderedmap.New[*int8, *int8]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(&k2, &v2)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78,"null":null}`))

	om1 := orderedmap.New[*int8, *int8]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 4)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int8(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int8(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int8(-78))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*int8)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int16(t *testing.T) {
	om0 := orderedmap.New[int16, int16]()
	om0.Store(int16(12), int16(34))
	om0.Store(int16(0), int16(0))
	om0.Store(int16(-56), int16(-78))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)

	v, b = om1.Load(-56)
	assert.True(t, b)
	assert.Equal(t, v, -78)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int16Pointer(t *testing.T) {
	k0, v0 := int16(12), int16(34)
	k1, v1 := int16(0), int16(0)
	k2, v2 := int16(-56), int16(-78)
	om0 := orderedmap.New[*int16, *int16]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(&k2, &v2)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78,"null":null}`))

	om1 := orderedmap.New[*int16, *int16]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 4)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int16(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int16(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int16(-78))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*int16)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int32(t *testing.T) {
	om0 := orderedmap.New[int32, int32]()
	om0.Store(int32(12), int32(34))
	om0.Store(int32(0), int32(0))
	om0.Store(int32(-56), int32(-78))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)

	v, b = om1.Load(-56)
	assert.True(t, b)
	assert.Equal(t, v, -78)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int32Pointer(t *testing.T) {
	k0, v0 := int32(12), int32(34)
	k1, v1 := int32(0), int32(0)
	k2, v2 := int32(-56), int32(-78)
	om0 := orderedmap.New[*int32, *int32]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(&k2, &v2)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78,"null":null}`))

	om1 := orderedmap.New[*int32, *int32]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 4)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int32(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int32(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int32(-78))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*int32)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int64(t *testing.T) {
	om0 := orderedmap.New[int64, int64]()
	om0.Store(int64(12), int64(34))
	om0.Store(int64(0), int64(0))
	om0.Store(int64(-56), int64(-78))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)

	v, b = om1.Load(-56)
	assert.True(t, b)
	assert.Equal(t, v, -78)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_int64Pointer(t *testing.T) {
	k0, v0 := int64(12), int64(34)
	k1, v1 := int64(0), int64(0)
	k2, v2 := int64(-56), int64(-78)
	om0 := orderedmap.New[*int64, *int64]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(&k2, &v2)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"-56":-78,"null":null}`))

	om1 := orderedmap.New[*int64, *int64]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 4)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int64(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int64(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, int64(-78))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*int64)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint(t *testing.T) {
	om0 := orderedmap.New[uint, uint]()
	om0.Store(uint(12), uint(34))
	om0.Store(uint(0), uint(0))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uintPointer(t *testing.T) {
	k0, v0 := uint(12), uint(34)
	k1, v1 := uint(0), uint(0)
	om0 := orderedmap.New[*uint, *uint]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*uint, *uint]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*uint)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint8(t *testing.T) {
	om0 := orderedmap.New[uint8, uint8]()
	om0.Store(uint8(12), uint8(34))
	om0.Store(uint8(0), uint8(0))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint8Pointer(t *testing.T) {
	k0, v0 := uint8(12), uint8(34)
	k1, v1 := uint8(0), uint8(0)
	om0 := orderedmap.New[*uint8, *uint8]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*uint8, *uint8]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint8(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint8(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*uint8)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint16(t *testing.T) {
	om0 := orderedmap.New[uint16, uint16]()
	om0.Store(uint16(12), uint16(34))
	om0.Store(uint16(0), uint16(0))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint16Pointer(t *testing.T) {
	k0, v0 := uint16(12), uint16(34)
	k1, v1 := uint16(0), uint16(0)
	om0 := orderedmap.New[*uint16, *uint16]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*uint16, *uint16]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint16(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint16(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*uint16)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint32(t *testing.T) {
	om0 := orderedmap.New[uint32, uint32]()
	om0.Store(uint32(12), uint32(34))
	om0.Store(uint32(0), uint32(0))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint32Pointer(t *testing.T) {
	k0, v0 := uint32(12), uint32(34)
	k1, v1 := uint32(0), uint32(0)
	om0 := orderedmap.New[*uint32, *uint32]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*uint32, *uint32]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint32(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint32(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*uint32)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint64(t *testing.T) {
	om0 := orderedmap.New[uint64, uint64]()
	om0.Store(uint64(12), uint64(34))
	om0.Store(uint64(0), uint64(0))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0}`))

	om1 := orderedmap.New[int, int]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 2)

	v, b := om1.Load(12)
	assert.True(t, b)
	assert.Equal(t, v, 34)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_uint64Pointer(t *testing.T) {
	k0, v0 := uint64(12), uint64(34)
	k1, v1 := uint64(0), uint64(0)
	om0 := orderedmap.New[*uint64, *uint64]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*uint64, *uint64]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint64(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, uint64(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*uint64)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_float32(t *testing.T) {
	om0 := orderedmap.New[float32, float32]()
	om0.Store(float32(1.2), float32(3.4))
	om0.Store(float32(0.0), float32(0.0))
	om0.Store(float32(-5.6), float32(-7.8))

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"1.2":3.4,"0":0,"-5.6":-7.8}`))

	om1 := orderedmap.New[float32, float32]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(1.2)
	assert.True(t, b)
	assert.Equal(t, v, float32(3.4))

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, float32(0))

	v, b = om1.Load(-5.6)
	assert.True(t, b)
	assert.Equal(t, v, float32(-7.8))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_float32Pointer(t *testing.T) {
	k0, v0 := float32(12), float32(34)
	k1, v1 := float32(0), float32(0)
	om0 := orderedmap.New[*float32, *float32]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*float32, *float32]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, float32(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, float32(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*float32)(nil))
}

func TestUnmarshalJSON_variousKeyAndValueTypes_float64(t *testing.T) {
	om0 := orderedmap.New[float64, float64]()
	om0.Store(1.2, 3.4)
	om0.Store(0.0, 0.0)
	om0.Store(-5.6, -7.8)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"1.2":3.4,"0":0,"-5.6":-7.8}`))

	om1 := orderedmap.New[float64, float64]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	v, b := om1.Load(1.2)
	assert.True(t, b)
	assert.Equal(t, v, 3.4)

	v, b = om1.Load(0)
	assert.True(t, b)
	assert.Equal(t, v, 0.0)

	v, b = om1.Load(-5.6)
	assert.True(t, b)
	assert.Equal(t, v, -7.8)
}

func TestUnmarshalJSON_variousKeyAndValueTypes_float64Pointer(t *testing.T) {
	k0, v0 := float64(12), float64(34)
	k1, v1 := float64(0), float64(0)
	om0 := orderedmap.New[*float64, *float64]()
	om0.Store(&k0, &v0)
	om0.Store(&k1, &v1)
	om0.Store(nil, nil)

	bs, err := om0.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, bs, []byte(`{"12":34,"0":0,"null":null}`))

	om1 := orderedmap.New[*float64, *float64]()
	err = om1.UnmarshalJSON(bs)
	assert.Nil(t, err)
	assert.Equal(t, om1.Len(), 3)

	ent := om1.Front()
	v, b := om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, float64(34))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, *v, float64(0))

	ent = ent.Next()
	v, b = om1.Load(ent.Key())
	assert.True(t, b)
	assert.Equal(t, v, (*float64)(nil))
}

func TestUnmarshalJSON_badFirstChar(t *testing.T) {
	om := orderedmap.New[string, string]()

	bs := []byte("@")
	err := om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "invalid character '@' looking for beginning of value")
	assert.Equal(t, om.Len(), 0)

	bs = []byte("{@}")
	err = om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "invalid character '@'")
	assert.Equal(t, om.Len(), 0)
}

func TestUnmarshalJSON_noCloseBracket(t *testing.T) {
	om := orderedmap.New[string, string]()

	bs := []byte(`{"A":"a"`)
	err := om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "The input JSON does not end with '}' (offset:8)")
	assert.Equal(t, om.Len(), 1)
}

func TestUnmarshalJSON_keyTypeMismatched_int(t *testing.T) {
	om := orderedmap.New[int, string]()

	bs := []byte(`{"1":"a","B":"b"}`)
	err := om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "invalid character 'B' looking for beginning of value")
	assert.Equal(t, om.Len(), 1)
}

func TestUnmarshalJSON_keyTypeMismatched_intPointer(t *testing.T) {
	om := orderedmap.New[*int, string]()

	bs := []byte(`{"1":"a","B":"b"}`)
	err := om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "invalid character 'B' looking for beginning of value")
	assert.Equal(t, om.Len(), 1)
}

// COMPILE ERROR: any to satisfy comparable requires go1.20 or later (-lang was set to go1.18; check go.mod)
// func TestUnmarshalJSON_keyTypeIsAny_unsupportedError(t *testing.T) {
// 	om := orderedmap.New[any, string]()
//
// 	bs := []byte(`{"A":"a","B":"b"}`)
// 	err := om.UnmarshalJSON(bs)
// 	assert.Equal(t, err.Error(), "json: unsupported key type: any")
// 	assert.Equal(t, om.Len(), 0)
// }

type B struct {
	S, t string
	U    string `json:"u"`
}

func TestUnmarshalJSON_valueTypeMisMatched(t *testing.T) {
	om := orderedmap.New[string, string]()

	bs := []byte(`{"A":{"B":"C":"abc1"}}}`)
	err := om.UnmarshalJSON(bs)
	assert.Equal(t, err.Error(), "Invalid character '{' (offset:6)")
}
