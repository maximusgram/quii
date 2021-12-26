package dep_inj

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Madhu")

	got := buffer.String()
	want := "Hello, Madhu"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
