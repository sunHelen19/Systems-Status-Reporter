package controller

import (
	"encoding/json"
	"finalWork/internal/entity"
	"finalWork/internal/usecase"
	"net/http"
	"sort"
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
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

func New(uc usecase.Controller) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) HandleConnection(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	result := ResultT{}
	data := c.GetResultData()
	status := false

	if data.Email != nil && data.Incidents != nil && data.MMS != nil && data.SMS != nil && data.VoiceCall != nil && data.Support != nil {
		status = true
		result.Data = data
	} else {
		result.Error = "Error on collect data"
	}

	result.Status = status

	resultJson, _ := json.Marshal(result)
	w.Write(resultJson)
}

func (c *Controller) GetResultData() ResultSetT {
	smsData := c.prepareSMSData()
	mmsData := c.prepareMMSData()
	voiceCallData := c.prepareVoiceCallData()
	emailData := c.prepareEmailData()
	billingData := c.prepareBillingData()
	supportData := c.uc.GetSupportData()
	incidentData := c.prepareIncidentData()

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

func (c *Controller) prepareEmailData() map[string][][]EmailData {
	data := c.uc.GetEmailData()

	dataStore := make(map[string][][]EmailData)

	for key, country := range data {
		providers := make([][]EmailData, 0, 2)
		fastProviders := make([]EmailData, 0, 3)

		for _, elem := range country[0] {
			provider := EmailData{
				Country:      elem.Country,
				Provider:     elem.Provider,
				DeliveryTime: elem.DeliveryTime,
			}
			fastProviders = append(fastProviders, provider)

		}

		slowProviders := make([]EmailData, 0, 3)
		for _, elem := range country[1] {
			provider := EmailData{
				Country:      elem.Country,
				Provider:     elem.Provider,
				DeliveryTime: elem.DeliveryTime,
			}
			slowProviders = append(slowProviders, provider)
		}
		sort.Slice(fastProviders, func(i, j int) bool { return fastProviders[i].DeliveryTime > fastProviders[j].DeliveryTime })
		providers = append(providers, fastProviders, slowProviders)
		dataStore[key] = providers
	}

	return dataStore

}

func (c *Controller) prepareBillingData() BillingData {
	data := c.uc.GetBillingData()

	billingData := BillingData{
		CreateCustomer: data.CreateCustomer,
		Purchase:       data.Purchase,
		Payout:         data.Payout,
		Recurring:      data.Recurring,
		FraudControl:   data.FraudControl,
		CheckoutPage:   data.CheckoutPage,
	}

	return billingData

}

func (c *Controller) prepareIncidentData() []IncidentData {
	data := c.uc.GetIncidentData()

	dataStore := make([]IncidentData, 0, len(data))
	for _, elem := range data {
		incident := IncidentData{
			Topic:  elem.Topic,
			Status: elem.Status,
		}
		dataStore = append(dataStore, incident)
	}

	return dataStore

}
