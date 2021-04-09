//
//  profile.service.go
//  services
//
//  Created by d-exclaimation on 3:36 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"context"
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/d-exclaimation/exclaimation-api/ent"
	"github.com/d-exclaimation/exclaimation-api/graph/libs"
	"github.com/d-exclaimation/exclaimation-api/server/models/raw"
	"time"
)

type ProfileService struct {
	client *ent.Client
}

// ProfileServiceProvider Fx Provider
func ProfileServiceProvider(client *ent.Client) *ProfileService {
	return &ProfileService{
		client: client,
	}
}

// GetProfile and Helper functions
func (t *ProfileService) GetProfile(ctx context.Context) (*ent.Profile, error) {
	// Grab from database, ok
	stale := false
	res, err := t.queryProfile(ctx)
	if err != nil || res == nil {
		stale = true
	}

	// Check for timing
	twelve, _ := time.ParseDuration("12h")
	if res != nil && time.Now().After(res.LastUpdated.Add(twelve)) {
		stale = true
	}

	// Re-fetch data, ok
	if stale {
		data, err := t.fetchProfile()
		if err != nil {
			return nil, err
		}

		// Insert or Create New Record, ok
		if res != nil {
			res, err = t.updateProfile(ctx, res.ID, data)
		} else {
			res, err = t.createProfile(ctx, data)
		}
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (t *ProfileService) queryProfile(ctx context.Context) (*ent.Profile, error) {
	res, err := t.client.Profile.
		Query().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) createProfile(ctx context.Context, raw *raw.Profile) (*ent.Profile, error) {
	res, err := t.client.Profile.Create().
		SetAvatarURL(raw.AvatarURL).
		SetGithubURL(raw.HtmlURL).
		SetName(raw.Name).
		SetLocation(raw.Location).
		SetBio(raw.Bio).
		SetTwitterUsername(raw.TwitterUsername).
		SetPublicRepo(raw.PublicRepos).
		SetFollowers(raw.Followers).
		SetFollowing(raw.Following).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) updateProfile(ctx context.Context, id int, raw *raw.Profile) (*ent.Profile, error) {
	res, err := t.client.Profile.UpdateOneID(id).
		SetAvatarURL(raw.AvatarURL).
		SetGithubURL(raw.HtmlURL).
		SetName(raw.Name).
		SetLocation(raw.Location).
		SetBio(raw.Bio).
		SetTwitterUsername(raw.TwitterUsername).
		SetPublicRepo(raw.PublicRepos).
		SetFollowers(raw.Followers).
		SetFollowing(raw.Following).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) fetchProfile() (*raw.Profile, error) {
	var profile *raw.Profile
	if err := libs.GetAndParse(config.GetProfileURL(), profile); err != nil {
		return nil, err
	}
	return profile, nil
}

// GetAllRepos and Helper methods
func (t *ProfileService) GetAllRepos(limit int) (ent.Repos, error) {
	return make(ent.Repos, 0), nil
}
