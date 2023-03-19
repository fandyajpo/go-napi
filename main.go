package main

import (
	"context"
	// "fmt"
	"go-napi/auth"
	"go-napi/connection"
	"go-napi/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	// PRODUCTION
// 	// if err := godotenv.Load("/home/rahman/cassanovaGoAPI/.env"); err != nil {
// 	if err := godotenv.Load(); err != nil {
// 		panic("Server Shutdown:")
// 	}
// }

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	ai := auth.NewAuth()
	di := connection.NewDB()

	srvs := service.NewHandler(ai, di)

	r.POST("/register", srvs.Register)
	r.GET("/books", srvs.Book)
	r.GET("/string", srvs.StringReturn)

	s := &http.Server{
		Addr:           ":9035",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
