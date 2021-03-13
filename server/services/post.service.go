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
)

// PostService Struct
type PostService struct {
}

// Fx Provider
func PostServiceProvider() *PostService {
	return &PostService{
	}
}

// Methods
func (srv *PostService) CreateNew(input model.NewThought) (*entities.Thought, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) GetAll() (entities.ThoughtsArray, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) GetOne(id int) (*entities.Thought, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) UpdateOne(id int, userId int, input model.NewThought) (*entities.Thought, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}

func (srv *PostService) DeleteOne(id int, userId int) (*entities.Thought, *errors.ServiceError) {
	// TODO: Delete One Item
	panic("Not implemented")
}