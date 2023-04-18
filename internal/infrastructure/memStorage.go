package infrastructure

import (
	"finalWork/internal/entity"
	"strings"
)

type Store struct {
	SMSDataStore []*entity.SMSData
}

func CreateStore() *Store {
	return &Store{
		make([]*entity.SMSData, 0, 0),
	}
}

func (s *Store) GetSMSData(data []string) []*entity.SMSData {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")
		str := entity.SMSData{
			Country:      elemSlice[0],
			Bandwidth:    elemSlice[1],
			ResponseTime: elemSlice[2],
			Provider:     elemSlice[3],
		}
		s.SMSDataStore = append(s.SMSDataStore, &str)

	}

	return s.SMSDataStore
}
