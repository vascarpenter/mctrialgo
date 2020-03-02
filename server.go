package main

import (
	"context"
	"html/template"
	"io"
	"mctrialgo/models"
	"mctrialgo/routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
	templates *template.Template
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func setUserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := routes.Repository()
			defer db.Close()

			if session, err := session.Get("oursession", c); err == nil {
				if userid, ok := session.Values["userid"]; ok {
					useridint := userid.(uint)
					hosp, _ := models.Hospitals(models.HospitalWhere.HospitalID.EQ(useridint), qm.Limit(1)).All(context.Background(), db)
					if hosp != nil {
						c.Set("UserName", hosp[0].Name)
						c.Set("UserID", useridint)
					}
				}
			}

			return next(c)
		}
	}
}

func redirectLoginWithoutAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userid := c.Get("UserID")
			if userid == nil {
				// not login'd, go to login page
				return c.Redirect(http.StatusFound, "/login")
			}
			return next(c)
		}
	}
}

func main() {
	// Echo instance
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} uri=${uri} path=${path} status=${status}\n",
	}))
	e.Use(middleware.Recover())

	var store = sessions.NewCookieStore([]byte("secret key"))
	e.Use(session.Middleware(store))

	// Routes
	e.Static("/css", "./static/css")
	e.GET("/login", routes.LoginRouter)
	e.POST("/login", routes.LoginRouterPost)
	e.GET("/", routes.IndexRouter, setUserMiddleware(), redirectLoginWithoutAuth())

	// handle error
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound, "404", nil)
			} else {
				c.Render(http.StatusInternalServerError, "500", nil)
			}
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
