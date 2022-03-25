package main

import (
	"TestBook/models/modelsjson"
	"TestBook/models/modelssql"
	"TestBook/sql"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

//func SetBook(){
//	Db_Connect, err := sql.Connectdb()
//
//			if err != nil {
//				panic(err)
//			}
//			Db_Connect.AutoMigrate(&modelssql.Books{})
//}
//
//func HandlerSetBook(c *gin.Context){
//	SetBook()
//}

//func SetGenre(){
//	Db_Connect, err := sql.Connectdb()
//
//		if err != nil {
//			panic(err)
//		}
//		Db_Connect.AutoMigrate(&modelssql.Zhanrs{})
//}
//
//func HandlerSetGenre(c *gin.Context){
//	SetGenre()
//}

func SetAuthor(request_SetAuthor []modelsjson.AuthorsBook) string {
	result := ""
	Db_Connect, err := sql.Connectdb()

	if err != nil {
		fmt.Println(err)
		result = "Ошибка подключения к БД."
		return result
	}

	for _, jsonAuthor := range request_SetAuthor {
		modelAuthorsBook := modelssql.AuthorsBook{}
		Db_Connect.Model(&modelAuthorsBook).Where("id = ?", jsonAuthor.Id).Find(&modelAuthorsBook)

		if modelAuthorsBook.Id == 0 {
			//create
			modelAuthorsBook.Name = jsonAuthor.Name
			modelAuthorsBook.Age = jsonAuthor.Age
			modelAuthorsBook.Website = jsonAuthor.Website
			modelAuthorsBook.Email = jsonAuthor.Email
			Db_Connect.Model(modelAuthorsBook).Create(&modelAuthorsBook)
			fmt.Println("create" + strconv.Itoa(jsonAuthor.Id))
		} else {
			//update
			Db_Connect.Model(&modelAuthorsBook).Where("id = ?", jsonAuthor.Id).
				Updates(map[string]interface{}{
					"name":    jsonAuthor.Name,
					"age":     jsonAuthor.Age,
					"website": jsonAuthor.Website,
					"email":   jsonAuthor.Email})
			fmt.Println("update " + strconv.Itoa(jsonAuthor.Id))
		}
	}
	return result
}

func HandlerSetAuthor(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)                        //читаем тело
	request_SetAuthor := []modelsjson.AuthorsBook{}                       //присваиваем переменной структуру json "SetAuthor"
	if err := json.Unmarshal(bodyBytes, &request_SetAuthor); err == nil { //запись из тела body в структуру
		result := SetAuthor(request_SetAuthor)
		c.JSON(http.StatusOK, result)
	}
}

func main() {

	handler := gin.New()
	handler.Use(gin.Logger(), gin.Recovery(), Logger())
	handler.POST("/setAuthor/", HandlerSetAuthor)
	//handler.GET("/setGenre/",HandlerSetGenre)
	//handler.GET("/setBook/",HandlerSetBook)
	//handler.POST("/setAuthor/",HandlerSetAuthor)
	//handler.POST("/setGenre/",HandlerSetGenre)
	//handler.POST("/setBook/",HandlerSetBook)
	//handler.DELETE("/deleteAuthor/",HandlerDeleteAuthor)
	//handler.DELETE("/deleteGenre /",HandlerDeleteGenre)
	//handler.DELETE("/deleteBook /",HandlerDeleteBook)
	//handler.DELETE("/deleteAuthor/",HandlerDeleteAuthor)
	//handler.DELETE("/deleteGenre/",HandlerDeleteGenre)
	//handler.DELETE("/deleteBook/",HandlerDeleteBook)
	//handler.GET("/viewAuthor/",HandlerViewAuthor)
	//handler.GET("/viewGenre/",HandlerViewGenre)
	//handler.GET("/viewBook/",HandlerViewBook)
	//handler.GET("/viewAuthor/",HandlerViewAuthor)
	//handler.GET("/viewGenre/",HandlerViewGenre)
	//handler.GET("/viewBook/",HandlerViewBook)

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
