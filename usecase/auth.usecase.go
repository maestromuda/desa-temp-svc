package usecase

import (
	"context"
	"desa-temp-svc/repository"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(ctx context.Context, username, password string) (bson.M, error) {
	// logic

	res, err := repository.Login(ctx, username, password)
	if err != nil {
		return bson.M{}, err
	}

	return res, nil
}
