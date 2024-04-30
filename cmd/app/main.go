package main
	
import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
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

   e.Start(":8080")
} 

