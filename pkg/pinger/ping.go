package pinger

import (
	"log"
	"net/http"
)

func Ping(url string) error {
	log.Printf("PINGING_URL: url=%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}