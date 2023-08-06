package events

import (
	"context"
	"net/http"
	"strconv"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventsRepo interface {
	CreateEvent(c context.Context, arg dbCon.CreateEventParams) error
	DeleteEvent(c context.Context, id uuid.UUID) error
	GetEvent(c context.Context, id uuid.UUID) (dbCon.Event, error)
	GetEvents(c context.Context, arg dbCon.GetEventsParams) ([]dbCon.Event, error)
	UpdateEvent(c context.Context, arg dbCon.UpdateEventParams) error
}

type Controller struct {
	Repo EventsRepo
}

func (ctrl *Controller) CreateEvent(c *gin.Context) {
	var event dbCon.CreateEventParams

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Repo.CreateEvent(c, event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func (ctrl *Controller) GetEvent(c *gin.Context) {
	eventID := c.Param("id")

	id, err := uuid.Parse(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	event, err := ctrl.Repo.GetEvent(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Cannot fetch event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (ctrl *Controller) ListEvents(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Invalid page size"})
		return
	}

	offset := (page - 1) * pageSize
	arg := dbCon.GetEventsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	events, err := ctrl.Repo.GetEvents(c, arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":   page,
		"events": events,
	})

}

func (ctrl *Controller) UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")

	id, err := uuid.Parse(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Event Not found"})
		return
	}

	var event dbCon.UpdateEventParams
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	event.ID = id

	err = ctrl.Repo.UpdateEvent(c, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
}

func (ctrl *Controller) DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")

	id, err := uuid.Parse(eventID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = ctrl.Repo.DeleteEvent(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Cannot delete event"})
		return
	}

	c.Status(http.StatusNoContent)
}
