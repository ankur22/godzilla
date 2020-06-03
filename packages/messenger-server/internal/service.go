package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel()
		sig := <-sigs
		log.Printf("Received sig %v\n", sig)
	}()

	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler)
	r.HandleFunc("/{name}", helloAnyHandler)
	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var err error
	go func() {
		err = s.ListenAndServe()
	}()

	<-ctx.Done()

	return err
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func helloAnyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Fprintf(w, "Hello %s", name)
}
