package recaptcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

// request is a structure for extracting g-recaptcha-response field from request
type request struct {
	GRecaptchaResponse string `json:"g-recaptcha-response" form:"g-recaptcha-response"`
}

// googleResponse is a structure with Google validation response data
type googleResponse struct {
	Success    bool     `json:"success"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

// Middleware is a fiber middleware for validating recaptchas
func Middleware(secretKey string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// g-recaptcha-response field extraction and testing
		req := request{}
		c.BodyParser(&req)
		if req.GRecaptchaResponse == "" {
			c.Locals("recaptchaSuccess", false)
			return c.Next()
		}

		// google validation request data
		postURL := "https://www.google.com/recaptcha/api/siteverify"
		postStr := url.Values{
			"secret":   {secretKey},
			"response": {req.GRecaptchaResponse},
			"remoteip": {c.IP()},
		}

		// validity check
		responsePost, err := http.PostForm(postURL, postStr)
		if err != nil {
			fmt.Println(err.Error())
			c.Locals("recaptchaSuccess", false)
			return c.Next()
		}
		defer responsePost.Body.Close()
		body, err := ioutil.ReadAll(responsePost.Body)
		if err != nil {
			fmt.Println(err.Error())
			c.Locals("recaptchaSuccess", false)
			return c.Next()
		}

		// unmarshal the response and test the success
		gres := googleResponse{}
		json.Unmarshal(body, &gres)
		if !gres.Success {
			c.Locals("recaptchaSuccess", false)
			return c.Next()
		}

		// success, execute the next method in router
		c.Locals("recaptchaSuccess", true)
		return c.Next()
	}
}
