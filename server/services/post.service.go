//
//  post.service.go
//  services
//
//  Created by d-exclaimation on 8:17 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"github.com/d-exclaimation/exclaimation-gql/db"
	"github.com/d-exclaimation/exclaimation-gql/db/entities"
	"github.com/d-exclaimation/exclaimation-gql/graph/model"
	"github.com/d-exclaimation/exclaimation-gql/server/errors"
	"gopkg.in/gorp.v1"
	"net/http"
)

// PostService Struct
type PostService struct {
	rp *gorp.DbMap
}

// Fx Provider
func PostServiceProvider(rp *gorp.DbMap) *PostService {
	return &PostService{
		rp: rp,
	}
}

// Methods
func (srv *PostService) CreateNew(input model.PostDto) (*entities.Post, *errors.ServiceError) {
	post := entities.NewPost(input.Title, input.Body)
	if err := srv.rp.Insert(post); err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Cannot connect to the database")
	}
	return post, nil
}

func (srv *PostService) GetAll() (entities.Posts, *errors.ServiceError) {
	var posts entities.Posts
	_, err := srv.rp.Select(&posts, "SELECT * FROM " + string(db.PostsTable) + " order by created_at")
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Cannot fetch data from the database")
	}
	return posts, nil
}

func (srv *PostService) GetOne(id int) (*entities.Post, *errors.ServiceError) {
	post, err := srv.rp.Get(entities.Post{}, id)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return post.(*entities.Post), nil
}

func (srv *PostService) UpdateOne(id int, input model.PostDto) (*entities.Post, *errors.ServiceError) {
	post, fail := srv.GetOne(id)
	if fail != nil {
		return nil, fail
	}

	post.Title = input.Title
	post.Body = input.Body
	_, err := srv.rp.Update(post)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Cannot update data to the database")
	}
	return post, nil
}

func (srv *PostService) DeleteOne(id int) (*entities.Post, *errors.ServiceError) {
	post, fail := srv.GetOne(id)
	if fail != nil {
		return nil, fail
	}

	copied := &entities.Post{
		Id:        post.Id,
		CreatedAt: post.CreatedAt,
		Title:     post.Title,
		Body:      post.Body,
		Agrees:    post.Agrees,
		Disagree:  post.Disagree,
	}

	_, err := srv.rp.Delete(post)
	if err != nil {
		return nil, errors.NewServiceError(http.StatusInternalServerError, "Cannot delete data from the database")
	}
	return copied, nil
}