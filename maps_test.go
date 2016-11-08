package goutils

import (
	"testing"
)

func TestMergeMap(t *testing.T) {
	m1 := map[string]string{"k1": "error", "k2": "v2"}
	m2 := map[string]string{"k1": "v1", "k3": "v3"}

	m3 := MergeStringMap(m1, m2)

	if len(m3) != 3 {
		t.Errorf("m3 should contains 3 elements, but %d", len(m3))
	}

	v1, _ := m3["k1"]
	if v1 != "v1" {
		t.Errorf("v1 should be v1, but %s", v1)
	}

	v2, _ := m3["k2"]
	if v2 != "v2" {
		t.Errorf("v2 should be v2, but %s", v2)
	}

	v3, _ := m3["k3"]
	if v3 != "v3" {
		t.Errorf("v3 should be v3, but %s", v3)
	}

}
