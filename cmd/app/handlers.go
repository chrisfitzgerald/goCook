package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
	"log"
)

type User struct {
	Username string
	Password string 
}

var users = []User{
	{Username: "admin", Password: "password"},
}

var jwtKey = []byte("secret")

// validateUser checks if the user exists in the users slice
func validateUser(username, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// LoginHandler is a handler for the login route
func loginHandler(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    // Validate the user's credentials.
    if !validateUser(username, password) {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "message": "user unauthorized",
        })
    }

    // Create a new token object, specifying signing method
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Minute * 5).Unix()	

    //Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        log.Printf("Error signing token: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "could not log in",
        })
    }

    //return the token
    return c.JSON(http.StatusOK, map[string]string{
        "token": tokenString,
        "redirectURL": "/home",
    })
}

func homeHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Welcome back " + c.FormValue("username") + "!")
}