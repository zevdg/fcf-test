package hello

import (
	"context"
	"time"

	"github.com/zevdg/fcf"
)

func Nested(ctx context.Context, e fcf.Event) error {
	dump("raw", e.Value)
	d := &Assessment{}
	e.Value.Decode(d)
	dump("clean", d)
	return nil
}

// Assessment stores assessment in db
type Assessment struct {
	TemplateID string           `fcf:"templateID"`
	Created    time.Time        `fcf:"created"`
	UserID     string           `fcf:"userID"`
	Slots      map[string]*Slot `fcf:"slots"`
}

type Slot struct {
	Value   interface{} `fcf:"value"`
	Created time.Time   `fcf:"created"`
}
