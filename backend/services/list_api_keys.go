// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

package services

import (
	"context"

	"github.com/theparanoids/ashirt/backend"
	"github.com/theparanoids/ashirt/backend/database"
	"github.com/theparanoids/ashirt/backend/dtos"
	"github.com/theparanoids/ashirt/backend/models"
	"github.com/theparanoids/ashirt/backend/policy"
	"github.com/theparanoids/ashirt/backend/server/middleware"

	sq "github.com/Masterminds/squirrel"
)

func ListAPIKeys(ctx context.Context, db *database.Connection, userSlug string) ([]*dtos.APIKey, error) {
	var userID int64
	var err error

	if userID, err = selfOrSlugToUserID(ctx, db, userSlug); err != nil {
		return nil, backend.DatabaseErr(err)
	}

	if err := policy.Require(middleware.Policy(ctx), policy.CanListAPIKeys{UserID: userID}); err != nil {
		return nil, backend.UnauthorizedReadErr(err)
	}

	var keys []models.APIKey
	err = db.Select(&keys, sq.Select("access_key", "last_auth").
		From("api_keys").
		Where(sq.Eq{"user_id": userID}))

	if err != nil {
		return nil, backend.DatabaseErr(err)
	}

	keysDTO := make([]*dtos.APIKey, len(keys))
	for i, key := range keys {
		keysDTO[i] = &dtos.APIKey{
			AccessKey: key.AccessKey,
			LastAuth:  key.LastAuth,
		}
	}
	return keysDTO, nil
}
