// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Name        string
	Description string
	Date        time.Time
}
