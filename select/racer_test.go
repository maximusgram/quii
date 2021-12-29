package racer

import "testing"

func TestRacer(t *testing.T) {
    slowUrl := "http://www.facebook.com"
    fastUrl := "http://www.quii.co.uk"

    want := fastUrl
    got := Racer(slowUrl, fastUrl)

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}