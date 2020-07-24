package server_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compara a velocidade de servidores, retornando o endereço do mais rápido", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf(" obteve '%s', queria '%s'", got, want)
		}
	})

	t.Run("retorna um erro se o servidor não responder dentro de 10s", func(t *testing.T) {
		server := makeDelayedServer(20 * time.Millisecond)

		defer server.Close()

		want := ErrorTimeout
		_, err := ConfigurableRacer(server.URL, server.URL, (25 * time.Millisecond))

		confirmError(t, err, want)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func confirmError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("é esperado um erro '%s'", got)
	}
}
