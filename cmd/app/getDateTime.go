package main

import (
    "net/http"
    "time"
    "github.com/labstack/echo/v4"
)

func getDateTimeHandler(c echo.Context) error {
    // Get the current date and time
    now := time.Now()
    // Format it as a string (e.g., "Mon Jan 2 15:04:05 MST 2006")
    formattedNow := now.Format("Mon Jan 2 15:04:05 MST 2006")
    // Pass the formatted date and time to the template
    data := map[string]interface{}{
        "DateTime": formattedNow,
    }
    return c.Render(http.StatusOK, "index.html", data)
}