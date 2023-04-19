package controller

import (
	"finalWork/internal/entity"
	"finalWork/internal/usecase"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Controller struct {
	uc usecase.Controller
}

func New(uc usecase.Controller) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) GetSMSData() []*entity.SMSData {
	data, err := readFile("src/simulator/sms.data")
	if err != nil {
		return nil
	}
	dataSlice := c.uc.GetSMSData(data)
	return dataSlice
}

func (c *Controller) GetMMSData() (result []*entity.MMSData) {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {
			return
		}

		rep, errGetData := c.uc.GetMMSData(body)
		if errGetData != nil {
			return
		}
		result = rep
		return
	}

	return
}

func (c *Controller) GetVoiceCallData() []*entity.VoiceCallData {
	data, err := readFile("src/simulator/voice.data")
	if err != nil {
		return nil
	}
	dataSlice := c.uc.GetVoiceCallData(data)
	return dataSlice
}

func (c *Controller) GetEmailData() []*entity.EmailData {
	data, err := readFile("src/simulator/email.data")
	if err != nil {
		return nil
	}
	dataSlice := c.uc.GetEmailData(data)
	return dataSlice
}

func (c *Controller) GetBillingData() []*entity.BillingData {
	data, err := readFile("src/simulator/billing.data")
	if err != nil {
		return nil
	}

	dataSlice := c.uc.GetBillingData(data)
	return dataSlice
}

func (c *Controller) GetSupportData() (result []*entity.SupportData) {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {

		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return
		}

		rep, errGetData := c.uc.GetSupportData(body)
		if errGetData != nil {

			return
		}
		result = rep
		return
	}

	return
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
