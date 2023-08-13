package events

import (
	"net/http"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/events"
	"github.com/gin-gonic/gin"
)

func EventValidators() []func(c *gin.Context, event dbCon.CreateEventParams) {
	return []func(c *gin.Context, event dbCon.CreateEventParams){
		ValidateName,
		ValidateDescription,
	}
}

func ValidateName(c *gin.Context, event dbCon.CreateEventParams) {
	if event.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event name is cannot be empty"})
		return
	}
}

func ValidateDescription(c *gin.Context, event dbCon.CreateEventParams) {
	if event.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event description is cannot be empty"})
		return
	}
}
