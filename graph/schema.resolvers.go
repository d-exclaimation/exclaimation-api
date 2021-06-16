package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/d-exclaimation/exclaimation-api/graph/generated"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	e "github.com/d-exclaimation/exclaimation-api/server/errors"
	"github.com/d-exclaimation/exclaimation-api/server/libs"
)

func (r *mutationResolver) LoginAsAdmin(ctx context.Context, options model.PasswordInput) (string, error) {
	if config.GetKey() != options.Pass {
		return "", e.InvalidKeyError()
	}
	log.Printf("[INFO] Admin logged in, at %s\n", options.Time)
	return libs.GiveTheCookie(ctx, options.Pass)
}

func (r *mutationResolver) NewPost(ctx context.Context, input model.PostDto) (*model.Post, error) {
	comp, err := libs.StealTheCookie(ctx, config.GetKey())
	if err != nil {
		return nil, e.InvalidKeyError()
	}
	if *comp != config.GetComputedKey() {
		return nil, e.InvalidKeyError()
	}

	res, err := r.post.CreateNew(ctx, input)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input model.PostDto) (*model.Post, error) {
	comp, err := libs.StealTheCookie(ctx, config.GetKey())
	if err != nil {
		return nil, e.InvalidKeyError()
	}
	if *comp != config.GetComputedKey() {
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

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*model.Post, error) {
	comp, err := libs.StealTheCookie(ctx, config.GetKey())
	if err != nil {
		return nil, e.InvalidKeyError()
	}
	if *comp != config.GetComputedKey() {
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

func (r *queryResolver) Posts(ctx context.Context, limit int, by string) ([]*model.Post, error) {
	res, err := r.post.QueryAll(ctx, limit, libs.ToSortBy(by))
	if err != nil {
		return nil, err
	}
	return res.ToGraphQLs(), nil
}

func (r *queryResolver) LatestPost(ctx context.Context) (*model.Post, error) {
	res, err := r.post.GrabLatest(ctx)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), err
}

func (r *queryResolver) Profile(ctx context.Context) (*model.Profile, error) {
	res, err := r.profile.GetProfile(ctx)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) Repos(ctx context.Context, limit int) ([]*model.Repo, error) {
	res, err := r.repo.GetAllRepos(ctx, limit)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQLs(), nil
}

func (r *queryResolver) LatestRepo(ctx context.Context) (*model.Repo, error) {
	res, err := r.repo.GrabLatest(ctx)
	if err != nil {
		return nil, err
	}
	return res.ToGraphQL(), nil
}

func (r *queryResolver) TopLang(ctx context.Context) (*model.Language, error) {
	top, percent := r.repo.GetTopLang(ctx)
	return &model.Language{
		Lang:       top,
		Percentage: percent,
	}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*string, error) {
	return libs.StealTheCookie(ctx, config.GetKey())
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
