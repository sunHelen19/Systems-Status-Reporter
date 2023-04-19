package usecase

import "finalWork/internal/entity"

type (
	Infrastructure interface {
		GetSMSData([]string) []*entity.SMSData
		GetMMSData(*entity.MMSData) []*entity.MMSData
		GetVoiceCallData([]string) []*entity.VoiceCallData
		GetEmailData([]string) []*entity.EmailData
	}

	Controller interface {
		GetSMSData([]byte) []*entity.SMSData
		GetMMSData([]byte) ([]*entity.MMSData, error)
		GetVoiceCallData([]byte) []*entity.VoiceCallData
		GetEmailData([]byte) []*entity.EmailData
	}
)
