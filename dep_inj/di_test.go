package dep_inj

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Maximus")

	got := buffer.String()
	want := "Hello, Maximus"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
