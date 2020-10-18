package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/kingofmidas/gfc-api/internal/api"
	"github.com/kingofmidas/gfc-api/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	storeConfig := store.NewConfig()
	_, err := toml.DecodeFile(configPath, storeConfig)
	if err != nil {
		log.Fatal(err)
	}

	s := store.New(storeConfig)
	if err = s.OpenConnection(); err != nil {
		log.Fatal(err)
	}
	defer s.CloseConnection()

	mux := mux.NewRouter()
	handler := api.Handler{Store: s}

	mux.HandleFunc("/orders/create", handler.CreateOrder).Methods("POST")
	mux.HandleFunc("/orders/mark-ready/{id}", handler.UpdateOrderReady).Methods("PATCH")
	mux.HandleFunc("/orders/mark-complete/{id}", handler.UpdateOrderComplete).Methods("PATCH")
	mux.HandleFunc("/orders/list-await", handler.GetOrdersAwait).Methods("GET")
	mux.HandleFunc("/orders/list-ready", handler.GetOrdersReady).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+v", err)
		}
	}()

	log.Println("Starting...")

	<-done

	log.Println("Server stopped ...")

	ctx := context.TODO()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s", err)
	}

	log.Println("Exited")
}
