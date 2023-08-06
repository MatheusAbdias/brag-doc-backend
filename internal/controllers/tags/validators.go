package controllers

import (
	"net/http"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"

	"github.com/gin-gonic/gin"
)

func TagValidators() []func(c *gin.Context, tag dbCon.Tag) {
	return []func(c *gin.Context, tag dbCon.Tag){
		ValidateName,
	}
}

func ValidateName(c *gin.Context, tag dbCon.Tag) {
	if tag.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag name is cannot be empty"})
		return
	}
}
