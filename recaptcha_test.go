package recaptcha

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// go test -run TestValidate
func TestValidate(t *testing.T) {
	tests := []struct {
		grr string
		ip  string
	}{
		{"", ""},
		{"no", "127.0.0.1"},
		{"", "1.2.3.4"},
		{"abcd", "2.3.4.5"},
	}
	for _, test := range tests {
		ok, err := validate(test.grr, test.ip)
		assert.False(t, ok, "Validation success")
		assert.Nil(t, err, "Validation error")
	}
}

// go test -run TestMiddleware
func TestMiddleware(t *testing.T) {
	app := fiber.New()
	app.Use(Middleware)
	app.Post("/test", func(c *fiber.Ctx) error {
		assert.Equal(t, false, c.Locals("recaptchaSuccess"))
		return nil
	})

	tests := []struct {
		body        string
		contentType string
	}{
		{"{\"g-recaptcha-response\":\"a\"}", "application/json"},
		{"{\"g-recaptcha-response\":\"\"}", "application/json"},
		{"", "application/json"},
		{"g-recaptcha-response=a", "application/x-www-form-urlencoded"},
		{"g-recaptcha-response=", "application/x-www-form-urlencoded"},
		{"", "application/x-www-form-urlencoded"},
	}
	for _, test := range tests {
		req := httptest.NewRequest("POST", "/test", strings.NewReader(test.body))
		req.Header.Add("Content-Type", test.contentType)
		resp, err := app.Test(req)
		assert.Nil(t, err, "Returned error")
		assert.Equal(t, 200, resp.StatusCode, "Status code")
	}
}
