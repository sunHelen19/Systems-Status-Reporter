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

type Store struct{}

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
	return &Store{}
}

func (s *Store) GetSMSData() []*entity.SMSData {
	dataStore := make([]*entity.SMSData, 0)
	data, err := readFile("src/simulator/data/sms.data")
	if err != nil {
		return nil
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
		dataStore = append(dataStore, &str)

	}

	return dataStore
}

func (s *Store) GetMMSData() []*entity.MMSData {
	dataStore := make([]*entity.MMSData, 0)
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {
			return nil
		}
		providers := []string{"Topolo", "Rond", "Kildy"}
		var str []*entity.MMSData
		if errJson := json.Unmarshal(body, &str); errJson != nil {
			return nil
		}

		for _, elem := range str {
			if elem.HasCountryAlpha2Code() {
				if elem.CheckCorrectProviders(providers) {
					dataStore = append(dataStore, elem)

				}
			}
		}

		return dataStore

	}
	return nil
}

func (s *Store) GetVoiceCallData() []*entity.VoiceCallData {
	dataStore := make([]*entity.VoiceCallData, 0)
	data, err := readFile("src/simulator/data/voice.data")
	if err != nil {
		return nil
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
		dataStore = append(dataStore, &str)

	}

	return dataStore
}

func (s *Store) GetEmailData() []*entity.EmailData {

	data, err := readFile("src/simulator/data/email.data")
	if err != nil {
		return nil
	}

	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
	dataSlice := getDataStringSlice(data, "\n", 3, providers, 1)
	dataStore := make([]*entity.EmailData, 0)
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
		dataStore = append(dataStore, &str)

	}

	return dataStore
}

func (s *Store) GetBillingData() *entity.BillingData {

	data, err := readFile("src/simulator/data/billing.data")
	if err != nil {
		return nil
	}
	if len(data) != 0 {
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
	return nil
}

func (s *Store) GetSupportData() []*entity.SupportData {
	dataStore := make([]*entity.SupportData, 0)
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {

		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return nil
		}

		var str []*entity.SupportData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return nil
		}

		dataStore = append(dataStore, str...)

		return dataStore
	}

	return nil
}

func (s *Store) GetIncidentData() []*entity.IncidentData {
	dataStore := make([]*entity.IncidentData, 0)
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {

		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return nil
		}
		var str []*entity.IncidentData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return nil
		}

		dataStore = append(dataStore, str...)

		return dataStore
	}

	return nil

}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {

		return nil, fmt.Errorf("Ошибка открытия файла %v", err)
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
