package infrastructure

import (
	"finalWork/internal/entity"
	"strconv"
	"strings"
)

type Store struct {
	SMSDataStore       []*entity.SMSData
	MMSDataStore       []*entity.MMSData
	VoiceCallDataStore []*entity.VoiceCallData
	EmailDataStore     []*entity.EmailData
	BillingDataStore   []*entity.BillingData
	SupportDataStore   []*entity.SupportData
}

func CreateStore() *Store {
	return &Store{
		make([]*entity.SMSData, 0, 0),
		make([]*entity.MMSData, 0, 0),
		make([]*entity.VoiceCallData, 0, 0),
		make([]*entity.EmailData, 0, 0),
		make([]*entity.BillingData, 0, 0),
		make([]*entity.SupportData, 0, 0),
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

func (s *Store) GetMMSData(str *entity.MMSData) []*entity.MMSData {
	s.MMSDataStore = append(s.MMSDataStore, str)
	return s.MMSDataStore
}

func (s *Store) GetVoiceCallData(data []string) []*entity.VoiceCallData {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")

		connectionStability, errCS := strconv.ParseFloat(elemSlice[4], 64)
		if errCS != nil {
			continue
		}
		TTFB, errTTFB := strconv.Atoi(elemSlice[5])
		if errTTFB != nil {
			continue
		}
		voicePurity, errVP := strconv.Atoi(elemSlice[6])
		if errVP != nil {
			continue
		}

		mediaOfCallsTime, errMOCT := strconv.Atoi(elemSlice[7])
		if errMOCT != nil {
			continue
		}

		str := entity.VoiceCallData{
			Country:             elemSlice[0],
			Bandwidth:           elemSlice[1],
			ResponseTime:        elemSlice[2],
			Provider:            elemSlice[3],
			ConnectionStability: float32(connectionStability),
			TTFB:                TTFB,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   mediaOfCallsTime,
		}
		s.VoiceCallDataStore = append(s.VoiceCallDataStore, &str)

	}

	return s.VoiceCallDataStore
}

func (s *Store) GetEmailData(data []string) []*entity.EmailData {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")

		deliveryTime, errDT := strconv.Atoi(elemSlice[2])
		if errDT != nil {
			continue
		}

		str := entity.EmailData{
			Country:      elemSlice[0],
			Provider:     elemSlice[1],
			DeliveryTime: deliveryTime,
		}
		s.EmailDataStore = append(s.EmailDataStore, &str)

	}

	return s.EmailDataStore
}

func (s *Store) GetBillingData(data []bool) []*entity.BillingData {

	str := entity.BillingData{
		CreateCustomer: data[0],
		Purchase:       data[1],
		Payout:         data[2],
		Recurring:      data[3],
		FraudControl:   data[4],
		CheckoutPage:   data[5],
	}
	s.BillingDataStore = append(s.BillingDataStore, &str)

	return s.BillingDataStore
}

func (s *Store) GetSupportData(str *entity.SupportData) []*entity.SupportData {
	s.SupportDataStore = append(s.SupportDataStore, str)
	return s.SupportDataStore
}
