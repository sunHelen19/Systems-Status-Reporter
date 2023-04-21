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
		//GetSMSData([]byte) []*entity.SMSData
		//	GetMMSData([]byte) ([]*entity.MMSData, error)
		//	GetVoiceCallData([]byte) []*entity.VoiceCallData
		//GetEmailData([]byte) []*entity.EmailData
		//	GetBillingData([]byte) []*entity.BillingData
		//GetSupportData([]byte) ([]*entity.SupportData, error)
		//GetIncidentData([]byte) ([]*entity.IncidentData, error)
	}
)
