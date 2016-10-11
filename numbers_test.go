package goutils

import (
	"testing"
)

func TestFloat64ToInt64(t *testing.T) {

	m := "Float64ToInt64"
	actual := Float64ToInt64(234234.345)

	expected := int64(234234)
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}
}
