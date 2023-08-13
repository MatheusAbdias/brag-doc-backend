package tests

import (
	"context"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/events"
	"github.com/google/uuid"
)

type FakeEventRepo struct {
	events []dbCon.Event
}

func (repo *FakeEventRepo) CreateEvent(c context.Context, arg dbCon.CreateEventParams) error {
	return nil
}

func (repo *FakeEventRepo) GetEvents(c context.Context, arg dbCon.GetEventsParams) ([]dbCon.Event, error) {
	return repo.events, nil
}

func (repo *FakeEventRepo) GetEvent(c context.Context, id uuid.UUID) (dbCon.Event, error) {
	return dbCon.Event{}, nil
}

func (repo *FakeEventRepo) UpdateEvent(c context.Context, arg dbCon.UpdateEventParams) error {
	return nil
}

func (repo *FakeEventRepo) DeleteEvent(c context.Context, id uuid.UUID) error {
	return nil
}
