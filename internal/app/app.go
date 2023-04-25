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

	server := &http.Server{Addr: "127.0.0.1:8282", Handler: service(c)}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

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
	r.HandleFunc("/json", c.HandleConnection)
	r.HandleFunc("/api", c.HandleConnection).Methods("GET", "OPTIONS")

	staticFileDirectory := http.Dir("./web/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	return r
}

/*
func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./web/index.html"
	}
	http.ServeFile(w, r, p)
}
*/
