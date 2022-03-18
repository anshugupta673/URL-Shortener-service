package routes

import (
	"time"
	"github.com/anshugupta673/URL-Shortener-service/helpers"
)

type request struct { /* golang do not nderstand json on it's own so it has to do a lot of encoding and decoding with json, searilization bascially, so we need to tell what it's going to look like when we come across json when we receive a request */
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"` /* we do not want frontend to make unlimited number of requests */
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	//implement rate limiting

	//check if the input sent by the user if an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid URL"})
	}

	//check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavaliable).JSON()
	}

	//enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}
