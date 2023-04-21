package controller

import (
	"finalWork/internal/entity"
	"finalWork/internal/usecase"
	"fmt"
	"net/http"
)

type Controller struct {
	uc usecase.Controller
}

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]entity.SMSData              `json:"sms"`
	MMS       [][]entity.MMSData              `json:"mms"`
	VoiceCall []entity.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]entity.EmailData `json: email"`
	Billing   entity.BillingData              `json: billing"`
	Support   []int                           `json: support"`
	Incidents []entity.IncidentData           `json:"incident"`
}

func New(uc usecase.Controller) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
