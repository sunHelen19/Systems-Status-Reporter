package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"netWorkService/internal"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]internal.SMSData              `json:"sms"`
	MMS       [][]internal.MMSData              `json:"mms"`
	VoiceCall []internal.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]internal.EmailData `json:"email"`
	Billing   internal.BillingData              `json:"billing"`
	Support   []int                             `json:"support"`
	Incidents []internal.IncidentData           `json:"incident"`
}

func main() {

	server := &http.Server{Addr: "127.0.0.1:8282", Handler: service()}
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

func service() http.Handler {
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
	result := ResultT{}
	data := GetResultData()
	status := false
	billingDataSlice := make([]*internal.BillingData, 0, 1)

	if data.Email != nil && data.Incidents != nil && data.MMS != nil && data.SMS != nil && data.VoiceCall != nil && data.Support != nil && billingDataSlice != nil {
		status = true
		result.Data = data
	} else {

		result.Error = "Error on collect data"
	}

	result.Status = status

	resultJson, _ := json.MarshalIndent(result, "", " ")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}

}

func GetResultData() ResultSetT {
	smsData := internal.PrepareSMSData()
	mmsData := internal.PrepareMMSData()
	voiceCallData := internal.PrepareVoiceCallData()
	emailData := internal.PrepareEmailData()
	billingData := internal.PrepareBillingData()
	supportData := internal.PrepareSupportData()
	incidentData := internal.PrepareIncidentData()

	resultSetT := ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceCallData,
		Email:     emailData,
		Billing:   billingData,
		Support:   supportData,
		Incidents: incidentData,
	}
	return resultSetT
}
