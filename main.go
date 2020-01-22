package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/meromen/db-inteface-v2/controller"
	dbpkg "github.com/meromen/db-inteface-v2/db"
)

func main() {
	db, err := dbpkg.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = dbpkg.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	c := controller.New(db)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      c,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Printf("Server successfully started at port %v\n", server.Addr)
	log.Println(server.ListenAndServe())
}
