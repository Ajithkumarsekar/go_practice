package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type RequestIdKey string

const requestIdKey = RequestIdKey("request_id")

func Println(ctx context.Context, msg string) {
	requestId, ok := ctx.Value(requestIdKey).(int64)
	if !ok {
		log.Printf("could not find %v in context", requestIdKey)
		return
	}

	log.Printf("[%d] %s", requestId, msg)
}

// Decorate sets a unique key in context
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestIdKey, rand.Int63())
		f(w, r.WithContext(ctx))
	}
}
