package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/d-exclaimation/exclaimation-gql/graph/generated"
	"github.com/d-exclaimation/exclaimation-gql/graph/model"
)

func (r *mutationResolver) CreatePost(_ context.Context, input model.PostDto, access string) (*model.Post, error) {
	if err := ResolverAccessHandler(access); err != nil {
		return nil, err
	}

	// Get from the service and handle error given
	post, err := r.srv.CreateNew(input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return post.ToGraphQL(), nil
}

func (r *mutationResolver) UpdatePost(_ context.Context, id int, input model.PostDto, access string) (*model.Post, error) {
	if err := ResolverAccessHandler(access); err != nil {
		return nil, err
	}

	// Update to the service and handle error given
	post, err := r.srv.UpdateOne(id, input)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return post.ToGraphQL(), nil
}

func (r *mutationResolver) DeletePost(_ context.Context, id int, access string) (*model.Post, error) {
	if err := ResolverAccessHandler(access); err != nil {
		return nil, err
	}

	// Delete to the service and handle error
	post, err := r.srv.DeleteOne(id)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return post.ToGraphQL(), nil
}

func (r *queryResolver) Posts(_ context.Context) ([]*model.Post, error) {
	posts, err := r.srv.GetAll()
	if err != nil {
		return nil, err.ToGQLError()
	}
	return posts.ToGraphQLs(), nil
}

func (r *queryResolver) Post(_ context.Context, id int) (*model.Post, error) {
	post, err := r.srv.GetOne(id)
	if err != nil {
		return nil, err.ToGQLError()
	}
	return post.ToGraphQL(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
