package main

import (
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	deletequote "quoter/internal/http-server/handlers/delete_quote"
	getrandom "quoter/internal/http-server/handlers/get_random"
	"quoter/internal/http-server/handlers/list"
	"quoter/internal/http-server/handlers/save"
	"quoter/internal/storage"

	"github.com/gorilla/mux"

	"os"
)

const addres = "localhost:8080"

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	storage := storage.New()

	r := mux.NewRouter()

	r.HandleFunc("/quotes", list.New(storage)).Methods(http.MethodGet)
	r.HandleFunc("/quotes", save.New(log, storage)).Methods(http.MethodPost)
	r.HandleFunc("/quotes/random", getrandom.New(storage)).Methods(http.MethodGet)
	r.HandleFunc("/quotes/{id}", deletequote.New(storage, log)).Methods(http.MethodDelete)

	go http.ListenAndServe(addres, r)
	log.Info("server started", "addr", addres)

	var sig os.Signal
	sysStop := make(chan os.Signal, 1)
	signal.Notify(sysStop, syscall.SIGTERM, syscall.SIGINT)
	sig = <-sysStop
	log.Info("application fully stoped", slog.String("SIGNAL", sig.String()))
}

