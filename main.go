package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/test3", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/test1", func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/test2", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.Render(200, SomeStruct{})
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/json2", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.Render(200, SomeStruct{})
		// c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/json3", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.Render(200, SomeStruct{})
		// c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/json", func(c *gin.Context) {
		// time.Sleep(2 * time.Second)
		c.Render(200, SomeStruct{})
		// c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}

type SomeStruct struct {
	Code int `json:"code"`
}

func (s SomeStruct) SomeHandler(w http.ResponseWriter, r *http.Request) {
	data := SomeStruct{Code: 201}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (s SomeStruct) Render(w http.ResponseWriter) error {
	data := SomeStruct{Code: 201}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func (s SomeStruct) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	data := SomeStruct{Code: 201}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(data)
// }
