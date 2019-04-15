package pars_test

import (
	"germanWords/pars"
	"testing"
)

func TestParse(t *testing.T) {
	s := `//0 – null`

	ss := pars.Parse(s)
	if len(ss) == 0 {
		t.Fatal("parse result empty - must be full")
	}
	if ss[0][0] != "0" {
		t.Fatal("Erroneous parse - key")
	}
	if ss[0][1] != "null" {
		t.Fatal("Erroneous parse - value")
	}
}
