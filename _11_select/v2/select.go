package server_racer

import (
	"errors"
	"net/http"
	"time"
)

const tenSecondTimeout = time.Duration(10 * time.Second)

var ErrorTimeout = errors.New("Tempo esgotado")

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrorTimeout
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}
