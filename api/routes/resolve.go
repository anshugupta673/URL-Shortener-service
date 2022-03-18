package routes

import (
	"github.com/anshugupta673/URL-Shortener-service/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

/* ResolveURL: once we use the shorten.go file(shorten function to create a shorten URL of the URL that we put) so bascially we put the man URL and we get a shorten URL(that's what our service does, but when somebody uses that shorten URL then it needs to redirect to the actual link, so what we are doing is we are saving the real link(long link) in our database and we create a short link and when somebody uses the short url link we check our database which is the actual link which has a relationship with this link and then we take the user to the sctual link) */
func ResolveURL(c *fiber.Ctx) error { /* http request/response context */
	url := c.Params("url")        /* url is the variable we get access to using the c.Params */
	r := database.CreateClient(0) /* db number 0 */
	defer r.Close()               /* runs when the call stack is at the end */

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to the db"}))
	}

	rInr := database.CreateClient(1) /* increment the counter to 1 */
	defer rInr.Close()

	_ = rInr.Incr(databse.Ctx, "counter")

	return c.Redirect(value, 301) /* with status 301 we redirect the uset to the actual URL what we just find in the database */
}
