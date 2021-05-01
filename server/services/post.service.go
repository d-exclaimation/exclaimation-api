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
	"github.com/d-exclaimation/exclaimation-api/server/libs"
	"net/http"
)

type PostService struct {
	client *ent.Client
}

// PostServiceProvider Fx Provider
func PostServiceProvider(client *ent.Client) *PostService {
	return &PostService{
		client: client,
	}
}

// QueryAll get all records given the specific limit and sorting method
func (t *PostService) QueryAll(ctx context.Context, limit int, by libs.SortBy) (ent.Posts, error) {
	res, err := t.client.
		Post.Query().
		Limit(limit).
		Order(libs.EntOrderBy(by)).
		All(ctx)
	if err != nil {
		return make(ent.Posts, 0), errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

// QueryOne get the correct given the ID
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

// CreateNew added a new record given the input model
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

// UpdateOne updates the given record of the id using the input model
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

// ChangeRave updates just the crabrave count (likes) to a specific record given a value
func (t *PostService) ChangeRave(ctx context.Context, id int, value int) (*ent.Post, error) {
	curr, err := t.client.Post.
		Query().
		Where(post.ID(id)).
		Select(post.FieldCrabrave).
		Int(ctx)
	if err != nil {
		return nil, err
	}
	res, err := t.client.Post.
		UpdateOneID(id).
		SetCrabrave(curr + value).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteOne removes a record given the id
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

// GrabLatest fetch the most recent record
func (t *PostService) GrabLatest(ctx context.Context) (*ent.Post, error) {
	res, err := t.client.Post.
		Query().
		Order(libs.EntOrderBy(libs.ByLatest)).
		First(ctx)
	if err != nil {
	    return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}