//
//  graphql.go
//  ent
//
//  Created by d-exclaimation on 12:56 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package ent

import (
	"fmt"
	"github.com/d-exclaimation/exclaimation-api/graph/model"
)

func (t *Post) ToGraphQL() *model.Post {
	return &model.Post{
		ID:       fmt.Sprintf("%d", t.ID),
		Title:    t.Title,
		Body:     t.Body,
		Crabrave: t.Crabrave,
	}
}


// Convert all to GraphQL Schema
func (po Posts) ToGraphQLs() []*model.Post {
	res := make([]*model.Post, len(po))
	for i, Post := range po {
		res[i] = Post.ToGraphQL()
	}
	return res
}
