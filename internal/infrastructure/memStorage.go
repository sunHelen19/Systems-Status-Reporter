package infrastructure

import (
	"encoding/json"
	"finalWork/internal/entity"
	"finalWork/src"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Store struct {
	SMSDataStore       []*entity.SMSData
	MMSDataStore       []*entity.MMSData
	VoiceCallDataStore []*entity.VoiceCallData
	EmailDataStore     []*entity.EmailData
	SupportDataStore   []*entity.SupportData
	IncidentDataStore  []*entity.IncidentData
}

type keySet uint8

const (
	CreateCustomer keySet = 1 << iota
	Purchase
	Payout
	Recurring
	FraudControl
	CheckoutPage
	maxKey
)

func CreateStore() *Store {
	return &Store{
		make([]*entity.SMSData, 0, 0),
		make([]*entity.MMSData, 0, 0),
		make([]*entity.VoiceCallData, 0, 0),
		make([]*entity.EmailData, 0, 0),
		make([]*entity.SupportData, 0, 0),
		make([]*entity.IncidentData, 0, 0),
	}
}

func (s *Store) GetSMSData() []*entity.SMSData {
	s.SMSDataStore = nil
	data, err := readFile("src/simulator/data/sms.data")
	if err != nil {
		return s.SMSDataStore
	}
	providers := []string{"Topolo", "Rond", "Kildy"}
	dataSlice := getDataStringSlice(data, "\n", 4, providers, 3)

	for _, elem := range dataSlice {
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

func (s *Store) GetMMSData() []*entity.MMSData {
	s.MMSDataStore = nil
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return s.MMSDataStore
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {
			return s.MMSDataStore
		}
		providers := []string{"Topolo", "Rond", "Kildy"}
		var str []*entity.MMSData
		if errJson := json.Unmarshal(body, &str); errJson != nil {
			return s.MMSDataStore
		}

		for _, elem := range str {
			if elem.HasCountryAlpha2Code() {
				if elem.CheckCorrectProviders(providers) {
					s.MMSDataStore = append(s.MMSDataStore, elem)

				}
			}
		}

		return s.MMSDataStore

	}
	return s.MMSDataStore
}

func (s *Store) GetVoiceCallData() []*entity.VoiceCallData {
	s.VoiceCallDataStore = nil
	data, err := readFile("src/simulator/data/voice.data")
	if err != nil {
		return s.VoiceCallDataStore
	}

	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	dataSlice := getDataStringSlice(data, "\n", 8, providers, 3)

	for _, elem := range dataSlice {
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

func (s *Store) GetEmailData() []*entity.EmailData {
	s.EmailDataStore = nil
	data, err := readFile("src/simulator/data/email.data")
	if err != nil {
		return s.EmailDataStore
	}

	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
	dataSlice := getDataStringSlice(data, "\n", 3, providers, 1)

	for _, elem := range dataSlice {
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

func (s *Store) GetBillingData() *entity.BillingData {

	data, err := readFile("src/simulator/data/billing.data")
	if err != nil {
		return nil
	}

	dataSum := getSumBits(data)
	dataSlice := dataSum.String()

	str := entity.BillingData{
		CreateCustomer: dataSlice[0],
		Purchase:       dataSlice[1],
		Payout:         dataSlice[2],
		Recurring:      dataSlice[3],
		FraudControl:   dataSlice[4],
		CheckoutPage:   dataSlice[5],
	}

	return &str
}

func (s *Store) GetSupportData() []*entity.SupportData {
	s.SupportDataStore = nil
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {

		return s.SupportDataStore
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return s.SupportDataStore
		}

		var str []*entity.SupportData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return s.SupportDataStore
		}
		for _, elem := range str {

			s.SupportDataStore = append(s.SupportDataStore, elem)

		}

	}

	return s.SupportDataStore
}

func (s *Store) GetIncidentData() []*entity.IncidentData {
	s.IncidentDataStore = nil
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {

		return s.IncidentDataStore
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return s.IncidentDataStore
		}
		var str []*entity.IncidentData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return s.IncidentDataStore
		}
		for _, elem := range str {

			s.IncidentDataStore = append(s.IncidentDataStore, elem)

		}

	}

	return s.IncidentDataStore

}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Ошибка открытия файла", err)
		return nil, err
	}
	defer file.Close()
	resultBytes, errRB := io.ReadAll(file)
	if errRB != nil {
		panic(err)
	}
	return resultBytes, nil
}

func getDataStringSlice(data []byte, sep string, fieldsAmount uint, providers []string, indexForProvider int) (dataSlice []string) {
	dataString := string(data)
	dataSlice = strings.Split(dataString, sep)

	dataSlice = checkDataStringFields(dataSlice, fieldsAmount, providers, indexForProvider)
	return
}

func checkDataStringFields(data []string, fieldsAmount uint, providers []string, indexForProvider int) (correctData []string) {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")
		if len(elemSlice) == int(fieldsAmount) {
			if hasStringCountryAlpha2(elemSlice[0]) {
				for _, rightProvider := range providers {
					if rightProvider == elemSlice[indexForProvider] {
						correctData = append(correctData, elem)
					}
				}

			}

		}

	}
	return
}

func hasStringCountryAlpha2(code string) (result bool) {
	country := src.Countries[code]
	if country != "" {
		result = true
	}
	return
}

func (k keySet) String() (data []bool) {
	if k >= maxKey {
		panic("Broken keyset")
	}

	for key := CreateCustomer; key < maxKey; key <<= 1 {
		if k&key != 0 {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}
	return
}

func getSumBits(data []byte) (sum keySet) {
	length := len(data)
	dataString := string(data)

	for index, elem := range dataString {
		elemStr := string(elem)
		elemInt, errElemInt := strconv.Atoi(elemStr)
		if errElemInt != nil {
			panic(errElemInt)
		}

		if elemInt == 1 {
			sum += keySet(math.Pow(2, float64(length-1-index)))

		}
	}
	return
}
