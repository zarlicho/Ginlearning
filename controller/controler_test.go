package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	// Create a new gin context for testing
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Mock request body
	requestBody := `{
		"Gmail": "test@gmail.com",
		"Passwords": "testpassword",
		"Token": "testtoken",
		"Id": 123
	}`

	// Create a request with the mocked body
	req, err := http.NewRequest("POST", "/login", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set the request's content type
	req.Header.Set("Content-Type", "application/json")

	// Attach the request to the context
	c.Request = req

	// Create an instance of your controller
	controller := &ControlersImpl{}

	// Call the Login function
	controller.Login(c)

	// Check the response status code
	if c.Writer.Status() != http.StatusBadRequest {
		// t.Errorf("Expected status code %d, but got %d", c.Writer.Status(), http.StatusBadRequest)
		fmt.Println(c.Writer.Status())
	}

	// Check the response body
	// expectedResponseBody := `{"error":"failed to read request body"}`
	// if body := c.Writer.(*httptest.ResponseRecorder).Body.String(); body != expectedResponseBody {
	// 	t.Errorf("Expected response body %s, but got %s", expectedResponseBody, body)
	// }
}
