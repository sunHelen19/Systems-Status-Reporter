package usecase

import "finalWork/internal/entity"

type (
	Infrastructure interface {
		GetSMSData() []*entity.SMSData
		GetMMSData() []*entity.MMSData
		GetVoiceCallData() []*entity.VoiceCallData
		GetEmailData() []*entity.EmailData
		GetBillingData() []*entity.BillingData
		GetSupportData() []*entity.SupportData
		GetIncidentData() []*entity.IncidentData
	}

	Controller interface {
		GetSMSData() ([]*entity.SMSData, []*entity.SMSData)
		GetMMSData() ([]*entity.MMSData, []*entity.MMSData)
		//	GetVoiceCallData([]byte) []*entity.VoiceCallData
		//GetEmailData([]byte) []*entity.EmailData
		//	GetBillingData([]byte) []*entity.BillingData
		//GetSupportData([]byte) ([]*entity.SupportData, error)
		//GetIncidentData([]byte) ([]*entity.IncidentData, error)
	}
)
