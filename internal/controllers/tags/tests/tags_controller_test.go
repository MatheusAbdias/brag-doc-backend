package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/tags"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestControllerCreateTag(t *testing.T) {
	testCases := []struct {
		Name          string
		Repo          controllers.TagRepo
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
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			requestBody, _ := json.Marshal(tc.RequestBody)
			req, _ := http.NewRequest("POST", "/tags", bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req

			Controller.CreateTag(ctx)

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

func TestControllerUdpate(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          controllers.TagRepo
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
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			requestBody, _ := json.Marshal(tc.RequestBody)
			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			Controller.UpdateTag(ctx)

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
func TestControllerDelete(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          controllers.TagRepo
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
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("DELETE", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			Controller.DeleteTag(ctx)

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

func TestControllerList(t *testing.T) {
	testCases := []struct {
		Name          string
		Repo          controllers.TagRepo
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
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			req, _ := http.NewRequest("GET", "/tags", nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			Controller.ListTags(ctx)

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

func TestControllerRetrive(t *testing.T) {
	testCases := []struct {
		Name          string
		TagID         string
		Repo          controllers.TagRepo
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
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			url := "/tags/" + tc.TagID
			req, _ := http.NewRequest("GET", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.TagID}}
			Controller.GetTag(ctx)

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
