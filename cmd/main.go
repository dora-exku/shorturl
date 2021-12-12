package main

import (
	"context"
	"fmt"
	"github.com/dora-exku/shorturl/config"
	"github.com/dora-exku/shorturl/controller"
	"github.com/dora-exku/shorturl/model"
	c "github.com/dora-exku/shorturl/pkg/config"
	"github.com/dora-exku/shorturl/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Initialize()

	database.ContentDB()

	database.DB.AutoMigrate(&model.Url{})

	route := gin.Default()
	route.Any("/url", controller.Store)
	route.GET("/:key", controller.Jump)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.GetString("app.host"), c.GetString("app.port")),
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
