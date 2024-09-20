package main

import (
   "net/http"
	"html/template"
	"io"
	"log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	_"github.com/labstack/echo-jwt/v4"
	_"github.com/labstack/echo/v4/middleware"
	_"golang.org/x/oauth2"
	_"golang.org/x/oauth2/google"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

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

func customJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        if c.Get("jwtLogFlag") == nil {
            cookie, err := c.Cookie(TokenCookieName)
            if err != nil {
                log.Printf("No JWT cookie found")
                c.Set("jwtLogFlag", true)
                return next(c)
            }

            log.Printf("JWT cookie found")

            token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
                return []byte(JWTKey), nil
            })

            if err != nil {
                log.Printf("JWT parsing error")
                // Clear the invalid cookie
                c.SetCookie(&http.Cookie{
                    Name:   TokenCookieName,
                    Value:  "",
                    Path:   "/",
                    MaxAge: -1,
                })
                return next(c)
            }

            if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
                c.Set("user", claims)
                log.Printf("Valid JWT found for user")
            } else {
                log.Printf("Invalid JWT token")
                // Clear the invalid cookie
                c.SetCookie(&http.Cookie{
                    Name:   TokenCookieName,
                    Value:  "",
                    Path:   "/",
                    MaxAge: -1,
                })
            }
        }
        return next(c)
    }
}

func debugHandler(c echo.Context) error {
   data := map[string]interface{}{
      "Cookies": c.Cookies(),
      "Headers": c.Request().Header,
   }

   if user := c.Get("user"); user != nil {
      data["User"] = user
   }

   return c.JSON(http.StatusOK, data)
}

func main() {
   // Create a new Echo instance
   e := echo.New()

   // Serve files from the template directory
   e.Static("/template", "template")
   // Load templates
   t := template.Must(template.ParseGlob("template/*.html"))
   e.Renderer = &TemplateRenderer{
       templates: t,
   }

   // Serve files from the static directory
   e.Static("/static", "static")


   // Apply JWT middleware to all routes, but allow requests to continue on error
   e.Use(customJWTMiddleware)

   // Public routes
   e.GET("/", homeHandler)
   e.GET("/login", serveLoginPage)
   e.GET("/auth/google/login", googleLoginHandler)
   e.GET("/auth/google/callback", googleCallbackHandler)
   e.GET("/logout", logoutHandler)
   e.GET("/game", gameHandler)
   // Protected routes can be added here if needed
   // r := e.Group("/protected")
   // r.Use(echojwt.WithConfig(config))
   // r.GET("/profile", profileHandler)

   e.GET("/debug", debugHandler)

   e.Logger.Fatal(e.Start(":8080"))
}

func gameHandler(c echo.Context) error {
    user := c.Get("user")
    if user == nil {
        return c.Redirect(http.StatusSeeOther, "/login") // Redirect to login if not authenticated
    }

    // Render the feature1 page with user data
    claims := user.(*jwt.RegisteredClaims)
    data := map[string]interface{}{
        "Email": claims.Subject,
    }
    return c.Render(http.StatusOK, "game.html", data)
}
