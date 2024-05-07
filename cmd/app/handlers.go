package main

import (
	"log"
    "fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
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

    //Set a cookie with the JWT
    cookie := new(http.Cookie)
    cookie.Name = "token"
    cookie.Value = tokenString
    cookie.Expires = time.Now().Add(time.Minute * 5)
    cookie.Path = "/"
    cookie.Secure = false
    cookie.SameSite = http.SameSiteLaxMode 
    c.SetCookie(cookie)

    //return the token
    return c.Redirect(http.StatusMovedPermanently, "/home")
}

func homeHandler(c echo.Context) error {
    // Get the JWT from the cookie.
    cookie, err := c.Cookie("token")
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "message": "missing jwt",
        })
    }

    // Parse the JWT.
    token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "message": "malformed jwt",
        })
    }

    // Check if the token is valid.
    if !token.Valid {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "message": "invalid jwt",
        })
    }

    // Get the username from the JWT.
    claims := token.Claims.(jwt.MapClaims)
    username := claims["username"].(string)

    // Return the welcome message.
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Welcome, " + username + "!",
    })
}