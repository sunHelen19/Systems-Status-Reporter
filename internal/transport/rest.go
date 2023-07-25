package transport

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"netWorkService/internal/config"
	"netWorkService/internal/data"
	"netWorkService/internal/result"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Service() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/json", HandleConnection)
	r.HandleFunc("/api", HandleConnection).Methods("GET", "OPTIONS")

	staticFileDirectory := http.Dir("./web/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	return r

}
func HandleConnection(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	resultStruct := result.ResultT{}
	resultData := result.GetResultData()
	//fmt.Println(resultData.SMS)
	status := false
	billingDataSlice := make([]*data.BillingData, 0, 1)

	if resultData.Email != nil && resultData.Incidents != nil && resultData.MMS != nil && resultData.SMS != nil && resultData.VoiceCall != nil && resultData.Support != nil && billingDataSlice != nil {
		status = true
		resultStruct.Data = resultData
	} else {

		resultStruct.Error = "Error on collect data"
	}

	resultStruct.Status = status

	resultJson, _ := json.MarshalIndent(resultStruct, "", " ")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}

}

func Run(cfg *config.Config) {

	server := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: Service(),
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	var sig = make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, time.Duration(cfg.Server.Timeout))
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
