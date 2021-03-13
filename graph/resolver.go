package graph

import (
	"github.com/d-exclaimation/exclaimation-gql/graph/generated"
	"github.com/d-exclaimation/exclaimation-gql/server/services"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
	srv *services.PostService
}

// Resolver Constructor
func NewResolver(srv *services.PostService) *Resolver {
	return &Resolver{
		srv: srv,
	}
}

// Fx Provider
func ModuleProvider(srv *services.PostService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(srv),
	}
}