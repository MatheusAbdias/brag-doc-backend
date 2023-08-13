package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	controllers "github.com/MatheusAbdias/brag-doc-backend/internal/controllers/events"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func TestControllerCreateEvent(t *testing.T) {
	testCases := []struct {
		Name          string
		Repo          controllers.EventsRepo
		RequestBody   interface{}
		ExpectedCode  int
		ExpectedError bool
	}{
		{
			Name: "Valid Event",
			Repo: &FakeEventRepo{},
			RequestBody: map[string]interface{}{
				"name":        "New event",
				"description": "New valid event",
				"date":        time.Now().Format(time.RFC3339),
			},
			ExpectedCode:  http.StatusCreated,
			ExpectedError: false,
		},
		{
			Name: "Invalid Event missing",
			Repo: &FakeEventRepo{},
			RequestBody: map[string]interface{}{
				"description": "New valid event",
				"date":        time.Now().Format(time.RFC3339),
			},
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: true,
		},
		{
			Name: "Invalid event missing description",
			Repo: &FakeEventRepo{},
			RequestBody: map[string]interface{}{
				"name": "New event",
				"date": time.Now().Format(time.RFC3339),
			},
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			requestBody, _ := json.Marshal(tc.RequestBody)
			req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req

			Controller.CreateEvent(ctx)

			if tc.ExpectedError {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d got %d.", tc.ExpectedCode, recorder.Code)

				}
			} else {
				if recorder.Code != tc.ExpectedCode {
					t.Errorf("Expected status code %d got %d.", tc.ExpectedCode, recorder.Code)
				}
			}
		})
	}
}

func TestControllerUdpate(t *testing.T) {
	testCases := []struct {
		Name          string
		EventID       string
		Repo          controllers.EventsRepo
		RequestBody   interface{}
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:    "Update Event",
			EventID: uuid.NewString(),
			Repo:    &FakeEventRepo{},
			RequestBody: map[string]string{
				"name": "Test",
			},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:    "Invalid Event id",
			EventID: "1",
			Repo:    &FakeEventRepo{},
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
			url := "/events/" + tc.EventID
			req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.EventID}}
			Controller.UpdateEvent(ctx)

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
		EventID       string
		Repo          controllers.EventsRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "Delete Event",
			EventID:       uuid.NewString(),
			Repo:          &FakeEventRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:          "Invalid Event id",
			EventID:       "1",
			Repo:          &FakeEventRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			url := "/events/" + tc.EventID
			req, _ := http.NewRequest("DELETE", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.EventID}}
			Controller.DeleteEvent(ctx)

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
		Repo          controllers.EventsRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "List Events",
			Repo:          &FakeEventRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			req, _ := http.NewRequest("GET", "/events", nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			Controller.ListEvents(ctx)

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
		EventID       string
		Repo          controllers.EventsRepo
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Name:          "Get Events",
			EventID:       uuid.NewString(),
			Repo:          &FakeEventRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusOK,
		},
		{
			Name:          "Invalid Event id",
			EventID:       "1",
			Repo:          &FakeEventRepo{},
			ExpectedError: false,
			ExpectedCode:  http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			Controller := controllers.Controller{
				Repo: tc.Repo,
			}

			url := "/events/" + tc.EventID
			req, _ := http.NewRequest("GET", url, nil)
			recorder := httptest.NewRecorder()

			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = req
			ctx.Params = []gin.Param{{Key: "id", Value: tc.EventID}}
			Controller.GetEvent(ctx)

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
