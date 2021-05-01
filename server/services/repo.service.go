//
//  repo.service.go
//  services
//
//  Created by d-exclaimation on 1:31 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package services

import (
	"context"
	"github.com/d-exclaimation/exclaimation-api/config"
	"github.com/d-exclaimation/exclaimation-api/ent"
	"github.com/d-exclaimation/exclaimation-api/ent/repo"
	"github.com/d-exclaimation/exclaimation-api/server/libs"
	"github.com/d-exclaimation/exclaimation-api/server/models/raw"
	"sort"
	"time"
)

type RepoService struct {
	client *ent.Client
}

// RepoServiceProvider Fx Provider
func RepoServiceProvider(client *ent.Client) *RepoService {
	return &RepoService{
		client: client,
	}
}

// Section: Repos Fetch Logic

// GetAllRepos will try to fetch all repos in the database, but will do checks either to fetch into Github's API
func (t *RepoService) GetAllRepos(ctx context.Context, limit int) (ent.Repos, error) {
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
func (t *RepoService) getOneRepo(ctx context.Context) (*ent.Repo, error) {
	res, err := t.client.Repo.
		Query().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *RepoService) getRepos(ctx context.Context, limit int) (ent.Repos, error) {
	res, err := t.client.Repo.
		Query().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *RepoService) updateReposIsh(ctx context.Context, wipe bool, data raw.Repos) (ent.Repos, error) {
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

func (t *RepoService) fetchRepos() (raw.Repos, error) {
	var repos raw.Repos
	if err := libs.GetAndParse(config.GetProfileURL() + "/repos?sort=updated", &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// Section: Top Lang fetch function

// GetTopLang and Helper Methods
func (t *RepoService) GetTopLang(ctx context.Context) (lang string, percentage float64) {
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

// GrabLatest get the most recent repo I have
func (t *RepoService) GrabLatest(ctx context.Context) (*ent.Repo, error) {
	rep, err := t.client.Repo.
		Query().
		Order(ent.Asc(repo.FieldID)).
		First(ctx)
	if err != nil {
	    return nil, err
	}
	return rep, nil
}