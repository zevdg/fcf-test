package hello

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var pp = spew.ConfigState{Indent: "→"}

func dump(prefix string, v interface{}) {
	for _, line := range strings.Split(pp.Sdump(v), "\n") {
		if line != "" {
			fmt.Printf("%s: %s\n", prefix, line)
		}
	}
}
