package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/object88/golang-relay-todo-modern/data"
)

func main() {
	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	port := ":8080"
	log.Printf(`GraphQL server starting up on http://localhost%s`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
