package usecase

import "finalWork/internal/entity"

type (
	Infrastructure interface {
		GetSMSData([]string) []*entity.SMSData
		GetMMSData(*entity.MMSData) []*entity.MMSData
		GetVoiceCallData([]string) []*entity.VoiceCallData
		GetEmailData([]string) []*entity.EmailData
		GetBillingData([]bool) []*entity.BillingData
		GetSupportData(data *entity.SupportData) []*entity.SupportData
		GetIncidentData(data *entity.IncidentData) []*entity.IncidentData
	}

	Controller interface {
		GetSMSData([]byte) []*entity.SMSData
		GetMMSData([]byte) ([]*entity.MMSData, error)
		GetVoiceCallData([]byte) []*entity.VoiceCallData
		GetEmailData([]byte) []*entity.EmailData
		GetBillingData([]byte) []*entity.BillingData
		GetSupportData([]byte) ([]*entity.SupportData, error)
		GetIncidentData([]byte) ([]*entity.IncidentData, error)
	}
)
