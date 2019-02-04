package hello

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/zevdg/fcf"
)

var pp = spew.ConfigState{Indent: "→"}

func dump(prefix string, v interface{}) {
	for _, line := range strings.Split(pp.Sdump(v), "\n") {
		if line != "" {
			fmt.Printf("%s: %s\n", prefix, line)
		}
	}
}

func FcfTest(ctx context.Context, e fcf.Event) error {
	dump("raw", e.Value)
	s := "bar"
	d := &Data{ZeroMe: "foo", NilMe: &s}
	e.Value.Decode(d)
	dump("clean", d)
	return nil
}

type Data struct {
	Name    string
	Age     int
	Office  string
	Temp    float32
	IsAdmin bool
	Time    time.Time
	NilMe   *string
	ZeroMe  string
}
