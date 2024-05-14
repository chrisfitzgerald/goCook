package main

import (
    "fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
    // Get the JWT from the cookie.
    cookie, err := c.Cookie(TokenCookieName)
    if err != nil {
		return handleError(c, http.StatusUnauthorized, "missing jwt")
    }

    // Parse the JWT.
    token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method is what we expect.
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
		
        return []byte(jwtKey), nil
    })

    if err != nil {
		return handleError(c, http.StatusUnauthorized, "invalid jwt")
    }

    // Check if the token is valid.
    if !token.Valid {
       return handleError(c, http.StatusUnauthorized, "invalid jwt") 
    }

    // Get the username from the JWT.
    //claims := token.Claims.(jwt.MapClaims)
    //username := claims["username"].(string)

    // Return the welcome message.
    return c.Redirect(http.StatusMovedPermanently, "./template/index.html")
}