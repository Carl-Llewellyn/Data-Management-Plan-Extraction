package postgres_routes_test

import (
	"dmp-api/api/postgres_routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSubmitDataCollection(t *testing.T) {
	// Set up the Fiber app
	app := fiber.New()

	// Set up the route
	app.Post("/submit", postgres_routes.SubmitDataCollection)

	// Test successful form submission
	req := httptest.NewRequest("POST", "/submit", strings.NewReader("uuid_fk=test_uuid&username=test_username&powerpoint_file_name=test.pptx&is_this_new_data=yes&dataset_name=test_dataset&data_storage_requirements=test_requirements"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test form submission with missing data
	req = httptest.NewRequest("POST", "/submit", strings.NewReader("uuid_fk=test_uuid"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	// Test form submission with invalid data
	req = httptest.NewRequest("POST", "/submit", strings.NewReader("uuid_fk=invalid_uuid&username=test_username&powerpoint_file_name=test.pptx&is_this_new_data=yes&dataset_name=test_dataset&data_storage_requirements=test_requirements"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}
