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
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

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

	log.Printf("Server successfully started at port %v\n", server.Addr)
	log.Println(server.ListenAndServe())

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
