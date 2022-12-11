package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const currentUser = "currentUser"
const secretKey = "secret"

func GenerateJWT(c *fiber.Ctx, id int) error {
	claims := &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), Issuer: strconv.Itoa(id)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(secretKey)); err != nil {
		return err
	} else {
		cookie := fiber.Cookie{Name: currentUser, Value: tokenString, Expires: time.Now().Add(time.Hour * 24), HTTPOnly: true}
		c.Cookie(&cookie)
		return err
	}
}

func ValidateJWT(c *fiber.Ctx) (int, error) {
	tokenString := c.Cookies(currentUser)
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	}); err != nil {
		return 0, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims["iss"].(string))
		return id, err
	} else {
		return 0, errors.New("token invalid")
	}
}

func ExpireCookie(c *fiber.Ctx) {
	cookie := fiber.Cookie{Name: currentUser, Value: "", Expires: time.Now().Add(-time.Hour), HTTPOnly: true}
	c.Cookie(&cookie)
}
