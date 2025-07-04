package selected

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowErrorServer := makeDelayedServer(20 * time.Second)

		defer slowErrorServer.Close()

		_output, err := Racer(slowErrorServer.URL, slowErrorServer.URL)

		fmt.Println(_output, err)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
	t.Run("return the url of the fastest server", func(t *testing.T) {
		slowServer := makeDelayedServer(50 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		fmt.Println(slowURL, fastURL)
		want := fastURL
		got, error := Racer(slowURL, fastURL)
		fmt.Println(got, error)
		if error != nil {
			t.Fatalf("did not expect an error but got one %v", error)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
