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

type (
	SMSData       entity.SMSData
	MMSData       entity.MMSData
	VoiceCallData entity.VoiceCallData
	EmailData     entity.EmailData
	BillingData   entity.BillingData
	SupportData   entity.SupportData
	IncidentData  entity.IncidentData
)

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json: email"`
	Billing   BillingData              `json: billing"`
	Support   []int                    `json: support"`
	Incidents []IncidentData           `json:"incident"`
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

func (c *Controller) GetResultData() ResultSetT {
	smsDataStoreOrdered := c.prepareSMSData()

	resultSetT := ResultSetT{
		SMS:       smsDataStoreOrdered,
		MMS:       nil,
		VoiceCall: nil,
		Email:     nil,
		Billing:   BillingData{},
		Support:   nil,
		Incidents: nil,
	}
	return resultSetT
}

func (c *Controller) prepareSMSData() [][]SMSData {
	sortByProvider, sortByCountry := c.uc.GetSMSData()

	smsDataStoreByProvider := make([]SMSData, len(sortByProvider))
	for _, elem := range sortByProvider {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		smsDataStoreByProvider = append(smsDataStoreByProvider, sms)
	}

	smsDataStoreByCountry := make([]SMSData, len(sortByCountry))
	for _, elem := range sortByCountry {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		smsDataStoreByCountry = append(smsDataStoreByCountry, sms)
	}
	smsDataStoreOrdered := make([][]SMSData, 0, 2)
	smsDataStoreOrdered = append(smsDataStoreOrdered, smsDataStoreByProvider, smsDataStoreByCountry)

	return smsDataStoreOrdered

}
