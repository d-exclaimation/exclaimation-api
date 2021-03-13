//
//  post.entity.go
//  entities
//
//  Created by d-exclaimation on 12:11 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package entities

import (
	"fmt"
	"github.com/d-exclaimation/exclaimation-gql/graph/model"
	_ "gopkg.in/gorp.v1"
	"time"
)

type Post struct {
	Id  	  int64  `db:"post_id"`
	CreatedAt int64  `db:"created_at"`
	Title 	  string `db:"title, size:60"`
	Body	  string `db:"body,size:1024"`
	Agrees 	  int	 `db:"agree"`
	Disagree  int	 `db:"disagree"`
}

func NewPost(title string, body string) *Post {
	return &Post{
		CreatedAt: time.Now().UnixNano(),
		Title:     title,
		Body:      body,
		Agrees:    0,
		Disagree:  0,
	}
}

func (post *Post) ToGraphQL() *model.Post {
	return &model.Post{
		ID:       fmt.Sprintf("%d", post.Id),
		Title:    post.Title,
		Body:     post.Body,
		Agree:    post.Agrees,
		Disagree: post.Disagree,
	}
}

type Posts []*Post

func (posts Posts) ToGraphQLs() []*model.Post {
	res := make([]*model.Post, len(posts))
	for i, post := range posts {
		res[i] = post.ToGraphQL()
	}
	return res
}
