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
	smsData := c.prepareSMSData()
	mmsData := c.prepareMMSData()
	voiceCallData := c.prepareVoiceCallData()

	resultSetT := ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceCallData,
		Email:     nil,
		Billing:   BillingData{},
		Support:   nil,
		Incidents: nil,
	}
	return resultSetT
}

func (c *Controller) prepareSMSData() [][]SMSData {
	sortByProvider, sortByCountry := c.uc.GetSMSData()

	dataStoreByProvider := make([]SMSData, 0, len(sortByProvider))
	for _, elem := range sortByProvider {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		dataStoreByProvider = append(dataStoreByProvider, sms)
	}

	dataStoreByCountry := make([]SMSData, 0, len(sortByCountry))
	for _, elem := range sortByCountry {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		dataStoreByCountry = append(dataStoreByCountry, sms)
	}
	dataStoreOrdered := make([][]SMSData, 0, 2)
	dataStoreOrdered = append(dataStoreOrdered, dataStoreByProvider, dataStoreByCountry)

	return dataStoreOrdered

}

func (c *Controller) prepareMMSData() [][]MMSData {
	sortByProvider, sortByCountry := c.uc.GetMMSData()

	dataStoreByProvider := make([]MMSData, 0, len(sortByProvider))
	for _, elem := range sortByProvider {
		mms := MMSData{
			elem.Country,
			elem.Provider,
			elem.Bandwidth,
			elem.ResponseTime,
		}
		dataStoreByProvider = append(dataStoreByProvider, mms)
	}

	dataStoreByCountry := make([]MMSData, 0, len(sortByCountry))
	for _, elem := range sortByCountry {
		mms := MMSData{
			elem.Country,
			elem.Provider,
			elem.Bandwidth,
			elem.ResponseTime,
		}
		dataStoreByCountry = append(dataStoreByCountry, mms)
	}
	dataStoreOrdered := make([][]MMSData, 0, 2)
	dataStoreOrdered = append(dataStoreOrdered, dataStoreByProvider, dataStoreByCountry)

	return dataStoreOrdered

}

func (c *Controller) prepareVoiceCallData() []VoiceCallData {
	data := c.uc.GetVoiceCallData()

	dataStore := make([]VoiceCallData, 0, len(data))
	for _, elem := range data {
		voiceCall := VoiceCallData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
			elem.ConnectionStability,
			elem.TTFB,
			elem.VoicePurity,
			elem.MedianOfCallsTime,
		}
		dataStore = append(dataStore, voiceCall)
	}

	return dataStore

}
