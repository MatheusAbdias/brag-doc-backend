package controllers

import (
	"context"
	"net/http"
	"strconv"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TagRepo interface {
	CreateTag(c context.Context, name string) error
	GetTags(c context.Context, arg dbCon.GetTagsParams) ([]dbCon.Tag, error)
	GetTag(c context.Context, id uuid.UUID) (dbCon.Tag, error)
	UpdateTag(c context.Context, arg dbCon.UpdateTagParams) error
	DeleteTag(c context.Context, id uuid.UUID) error
}

type Controller struct {
	Repo TagRepo
}

func (ctrl *Controller) CreateTag(c *gin.Context) {
	var tag dbCon.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validators := TagValidators()
	for _, validator := range validators {
		validator(c, tag)
	}

	err := ctrl.Repo.CreateTag(c, tag.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (ctrl *Controller) ListTags(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset := (page - 1) * pageSize
	arg := dbCon.GetTagsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	tags, err := ctrl.Repo.GetTags(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page": page,
		"tags": tags,
	})
}

func (ctrl *Controller) GetTag(c *gin.Context) {
	tagID := c.Param("id")

	id, err := uuid.Parse(tagID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return

	}

	tag, err := ctrl.Repo.GetTag(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (ctrl *Controller) UpdateTag(c *gin.Context) {
	tagID := c.Param("id")
	id, err := uuid.Parse(tagID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Tag not found"})
		return
	}

	var tag dbCon.UpdateTagParams
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Invalid data"})
		return
	}

	tag.ID = id

	err = ctrl.Repo.UpdateTag(c, tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": tag})
}

func (ctrl *Controller) DeleteTag(c *gin.Context) {
	tagID := c.Param("id")

	id, err := uuid.Parse(tagID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Tag not found"})
		return
	}

	err = ctrl.Repo.DeleteTag(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Cannot delete tag"})
		return
	}

	c.Status(http.StatusNoContent)
}
