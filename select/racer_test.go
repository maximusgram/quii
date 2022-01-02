package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
    t.Run("compare speeds, return the server with fastest response", func(t *testing.T){
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(5 * time.Millisecond)
    
        slowUrl := slowServer.URL
        fastUrl := fastServer.URL
    
        want := fastUrl
        got, _ := Racer(slowUrl, fastUrl)
    
        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    }) 

    t.Run("returns an error if a server doesn't respond within 100ms", func(t *testing.T){
        serverA := makeDelayedServer(121 * time.Millisecond)
        serverB := makeDelayedServer(101 * time.Millisecond)

        // By prefixing a function call with defer, it will now call that function 
        // ** at the end of THE _containing function_ **
        defer serverA.Close()
        defer serverB.Close()

        _, err := ConfigurableRacer(serverA.URL, serverB.URL, 10 * time.Millisecond)

        if err == nil {
            t.Error("expected an error but didn't get one")
        }
    })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}