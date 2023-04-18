package usecase

import "finalWork/internal/entity"

type (
	Infrastructure interface {
		GetSMSData([]string) []*entity.SMSData
	}

	Controller interface {
		GetSMSData([]byte) []*entity.SMSData
	}
)
