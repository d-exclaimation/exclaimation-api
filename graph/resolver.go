package graph

import (
	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/server/services"
)

//go:generate go run github.com/99designs/gqlgen

// Resolver Struct
type Resolver struct {
	post *services.PostService
	profile *services.ProfileService
	repo *services.RepoService
}

// NewResolver Constructor
func NewResolver(
		post *services.PostService,
		profile *services.ProfileService,
		repo *services.RepoService,
	) *Resolver {
	return &Resolver{
		post: post,
		profile: profile,
		repo: repo,
	}
}

// ModuleProvider Fx Provider
func ModuleProvider(post *services.PostService, profile *services.ProfileService, repo *services.RepoService) generated.Config {
	return generated.Config {
		Resolvers: NewResolver(post, profile, repo),
	}
}