package app

import (
	"context"
	"finalWork/internal/controller"
	"finalWork/internal/infrastructure"
	"finalWork/internal/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	repository := infrastructure.CreateStore()
	useCase := usecase.New(repository)
	c := controller.New(useCase)

	server := &http.Server{Addr: "localhost:8282", Handler: service(c)}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()

}

func service(c *controller.Controller) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", c.HandleConnection)
	return r
}
