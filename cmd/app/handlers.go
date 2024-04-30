package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func loginHandler(c echo.Context) error {
    // Check if user is already logged in
    // This is just a placeholder, replace with your actual login check
    isLoggedIn := false

    if isLoggedIn {
        // If user is already logged in, redirect to another page
        return c.Redirect(http.StatusSeeOther, "/home")
    } else {
        // If user is not logged in, serve the login page
        return c.Render(http.StatusOK, "login.html", nil)
    }
}

func homeHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Home Page")
}