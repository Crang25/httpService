package router

import (
	"net/http"

	"github.com/Crang25/httpService/internal/storages"
)

type Router struct {
	rootHandler rootHandler
}

func New(store storages.Store) *Router {
	r := &Router{
		rootHandler: newRootHandler(store),
	}
	return r
}

func (r *Router) RootHandler() http.Handler {
	return &r.rootHandler
}
