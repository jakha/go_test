package log

import (
	"context"
	"log"
	"net/http"
	"math/rand"
)

const requestIDKey = 42

func Println(ctx context.Context, msg string){
	id, ok :=	ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("not found req id ")
		return
	}
	log.Printf("[%d] %s", id, msg)
}


func Decorate ( f http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
