package result

import (
	"log"
	"netWorkService/internal/config"
	"netWorkService/internal/data"
)

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]data.SMSData              `json:"sms"`
	MMS       [][]data.MMSData              `json:"mms"`
	VoiceCall []data.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]data.EmailData `json:"email"`
	Billing   data.BillingData              `json:"billing"`
	Support   []int                         `json:"support"`
	Incidents []data.IncidentData           `json:"incident"`
}

func RunGetData(cfg *config.DataConfig) ResultSetT {

	providers, err := data.GetProviders(cfg.ProvidersFile.Path)
	if err != nil {

		panic(err)
	}

	smsData := data.PrepareSMSData(cfg.FileName.Sms, providers.SMS)
	mmsData := data.PrepareMMSData("http://"+cfg.Server.Host+":"+cfg.Server.Port+"/"+cfg.Server.Mms, providers.MMS)
	voiceCallData := data.PrepareVoiceCallData(cfg.FileName.Voice, providers.VoiceCall)
	emailData := data.PrepareEmailData(cfg.FileName.Email, providers.Email)
	billingData := data.PrepareBillingData(cfg.FileName.Billing)
	supportData := data.PrepareSupportData("http://" + cfg.Server.Host + ":" + cfg.Server.Port + "/" + cfg.Server.Support)
	incidentData := data.PrepareIncidentData("http://" + cfg.Server.Host + ":" + cfg.Server.Port + "/" + cfg.Server.Accendent)

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

func GetResultData() ResultSetT {

	cfg, err := config.NewConfigData("./configs/data.yml")
	if err != nil {
		log.Fatal(err)

	}

	data := RunGetData(cfg)
	return data

}
