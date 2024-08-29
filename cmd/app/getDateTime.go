package main

import (
    "net/http"
    "time"
    "github.com/labstack/echo/v4"
    "log" // Add this import
)

func getDateTimeHandler(c echo.Context) error {
    log.Println("getDateTimeHandler called") // Log when the handler is called

    // Get the current date and time
    now := time.Now()
    // Format it as a string (e.g., "Mon Jan 2 15:04:05 MST 2006")
    formattedNow := now.Format("Mon Jan 2 15:04:05 MST 2006")
    // Pass the formatted date and time to the template
    data := map[string]interface{}{
        "DateTime": formattedNow,
    }
    // Attempt to render the template and log any errors
    if err := c.Render(http.StatusOK, "index.html", data); err != nil {
        log.Printf("Error rendering template: %v", err) // Log the error
        return err // Return the error to the client
    }
    log.Println("Template rendered successfully") // Log successful rendering
    return nil // Return nil if no error occurred
}