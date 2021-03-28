//
//  post.service.go
//  services
//
//  Created by d-exclaimation on 1:03 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"context"
	"github.com/d-exclaimation/exclaimation-api/ent"
	"github.com/d-exclaimation/exclaimation-api/ent/post"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
	"github.com/d-exclaimation/exclaimation-api/server/errors"
	"net/http"
	"sort"
)

type PostService struct {
	client *ent.Client
}

func PostServiceProvider(client *ent.Client) *PostService {
	return &PostService{
		client: client,
	}
}

func (t *PostService) QueryAll(ctx context.Context) (ent.Posts, error) {
	res, err := t.client.
		Post.Query().
		All(ctx)
	if err != nil {
		return make(ent.Posts, 0), errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	sort.Slice(res, func(i, j int) bool { return res[i].ID < res[j].ID })
	return res, nil
}

func (t *PostService) QueryOne(ctx context.Context, id int) (*ent.Post, error) {
	res, err := t.client.Post.
		Query().
		Where(post.ID(id)).
		First(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (t *PostService) CreateNew(ctx context.Context, input model.PostDto) (*ent.Post, error) {
	res, err := t.client.Post.
		Create().
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (t *PostService) UpdateOne(ctx context.Context, id int, input model.PostDto) (*ent.Post, error) {
	res, err := t.client.Post.
		UpdateOneID(id).
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (t *PostService) ChangeRave(ctx context.Context, id int, value int) (*ent.Post, error) {
	prev, err := t.QueryOne(ctx, id)
	if err != nil {
		return nil, err
	}
	res, err := t.client.Post.
		UpdateOneID(id).
		SetCrabrave(prev.Crabrave + value).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *PostService) DeleteOne(ctx context.Context, id int) (*ent.Post, error) {
	curr, fail := t.QueryOne(ctx, id)
	if fail != nil {
		return nil, fail
	}

	if err := t.client.Post.
		DeleteOneID(id).
		Exec(ctx); err != nil {
			return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return curr, nil
}