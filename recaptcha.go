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

// SecretKey stores the secret key for your reCAPTCHA. You can get your keys
// here: https://www.google.com/recaptcha/admin/
var SecretKey string

// validates sends the request to the Google API server and gets the results
// about the whole recaptcha process, that means the boolean value which is
// then returned from this function to be processed
func validate(grr string, ip string) (bool, error) {
	// google validation request data
	postURL := "https://www.google.com/recaptcha/api/siteverify"
	postStr := url.Values{
		"secret":   {SecretKey},
		"response": {grr},
		"remoteip": {ip},
	}

	// validity check
	responsePost, err := http.PostForm(postURL, postStr)
	if err != nil {
		return false, err
	}
	defer responsePost.Body.Close()
	body, err := ioutil.ReadAll(responsePost.Body)
	if err != nil {
		return false, err
	}

	// unmarshal the response and test the success
	gres := googleResponse{}
	json.Unmarshal(body, &gres)
	if !gres.Success {
		return false, nil
	}

	return true, nil
}

// Middleware is a fiber middleware for validating recaptchas
func Middleware(c *fiber.Ctx) error {
	// g-recaptcha-response field extraction and testing
	req := request{}
	c.BodyParser(&req)
	if req.GRecaptchaResponse == "" {
		c.Locals("recaptchaSuccess", false)
		return c.Next()
	}

	if len(req.GRecaptchaResponse) <= 500 {
		c.Locals("recaptchaSuccess", false)
		return c.Next()
	}

	// get the google validation success value (true = success)
	ok, err := validate(req.GRecaptchaResponse, c.IP())
	if err != nil {
		fmt.Println(err.Error())
	}

	// store the result and execute the next method in router
	c.Locals("recaptchaSuccess", ok)
	return c.Next()
}
