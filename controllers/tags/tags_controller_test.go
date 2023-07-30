package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dbCon "github.com/MatheusAbdias/brag-doc-backend/internal/db/tags"
	"github.com/gin-gonic/gin"
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

func TestTagControllerCreateTag(t *testing.T) {
	testCases := []struct {
		Name          string
		Repo          TagRepo
		RequestBody   interface{}
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name: "Valid Tag Creation",
			Repo: &FakeTagRepo{},
			RequestBody: map[string]string{
				"name": "Test",
			},
			ExpectedError: false,
			ExpectedCode:  http.StatusCreated,
		},
		{
			Name: "Invalid Tag missing name",
			Repo: &FakeTagRepo{},
			RequestBody: map[string]string{
				"name": "",
			},
			ExpectedError: true,
			ExpectedCode:  http.StatusBadRequest,
		},
		{
			Name:          "Invalid Tag Payload missing name",
			Repo:          &FakeTagRepo{},
			RequestBody:   map[string]string{},
			ExpectedError: true,
			ExpectedCode:  http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tagController := TagController{
				Repo: tc.Repo,
			}

			requestBody, _ := json.Marshal(tc.RequestBody)
			req, _ := http.NewRequest("POST", "/tags", bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req

			tagController.CreateTag(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}
}

func TestTagControllerUdpate(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          TagRepo
		RequestBody   interface{}
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:  "Update Tag",
			TagID: uuid.NewString(),
			Repo:  &FakeTagRepo{},
			RequestBody: map[string]string{
				"name": "Test",
			},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:  "Invalid Tag id",
			TagID: "1",
			Repo:  &FakeTagRepo{},
			RequestBody: map[string]string{
				"name": "Test",
			},
			ExpectedError: false,
			ExpectedCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tagController := TagController{
				Repo: tc.Repo,
			}

			requestBody, _ := json.Marshal(tc.RequestBody)
			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			tagController.UpdateTag(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}

}
func TestTagControllerDelete(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          TagRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "Delete Tag",
			TagID:         uuid.NewString(),
			Repo:          &FakeTagRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:          "Invalid Tag id",
			TagID:         "1",
			Repo:          &FakeTagRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tagController := TagController{
				Repo: tc.Repo,
			}

			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("DELETE", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			tagController.DeleteTag(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}

}

func TestTagControllerList(t *testing.T) {
	testCases := []struct {
		Name          string
		Repo          TagRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "List Tags",
			Repo:          &FakeTagRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tagController := TagController{
				Repo: tc.Repo,
			}

			req, _ := http.NewRequest("GET", "/tags", nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			tagController.ListTags(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}

}

func TestTagControllerRetrive(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          TagRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "Get Tags",
			TagID:         uuid.NewString(),
			Repo:          &FakeTagRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:          "Invalid Tag id",
			TagID:         "1",
			Repo:          &FakeTagRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tagController := TagController{
				Repo: tc.Repo,
			}

			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("GET", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			tagController.GetTag(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d, got %d", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}

}
