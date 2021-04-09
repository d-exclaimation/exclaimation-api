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
	"github.com/d-exclaimation/exclaimation-api/server/libs"
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
	if err != nil {
		stale = true
	}

	// Check for timing
	twelve, _ := time.ParseDuration(config.GetRefreshRate())
	if res != nil && time.Now().After(res.LastUpdated.Add(twelve)) {
		stale = true
	}

	// Re-fetch data, ok
	if stale {
		data, fail := t.fetchProfile()
		if fail != nil {
			return nil, fail
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
		SetLastUpdated(time.Now()).
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
		SetLastUpdated(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) fetchProfile() (*raw.Profile, error) {
	var profile raw.Profile
	if err := libs.GetAndParse(config.GetProfileURL(), &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetAllRepos and Helper methods
func (t *ProfileService) GetAllRepos(ctx context.Context, limit int) (ent.Repos, error) {
	// Get from database
	stale := false
	rep, err := t.getOneRepo(ctx)
	if err != nil {
		stale = true
	}
	
	twelve, _ := time.ParseDuration(config.GetRefreshRate())
	if rep != nil && time.Now().After(rep.LastUpdated.Add(twelve)) {
		stale = true
	}
	res := make(ent.Repos, 0)
	// if stale, re-fetch
	if stale {
		data, fail := t.fetchRepos()
		if fail != nil {
			return nil, fail
		}

		// if state, update or create
		res, err = t.updateReposIsh(ctx, rep != nil, data)
		if len(res) > limit {
			res = res[:limit]
		}
	} else {
		res, err = t.getRepos(ctx, limit)
	}
	
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) getOneRepo(ctx context.Context) (*ent.Repo, error) {
	res, err := t.client.Repo.
		Query().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) getRepos(ctx context.Context, limit int) (ent.Repos, error) {
	res, err := t.client.Repo.
		Query().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) updateReposIsh(ctx context.Context, wipe bool, data raw.Repos) (ent.Repos, error) {
	if wipe {
		_, err := t.client.Repo.
			Delete().
			Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	bulk := make([]*ent.RepoCreate, len(data))
	for i, dat := range data {
		lang := ""
		if dat.Language != nil { lang = *dat.Language }
		bulk[i] = t.client.Repo.Create().
			SetRepoName(dat.FullName).
			SetName(dat.Name).
			SetURL(dat.URL).
			SetLanguage(lang).
			SetDescription(dat.Description).
			SetLastUpdated(time.Now())
	}
	res, err := t.client.Repo.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *ProfileService) fetchRepos() (raw.Repos, error) {
	var repos raw.Repos
	if err := libs.GetAndParse(config.GetProfileURL() + "/repos", &repos); err != nil {
		return nil, err
	}
	return repos, nil
}
