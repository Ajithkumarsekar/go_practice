package main

import (
	"fmt"
	"github.com/go_practice/context/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorate(rootHandler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler exited")

	select {
	case <-time.After(time.Second * 10):
		_, _ = fmt.Fprint(w, "Hello")
	// context will be canceled as soon as the client SYN-ACK in the http connection
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, fmt.Sprintf("context go cancled : %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
