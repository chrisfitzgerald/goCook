package main

import (
   "os"
   "net/http"
	"html/template"
	"io"
	"log"
   "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
   "github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func serveLoginPage(c echo.Context) error {
    // Render the login page
    return c.Render(http.StatusOK, "login.html", nil)
}

func main() {
   // Create a new Echo instance
   e := echo.New()

   //Set up the template renderer
   templates, err := template.ParseGlob("template/*.html")
   if err != nil {
      log.Fatalf("failed to parse templates: %v", err)
   }
   renderer := &TemplateRenderer{
      templates: templates,
   }
   e.Renderer = renderer
    
   // Define the JWT middleware configuration
   config := middleware.JWTConfig{
      Claims:     &jwt.MapClaims{},
      SigningKey: jwtKey,
      TokenLookup: "cookie:token",
   }

   //set up the routes
   setupRoutes(e, config)

   // Start the server
   port := os.Getenv("PORT")
   if port == "" {
       port = "8080" // Default port if not specified
   }
   err = e.Start(":" + port)
   if err != nil {
       log.Fatalf("failed to start server: %v", err)
   }
} 

func setupRoutes(e *echo.Echo, config middleware.JWTConfig) {
    // Define the login route
    e.POST("/login", loginHandler)

    // Define the home route with the JWT middleware
    e.GET("/home", homeHandler, middleware.JWTWithConfig(config))

    // Define the root route
    e.GET("/", serveLoginPage)
}
