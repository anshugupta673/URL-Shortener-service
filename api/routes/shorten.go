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

/* POST request... */
func ShortenURL(c *fiber.Ctx) error {
	body := new(request) /* have a new request(empty), body of type request and use body-parser to make sense of the request received from the user to convert into struct which is understood by golang */

	if err := c.BodyParser(&body); err != nil { /* body parser->to parse the json that we get as a part of the request into struct which is understood by the golang */
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	//implement rate limiting
	//...we check the IP of the user, and we check if IP is stored in the database, if IP is stored in the database then that means user has already used our service, if yes then devrement the number of rateRemaining by 1

	//check if the input sent by the user if an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid URL"})
	}

	//also there's something that the user can do, that user can use local host 3000 as the URL to get shortened and then program can enter a infinite loop and we dont want that to happen(so we need some helper functions to hold that... so we check for domain error also we enforce https or ssl)

	//check for domain error, stop the infinite loop
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavaliable).JSON(fiber.Map{"...": "..."})
	}

	//enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL) 
}
