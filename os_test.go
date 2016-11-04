package goutils

import (
	"os"
	"testing"
)

func TestGetEnvInt64(t *testing.T) {
	os.Clearenv()

	key := "key"

	i := GetEnvInt64(key, -1)
	var expected int64 = -1

	if i != expected {
		t.Errorf("i should be %d, but %d", expected, i)
	}

	os.Setenv(key, "123")
	expected = 123
	i = GetEnvInt64(key, -1)
	if i != expected {
		t.Errorf("i should be %d, but %d", expected, i)
	}

	os.Setenv(key, "a123")
	i = GetEnvInt64(key, -1)
	expected = -1
	if i != expected {
		t.Errorf("i should be %d, but %d", expected, i)
	}

}

func TestGetEnvString(t *testing.T) {

	os.Clearenv()
	key := "key"

	s := GetEnvString(key, "abc")
	var expected string = "abc"

	if s != expected {
		t.Errorf("s should be %s, but %s", expected, s)
	}

	os.Setenv(key, "123")
	s = GetEnvString(key, "abc")
	expected = "123"

	if s != expected {
		t.Errorf("s should be %s, but %s", expected, s)
	}

}
