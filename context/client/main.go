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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)

	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// why not pass context to http.NewRequest?
	// because context was introduced in go 1.7 and adding a new field to http.NewRequest would break backward compatibility
	req = req.WithContext(ctx)

	// client will send SYN-ACK on the TCP connection when the context gets canceled
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error while making the request : %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("server responded with status code : %v", resp.StatusCode)
	}

	io.Copy(os.Stdout, resp.Body)
}
