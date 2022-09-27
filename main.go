package main

import (
	"os"
	"strings"

	"github.com/hongsub/learngo/Go-Practice/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv") //첨부파일을 리턴(받아올파일이름,사용자에게 줄 파일이름)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323")) //포트 1323번 연결
}
