package usecase

import "finalWork/internal/entity"

type (
	Infrastructure interface {
		GetSMSData() []*entity.SMSData
		GetMMSData() []*entity.MMSData
		GetVoiceCallData() []*entity.VoiceCallData
		GetEmailData() []*entity.EmailData
		GetBillingData() *entity.BillingData
		GetSupportData() []*entity.SupportData
		GetIncidentData() []*entity.IncidentData
	}

	Controller interface {
		GetSMSData() ([]*entity.SMSData, []*entity.SMSData)
		GetMMSData() ([]*entity.MMSData, []*entity.MMSData)
		GetVoiceCallData() []*entity.VoiceCallData
		GetEmailData() map[string][][]*entity.EmailData
		GetBillingData() *entity.BillingData
		GetSupportData() []int
		GetIncidentData() []*entity.IncidentData
	}
)
