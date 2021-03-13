package graph

import (
	"github.com/d-exclaimation/exclaimation-gql/config"
	"github.com/d-exclaimation/exclaimation-gql/graph/generated"
	"github.com/d-exclaimation/exclaimation-gql/server/services"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
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

// Access Token Handling
func ResolverAccessHandler(access string) error {
	if access != config.GetAccessToken() {
		return gqlerror.Errorf("(%d) %s", http.StatusForbidden, "Invalid Access")
	}
	return nil
}