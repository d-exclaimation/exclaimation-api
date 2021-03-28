package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	e "github.com/d-exclaimation/exclaimation-api/server/errors"
)

func (r *mutationResolver) NewPost(ctx context.Context, input model.PostDto, key string) (*model.Post, error) {
	if key != config.GetKey() {
		return nil, e.InvalidKeyError()
	}

	res, err := r.post.CreateNew(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input model.PostDto, key string) (*model.Post, error) {
	if key != config.GetKey() {
		return nil, e.InvalidKeyError()
	}
	res, err := r.post.UpdateOne(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) IncrementCrabRave(ctx context.Context, id int) (*model.Post, error) {
	res, err := r.post.ChangeRave(ctx, id, 1)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), err
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int, key string) (*model.Post, error) {
	if key != config.GetKey() {
		return nil, e.InvalidKeyError()
	}

	res, err := r.post.DeleteOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	res, err := r.post.QueryOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	res, err := r.post.QueryAll(ctx)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQLs(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
