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
	"github.com/d-exclaimation/exclaimation-api/ent"
	"github.com/d-exclaimation/exclaimation-api/server/models/raw"
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

func (t *ProfileService) GetProfile(ctx context.Context) (*ent.Profile, error) {
	// TODO: 1. Grab from database
	// TODO: 2. Check for timing
	// TODO: 2.1. Re-fetch data
	// TODO: 2.2. Insert or Create New Record
	// TODO: 3. Send Data
	return &ent.Profile{}, nil
}

func (t *ProfileService) GetAllRepos(limit int) (ent.Repos, error) {
	return make(ent.Repos, 0), nil
}


func (t *ProfileService) FetchProfile() (*raw.Profile, error) {
	//res, err := http.Get(config.GetProfileURL())
	panic("not implemented")
}