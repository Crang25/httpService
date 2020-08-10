package main

import (
	"net/http"

	"github.com/Crang25/httpService/internal/storages/memstore"

	"github.com/Crang25/httpService/internal/router"
)

func main() {
	r := router.New(memstore.New())
	http.ListenAndServe(":8080", r.RootHandler())
}
