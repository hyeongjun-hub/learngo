package main

import (
	"github.com/hyeongjun-hub/learngo/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("home.html")
	})
	e.POST("/scrape", func(c echo.Context) error {
		term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
		defer os.RemoveAll("jobs.csv")
		scrapper.Scrape(term)
		return c.Attachment("jobs.csv", "jobs.csv")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
