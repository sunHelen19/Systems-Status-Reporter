package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func PrepareSupportData() []int {
	result := make([]int, 0, 2)
	data := getSupportData()
	sumTickets := 0
	workload := 1

	for _, topic := range data {
		sumTickets += topic.ActiveTickets
	}

	if sumTickets >= 9 && sumTickets <= 16 {
		workload = 2
	} else if sumTickets > 16 {
		workload = 3
	}
	result = append(result, workload)

	var waitTime int
	oneTicketTime := 60 / 18
	waitTime = oneTicketTime * sumTickets
	result = append(result, waitTime)

	return result

}

func getSupportData() []*SupportData {
	dataStore := make([]*SupportData, 0)
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

		var str []*SupportData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return nil
		}

		dataStore = append(dataStore, str...)

		return dataStore
	}

	return nil
}
