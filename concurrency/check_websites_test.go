package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
    if url == "http://fail.com" {
        return false
    }
    return true
}

func TestCheckWebsites(t *testing.T) {
    websites := []string {
        "http://google.com",
        "http://blog.com",
        "http://fail.com",
    }

    got := CheckWebsites(mockWebsiteChecker, websites)
    want := map[string]bool{
        "http://google.com": true,
        "http://blog.com": true,
        "http://fail.com": false,
    }

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}