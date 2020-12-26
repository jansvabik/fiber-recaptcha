# Golang Fiber reCAPTCHA middleware

[![Go Report Card](https://goreportcard.com/badge/github.com/jansvabik/fiber-recaptcha)](https://goreportcard.com/report/github.com/jansvabik/fiber-recaptcha)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/jansvabik/fiber-recaptcha/blob/master/LICENSE)
[![Maintainability](https://api.codeclimate.com/v1/badges/06a33ed30e237fa413ee/maintainability)](https://codeclimate.com/github/jansvabik/fiber-recaptcha/maintainability)

Golang reCAPTCHA middleware for [Fiber](https://github.com/gofiber/fiber). This package takes care of getting the `g-recaptcha-response` field from the sent form (or other request) data and validating the reCAPTCHA using Google's API.

This package should support at least these types of reCAPTCHA:

* reCAPTCHA V2, "I'm not a robot" Checkbox
* reCAPTCHA V2, Invisible reCAPTCHA badge

## Installation
You can add this package to your code by two simple steps, which are:

1. get the package by `go get github.com/jansvabik/fiber-recaptcha`
2. import the package manually by `import "github.com/jansvabik/fiber-recaptcha"` or by calling the exported `Middleware(secretKey string)` function from inside of the package
