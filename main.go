package main

import (
	"akrab-bangkit2022-api/config"
	"akrab-bangkit2022-api/infrastructure/io"
	"akrab-bangkit2022-api/infrastructure/registry"
	"akrab-bangkit2022-api/infrastructure/router"
	"context"
	"fmt"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	config.ReadConfig()

	isDbLogged := false
	if config.C.Server.Env == "stage" {
		gin.SetMode(gin.ReleaseMode)
		isDbLogged = true
	}

	db := io.InitDB(isDbLogged)

	reg := registry.NewRegistry(db)

	r := gin.Default()
	r = router.Routes(r, reg)


	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Port)

	srv := &http.Server{
		Addr:    ":" + config.C.Server.Port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
	}

	fmt.Println("Server exiting")
}
