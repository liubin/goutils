package goutils

import (
	"testing"
)

func TestPadString(t *testing.T) {
	m := "PadString"
	actual := PadString("12345", 10, "*")

	expected := "12345*****"
	if actual != expected {
		t.Errorf("PadString error,should %s, but %s", expected, actual)
	}

	actual = PadString("123456", 5, "*")

	expected = "12345"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

	actual = PadString("12345", 5, "*")

	expected = "12345"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

	actual = PadString("12", 5, "△")

	expected = "12△△△"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

}

func TestMaskString(t *testing.T) {
	m := "MaskString"

	actual := MaskString("123456", MaskStringParam{
		MaskChar:          "#",
		DisplayCharsCount: 3,
	})

	expected := "123###"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

	actual = MaskString("123456", MaskStringParam{
		MaskChar:          "#",
		DisplayCharsCount: 10,
	})

	expected = "123456"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

	actual = MaskString("123456789", MaskStringParam{
		MaskChar:          "*",
		DisplayCharsCount: 3,
	})

	expected = "123******"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

	actual = MaskString("123456789", MaskStringParam{
		MaskChar:          "*",
		DisplayCharsCount: 3,
		MaskAtTail:        true,
	})

	expected = "******789"
	if actual != expected {
		t.Errorf("%s error,should %s, but %s", m, expected, actual)
	}

}
