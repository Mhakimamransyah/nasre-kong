package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const secrets = "secret" // JWT secrets
const issuer = "nasre"   // JWT issuer

func main() {

	app := fiber.New()

	v1 := app.Group("/api").Group("/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {

		username := c.FormValue("username")
		password := c.FormValue("password")

		if username != "admin" || password != "admin" {
			return c.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}

		claims := jwt.MapClaims{
			"iss": issuer,
			"exp": time.Now().Add(time.Hour * 2).Unix(),
			"iat": time.Now().Unix(),
			"nbf": time.Now().Add(time.Hour * 2).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(secrets))

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "login success",
			"token":   t,
		})
	})

	app.Listen(":5000")
}
