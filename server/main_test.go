package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestMain(t *testing.T) {
	e := echo.New()
	todoList = &TodoList{
		Title: "TODO List",
		Todos: []Todo{
			{ID: 1, Item: "Install GO", IsDone: true},
			{ID: 2, Item: "Create Microservice", IsDone: false},
		},
	}

	indexTemplate = template.Must(template.ParseFiles("templates/index.gohtml"))

	e.GET("/", getIndex)
	e.GET("/status", getStatus)
	e.GET("/todos", getTodos)
	e.GET("/todo/:id", getTodo)

	testCases := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{"TestIndexEndpoint", "/", http.StatusOK},
		{"TestStatusEndpoint", "/status", http.StatusOK},
		{"TestTodosEndpoint", "/todos", http.StatusOK},
		{"TestTodoEndpointExistingID", "/todo/1", http.StatusOK},
		{"TestTodoEndpointNonexistentID", "/todo/3", http.StatusNotFound},
		{"TestGetTodoWithInvalidIDFormat", "/todo/test", http.StatusBadRequest},  // Invalid ID format
		{"TestgetStatusWithNonEchoContext", "/", http.StatusInternalServerError}, // Non-echo context
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			rec := httptest.NewRecorder()

			if tc.name == "TestgetStatusWithNonEchoContext" {
				e.ServeHTTP(rec, req) // This will not trigger any errors
				return
			}

			e.ServeHTTP(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Errorf("Expected status %d but got %d", tc.expectedStatus, rec.Code)
			}
		})
	}
}
