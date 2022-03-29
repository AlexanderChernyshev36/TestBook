package main

import (
	"TestBook/api/author"
	"TestBook/api/book"
	"TestBook/api/genre"
	"TestBook/models/modelsjson"
	"TestBook/sql"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Db_Connect *gorm.DB

func main() {

	db, err := sql.Connectdb()
	Db_Connect = db
	if err != nil {
		fmt.Println(err)
		fmt.Println("ошибка подключения к БД")
		panic(err)
	}
	handler := gin.New()
	handler.Use(gin.Logger(), gin.Recovery(), Logger())
	handler.POST("/setAuthor/", HandlerSetAuthor)
	handler.POST("/setGenre/", HandlerSetGenre)
	handler.POST("/setBook/", HandlerSetBook)
	handler.POST("/saveBook/", HandlerSaveBook)
	handler.DELETE("/deleteAuthor/", HandlerDeleteAuthor)
	handler.DELETE("/deleteGenre/", HandlerDeleteGenre)
	handler.DELETE("/deleteBook/", HandlerDeleteBook)
	handler.GET("/viewAuthor/", HandlerViewAuthor)
	handler.GET("/viewGenre/", HandlerViewGenre)
	handler.GET("/viewBook/", HandlerViewBook)

	///////////////////НАСТРОЙКА СЕРВЕРА//////////////////////////////////////
	s := &http.Server{
		Addr:           ":5999",
		Handler:        handler,           // if nil use default http.DefaultServeMux
		ReadTimeout:    100 * time.Second, // max duration reading entire request
		WriteTimeout:   100 * time.Second, // max timing write response
		IdleTimeout:    150 * time.Second, // max time wait for the next request
		MaxHeaderBytes: 1 << 20,           // 2^20 or 128kbytes
	}

	go func() {
		textlog := "Listening on https://" + s.Addr
		log.Printf(textlog)
		log.Fatal(s.ListenAndServe()) //для запуска http сервера
	}()

	graceful(s, 5*time.Second)

}

func HandlerSetAuthor(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)                        //читаем тело
	request_SetAuthor := []modelsjson.AuthorsBook{}                       //присваиваем переменной структуру json "SetAuthor"
	if err := json.Unmarshal(bodyBytes, &request_SetAuthor); err == nil { //запись из тела body в структуру
		result := author.SetAuthor(Db_Connect, request_SetAuthor)
		c.JSON(http.StatusOK, result)
	}
}

func HandlerSetGenre(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	request_SetGenre := []modelsjson.Zhanrs{}
	if err := json.Unmarshal(bodyBytes, &request_SetGenre); err == nil {
		result := genre.SetGenre(Db_Connect, request_SetGenre)
		c.JSON(http.StatusOK, result)
	}
}

func HandlerSetBook(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	request_SetBook := []modelsjson.Books{}
	if err := json.Unmarshal(bodyBytes, &request_SetBook); err == nil {
		result := book.SetBook(Db_Connect, request_SetBook)
		c.JSON(http.StatusOK, result)
	}
}

func HandlerSaveBook(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	request_SaveBook := modelsjson.QuerySaveBook{}
	if err := json.Unmarshal(bodyBytes, &request_SaveBook); err == nil {
		result := genre.ViewQuerySaveBook(Db_Connect, request_SaveBook)
		c.JSON(http.StatusOK, result)
	}
}

func HandlerDeleteAuthor(c *gin.Context) {
	id := c.Query("id")
	result := author.DeleteAuthor(Db_Connect, id)
	c.JSON(http.StatusOK, result)
}

func HandlerDeleteGenre(c *gin.Context) {
	id := c.Query("id")
	result := genre.DeleteGenre(Db_Connect, id)
	c.JSON(http.StatusOK, result)
}

func HandlerDeleteBook(c *gin.Context) {
	id := c.Query("id")
	result := book.DeleteBook(Db_Connect, id)
	c.JSON(http.StatusOK, result)
}

func HandlerViewGenre(c *gin.Context) {
	name := c.Query("name")
	result := genre.ViewGenre(Db_Connect, name)
	c.JSON(http.StatusOK, result)
}

func HandlerViewAuthor(c *gin.Context) {
	id := c.Query("id")
	result := author.ViewAuthor(Db_Connect, id)
	c.JSON(http.StatusOK, result)
}

func HandlerViewBook(c *gin.Context) {
	id := c.Query("id")
	result := book.ViewBook(Db_Connect, id)
	c.JSON(http.StatusOK, result)
}

//логирование подключений к серверу http
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		textlog := "method " + c.Request.Method + "  connection from " + c.Request.RemoteAddr + " url " + c.Request.URL.Path + c.Request.URL.RawQuery
		log.Printf(textlog)
		c.Next()
	}
}

//корректное отключение http сервера ctr+c
func graceful(hs *http.Server, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Println("Server stopped")
	}

}
