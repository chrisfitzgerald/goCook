
package main

import (
    "log"
    "net/http"
    "time"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
)

const (
    jwtKey        = "secret"
    TokenCookieName = "token"
)

func createJWT(username string) (string, error) {
    // Create a new token object, specifying signing method
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

    // Sign and get the complete encoded token as a string using the secret
    return token.SignedString([]byte(jwtKey))
}

func validateUser(username, password string) bool {
    const validUsername = "admin"
    const validPassword = "password"

    return username == validUsername && password == validPassword
}

func handleError(c echo.Context, status int, message string) error {
    log.Printf("Error: %s", message)
    return c.JSON(status, map[string]string{
        "message": message,
    })
}

func loginHandler(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    // Validate the user's credentials.
    if !validateUser(username, password) {
        return handleError(c, http.StatusUnauthorized, "user unauthorized")
    }

    tokenString, err := createJWT(username)
    if err != nil {
        return handleError(c, http.StatusInternalServerError, "could not log in")
    }

    // Set a cookie with the JWT
    cookie := new(http.Cookie)
    cookie.Name = TokenCookieName
    cookie.Value = tokenString
    cookie.Expires = time.Now().Add(time.Minute * 5)
    cookie.Path = "/"
    cookie.HttpOnly = true
    cookie.Secure = true 
    cookie.SameSite = http.SameSiteLaxMode
    c.SetCookie(cookie)

    // Redirect to home
    return c.Redirect(http.StatusMovedPermanently, "/home")
}