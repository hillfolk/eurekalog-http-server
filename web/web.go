package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	t := time.Now()
	folderPath := t.Format("2006-01-02")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		fmt.Println("mkdir All")
		os.MkdirAll("./data/"+folderPath, os.ModePerm)
	}
	

	for _, files := range form.File {
		
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("./data/"+folderPath+"/"+file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields.</p>", "file"))
		
	
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func RunServer(port string){
	e := echo.New()
	e.Debug = true
	// Server header
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServerHeader)
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Eurekolog-server")
	})
	
	e.POST("/upload/", upload)
	e.Logger.Fatal(e.Start(port))
	
}
