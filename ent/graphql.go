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

func (po *Post) ToGraphQL() *model.Post {
	return &model.Post{
		ID:       fmt.Sprintf("%d", po.ID),
		Title:    po.Title,
		Body:     po.Body,
		Crabrave: po.Crabrave,
	}
}

// ToGraphQLs Convert all to GraphQL Schema
func (po Posts) ToGraphQLs() []*model.Post {
	res := make([]*model.Post, len(po))
	for i, Post := range po {
		res[i] = Post.ToGraphQL()
	}
	return res
}

func (pr *Profile) ToGraphQL() *model.Profile {
	return &model.Profile{
		AvatarURL:       pr.AvatarURL,
		GithubURL:       pr.GithubURL,
		Name:            pr.Name,
		Location:        pr.Location,
		Bio:             pr.Bio,
		TwitterUsername: pr.TwitterUsername,
		ReposCount:      pr.PublicRepo,
		Followers:       pr.Followers,
		Following:       pr.Following,
	}
}

func (r *Repo) ToGraphQL() *model.Repo {
	lang := &r.Language
	if r.Language == "" {
		lang = nil
	}
	return &model.Repo{
		ID:          fmt.Sprintf("%d", r.ID),
		Name:        r.Name,
		RepoName:    r.RepoName,
		URL:         r.URL,
		Description: r.Description,
		Language:    lang,
	}
}

// ToGraphQLs Convert all to GraphQL schema
func (r Repos) ToGraphQLs() []*model.Repo {
	res := make([]*model.Repo, len(r))
	for i, repo := range r {
		res[i] = repo.ToGraphQL()
	}
	return res
}