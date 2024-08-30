package main

import (
    _"fmt"
    "log"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"time"
)

func homeHandler(c echo.Context) error {
	log.Println("Home handler started")
	// Get the current date and time
	now := time.Now()
	formattedNow := now.Format("Mon Jan 2 15:04:05 MST 2006")
	
	data := map[string]interface{}{
		"DateTime": formattedNow,
	}

	// Try to get the user from the context (it will be nil if not authenticated)
	user := c.Get("user")
	log.Printf("User object from context: %+v", user)

	if claims, ok := user.(*jwt.RegisteredClaims); ok {
		log.Printf("JWT Claims: %+v", claims)
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			log.Println("Token has expired")
			data["Email"] = ""
		} else {
			email := claims.Subject
			data["Email"] = email
			log.Printf("User authenticated: %s", email)
		}
	} else {
		log.Println("User not authenticated")
		// Log the cookies
		for _, cookie := range c.Cookies() {
			log.Printf("Cookie: %s = %s", cookie.Name, cookie.Value)
		}
		data["Email"] = "" // Ensure Email is set to an empty string for unauthenticated users
	}

	return c.Render(http.StatusOK, "index.html", data)
}