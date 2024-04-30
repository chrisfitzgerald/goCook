package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	Username string
	Password string 
}

var users = []User{
	{Username: "admin", Password: "password"},
}

func loginHandler(c echo.Context) error {
	isLoggedIn := false
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Validate the user's credentials.
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return c.Redirect(http.StatusSeeOther, "/home")
		}
	}
	

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