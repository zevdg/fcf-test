package hello

import (
	"context"
	"time"

	"github.com/zevdg/fcf"
)

func FcfTest(ctx context.Context, e fcf.Event) error {
	dump("raw", e.Value)
	s := "bar"
	d := &Data{ZeroMe: "foo", NilMe: &s}
	// d := &Data{}
	e.Value.Decode(d)
	dump("clean", d)
	return nil
}

type Data struct {
	Name      string
	Age       int
	Office    string
	Temp      float32
	IsAdmin   bool
	Time      time.Time
	NilMe     *string
	ZeroMe    string
	ReMap     int `fcf:"other"`
	Dynamic   interface{}
	MapStruct struct {
		Baz    int
		Foo    string
		SubMap struct {
			X0 string `fcf:"x0"`
			X1 string `fcf:"x1"`
		}
		Inner1 map[string]string      `fcf:"SubMap"`
		Inner2 map[string]interface{} `fcf:"SubMap"`
		Inner3 interface{}            `fcf:"SubMap"`
	} `fcf:"Map"`
	MapDynamic   map[string]interface{} `fcf:"Map"`
	MapInterface interface{}            `fcf:"Map"`
	TypedMap     map[string]string
	Location     struct {
		Lat  float32 `fcf:"latitude"`
		Long float32 `fcf:"longitude"`
	}
	TypedArray     []string
	DynamicArray   []interface{}
	ArrayInterface interface{}  `fcf:"DynamicArray"`
	Loc2           fcf.GeoPoint `fcf:"Location"`
	Bytes          []byte
	RMap           rMap
}

type rMap struct {
	Name string
	Arr  []rMap
}
