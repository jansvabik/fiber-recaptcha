# Golang Fiber reCAPTCHA middleware

[![Go Report Card](https://goreportcard.com/badge/github.com/jansvabik/fiber-recaptcha)](https://goreportcard.com/report/github.com/jansvabik/fiber-recaptcha)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/jansvabik/fiber-recaptcha/blob/master/LICENSE)
[![Maintainability](https://api.codeclimate.com/v1/badges/06a33ed30e237fa413ee/maintainability)](https://codeclimate.com/github/jansvabik/fiber-recaptcha/maintainability)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/jansvabik/fiber-recaptcha/issues)

Golang reCAPTCHA middleware for [Fiber](https://github.com/gofiber/fiber). This package takes care of getting the `g-recaptcha-response` field from the sent form (or other request) data and validating the reCAPTCHA using [Google's API](https://developers.google.com/recaptcha/docs/verify).

This package should support at least these types of reCAPTCHA:

* reCAPTCHA V2, "I'm not a robot" Checkbox
* reCAPTCHA V2, Invisible reCAPTCHA badge

## Installation
You can add this package to your code by two simple steps, which are:

1. get the package by `go get github.com/jansvabik/fiber-recaptcha`
2. import the package manually by importing `github.com/jansvabik/fiber-recaptcha` or by calling the exported `Middleware(c *fiber.Ctx) error` function from inside of the package

## Usage
To use the middleware within your web server, you should add the `Middleware` function to the queue of functions in your Fiber router, like in this example:

```go
// pass your secret key to the package
recaptcha.SecretKey = "place-your-recaptcha-secret-key-here"

// create new fiber router with one endpoint
router := fiber.New()
router.Post("/endpoint", recaptcha.Middleware, endpoint.YourHandler)
```

When there is a request to this endpoint, for first it will go through this reCAPTCHA middleware which sends a request to the Google API and validates the whole request to be sent by human or by robot (depending to the Google response). Then, it sets up a local variable `recaptchaSuccess` with boolean value which you can access and use in an `if` statement.

```go
// test the recaptcha success in your endpoint handler
if c.Locals("recaptchaSuccess") == false {
    // code to do if the validation fails
}

// do other job if the validation succeeds
```

There are only two possible values which can appear in the local variable:

* boolean `true` if the validation succeeds
* boolean `false` if the validation fails
