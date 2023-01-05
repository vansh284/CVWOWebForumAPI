package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vansh284/CVWOWebForumAPI/pkg/config"
)

var envMap map[string]string = config.GetEnvMap()

func GenerateJWT(c *fiber.Ctx, id int) error {
	claims := &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), Issuer: strconv.Itoa(id)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(envMap["SECRET_KEY"])); err != nil {
		return err
	} else {
		cookie := fiber.Cookie{
			Name:     envMap["CURRENT_USER"],
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
			Domain:   "cvwo-web-forum.onrender.com",
			Secure:   true,
			SameSite: "None"}
		c.Cookie(&cookie)
		return err
	}
}

func ValidateJWT(c *fiber.Ctx) (int, error) {
	tokenString := c.Cookies(envMap["CURRENT_USER"])
	fmt.Println(tokenString)
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(envMap["SECRET_`KEY"]), nil
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
	cookie := fiber.Cookie{Name: envMap["CURRENT_USER"], Value: "", Expires: time.Now().Add(-time.Hour), HTTPOnly: true}
	c.Cookie(&cookie)
}
