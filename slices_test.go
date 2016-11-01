package goutils

import (
	"testing"
)

func TestSliceToMapNormal(t *testing.T) {
	ss := []string{"a=b", "c=d"}

	m, err := SliceToMap(ss)

	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	if len(m) != 2 {
		t.Errorf("map's length should be 2, but %d", len(m))
	}

	v, ok := m["a"]

	if !ok {
		t.Errorf("map should contains element `a`")
	}
	if v != "b" {
		t.Errorf("a's value should be `b`, but is %s", v)
	}

	v, ok = m["c"]

	if !ok {
		t.Errorf("map should contains element `c`")
	}
	if v != "d" {
		t.Errorf("c's value should be `d`, but is %s", v)
	}

	v, ok = m["e"]

	if ok {
		t.Errorf("map should NOT contains element `e`")
	}

	if v != "" {
		t.Errorf("e's value should be ``, but is %s", v)
	}

}

func TestSliceToMapWithDuplicatedKey(t *testing.T) {
	ss := []string{"a=b", "c=d", "a=f"}

	m, err := SliceToMap(ss)

	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	if len(m) != 2 {
		t.Errorf("map's length should be 2, but %d", len(m))
	}

	v, ok := m["a"]

	if !ok {
		t.Errorf("map should contains element `a`")
	}
	if v != "f" {
		t.Errorf("a's value should be `f`, but is %s", v)
	}

	v, ok = m["c"]

	if !ok {
		t.Errorf("map should contains element `c`")
	}
	if v != "d" {
		t.Errorf("c's value should be `d`, but is %s", v)
	}

	v, ok = m["e"]

	if ok {
		t.Errorf("map should NOT contains element `e`")
	}

	if v != "" {
		t.Errorf("e's value should be ``, but is %s", v)
	}

}
