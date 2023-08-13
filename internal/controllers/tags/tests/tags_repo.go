package tests

import (
	"context"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"
	"github.com/google/uuid"
)

type FakeTagRepo struct {
	tags []dbCon.Tag
}

func (repo *FakeTagRepo) CreateTag(c context.Context, name string) error {
	return nil
}

func (repo *FakeTagRepo) GetTags(c context.Context, arg dbCon.GetTagsParams) ([]dbCon.Tag, error) {
	return repo.tags, nil
}

func (repo *FakeTagRepo) GetTag(c context.Context, id uuid.UUID) (dbCon.Tag, error) {
	return dbCon.Tag{}, nil
}

func (repo *FakeTagRepo) UpdateTag(c context.Context, arg dbCon.UpdateTagParams) error {
	return nil
}

func (repo *FakeTagRepo) DeleteTag(c context.Context, id uuid.UUID) error {
	return nil
}
