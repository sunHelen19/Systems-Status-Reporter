package data

import (
	"fmt"
	"io"
	"math"
	"netWorkService/src"
	"os"
	"strconv"
	"strings"
)

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

func getCountryName(code string) string {
	countryName := src.Countries[code]
	return countryName
}
