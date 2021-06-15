package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8011/", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), "time_out", time.Duration(1000))
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
