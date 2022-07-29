package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type key int

const (
	userIDctx key = 0
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-id")

	ctx := context.WithValue(r.Context(), userIDctx, id)

	result := processLongTask(ctx)

	w.Write([]byte(result))

	r.ParseForm()
	ids := r.FormValue("id")
	if ids == "id1" {
		fmt.Println("Проверочка")
		io.WriteString(w, ids)
	}
}

func processLongTask(ctx context.Context) string {
	id := ctx.Value(userIDctx)

	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprint("нет ид", id)
	case <-ctx.Done():
		log.Println("завершаем")
		return "жжжопа"
	}
}
