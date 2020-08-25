package db

import (
	"github.com/AlexHusleag/bookstore_oauth-api/src/domain/access_token"
	"github.com/AlexHusleag/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	return nil, rest_errors.NewInternalServerError("Database connection not implemented yet")
}
