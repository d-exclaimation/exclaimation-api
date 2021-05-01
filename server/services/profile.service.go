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
	"sort"
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

/// Section: Profile Fetch Logic

// GetProfile will try to look up profile in the database, if not available or outdated, will do re-fetch onto Github's API
func (t *ProfileService) GetProfile(ctx context.Context) (*ent.Profile, error) {
	// Grab from database if exist, and check for stale data
	stale := false
	res, err := t.queryProfile(ctx)
	if err != nil {
		stale = true
	}

	// Check for timing (A bit explicit)
	twelve, _ := time.ParseDuration(config.GetRefreshRate())
	if res != nil && time.Now().After(res.LastUpdated.Add(twelve)) {
		stale = true
	}

	// Re-fetch data, ok
	if stale {
		data, fail := t.fetchProfile()

		// Failure to fetch, if stale exist, return stale data. Otherwise, tell failure
		if fail != nil {
			if res != nil {
				return res, nil
			}
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

// Try query profile from the database
func (t *ProfileService) queryProfile(ctx context.Context) (*ent.Profile, error) {
	res, err := t.client.Profile.
		Query().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create a new profile record in the database
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

// Updated existing record of the profile given a outdated record id and the raw fetch result
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

// Try to fetch to Github's API, the profile information
func (t *ProfileService) fetchProfile() (*raw.Profile, error) {
	var profile raw.Profile
	if err := libs.GetAndParse(config.GetProfileURL(), &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// Section: Repos Fetch Logic

// GetAllRepos will try to fetch all repos in the database, but will do checks either to fetch into Github's API
func (t *ProfileService) GetAllRepos(ctx context.Context, limit int) (ent.Repos, error) {
	// Get from database, if exist and if not stale
	stale := false
	res := make(ent.Repos, 0)
	rep, err := t.getOneRepo(ctx)
	if err != nil {
		stale = true
	}
	
	twelve, _ := time.ParseDuration(config.GetRefreshRate())
	if rep != nil && time.Now().After(rep.LastUpdated.Add(twelve)) {
		stale = true
	}

	// if stale, re-fetch
	data, fail := make(raw.Repos, 0), err
	if stale {
		data, fail = t.fetchRepos()
	}

	// if fetch fails and there is no previous, throw error
	if fail != nil && err != nil {
		return nil, fail
	}

	// if stale, does not fail to fetch, do updates
	if stale && fail == nil {
		res, err = t.updateReposIsh(ctx, rep != nil, data)
		if len(res) > limit { res = res[:limit] }

	// else (not stale or stale but fetch fail) do query only
	} else {
		res, err = t.getRepos(ctx, limit)
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

// Try to query one
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
		_, err := t.client.DB().Exec("TRUNCATE repos RESTART IDENTITY;")
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
			SetURL(dat.HTMLURL).
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

// Section: Top Lang fetch function

// GetTopLang and Helper Methods
func (t *ProfileService) GetTopLang(ctx context.Context) (lang string, percentage float64) {
	repos, err := t.GetAllRepos(ctx, 100)
	if err != nil {
		return "English", 69.420
	}

	langs := make(map[string]int)
	for _, rep := range repos {
		prev, exist := langs[rep.Language]
		if !exist { prev = 0 }
		langs[rep.Language] = prev + 1
	}

	ranking := make([]string, len(langs))
	k := 0
	for lang := range langs {
		ranking[k] = lang
		k++
	}

	sort.Slice(ranking, func(i, j int) bool {
		lhs, rhs := langs[ranking[i]], langs[ranking[j]]
		return lhs > rhs
	})
	if len(ranking) < 1 {
		return "English", 69.420
	}
	return ranking[0], (float64(langs[ranking[0]]) / float64(len(repos))) * 100
}