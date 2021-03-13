//
//  post.service.go
//  services
//
//  Created by d-exclaimation on 8:17 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"github.com/d-exclaimation/exclaimation-gql/db/entities"
	"github.com/d-exclaimation/exclaimation-gql/graph/model"
	"github.com/d-exclaimation/exclaimation-gql/server/errors"
	"gopkg.in/gorp.v1"
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
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) GetAll() (entities.Posts, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) GetOne(id int) (*entities.Post, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) UpdateOne(id int, input model.PostDto) (*entities.Post, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) DeleteOne(id int) (*entities.Post, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}