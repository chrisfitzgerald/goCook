package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
   e := echo.New()
   e.POST("/login", loginHandler)
   
   renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
   e.Renderer = renderer

   e.GET("/", loginHandler)
   e.GET("/home", homeHandler)

   err := e.Start(":8080")
   if err != nil {
      log.Fatalf("failed to start server: %v", err)
   }
} 

