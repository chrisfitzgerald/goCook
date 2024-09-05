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
    now := time.Now()
    formattedNow := now.Format("Mon Jan 2 15:04:05 MST 2006")
    
    data := map[string]interface{}{
        "DateTime": formattedNow,
    }

    user := c.Get("user")
    log.Println("User object retrieved from context")

    if claims, ok := user.(*jwt.RegisteredClaims); ok {
        if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
            log.Println("Token has expired")
            data["Email"] = ""
        } else {
            data["Email"] = claims.Subject
            log.Println("User authenticated")
        }
    } else {
        log.Println("User not authenticated")
        data["Email"] = ""
    }

    log.Println("Rendering home page")
    return c.Render(http.StatusOK, "index.html", data)
}