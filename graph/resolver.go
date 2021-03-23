package graph

import (
	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/server/services"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
	post *services.PostService
}

// Resolver Constructor
func NewResolver(srv *services.PostService) *Resolver {
	return &Resolver{
		post: srv,
	}
}

// Fx Provider
func ModuleProvider(srv *services.PostService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(srv),
	}
}