package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func PrepareIncidentData() []IncidentData {
	data := getIncidentData()
	if len(data) == 0 {
		return nil
	}
	dataStore := make([]IncidentData, 0, len(data))
	for _, elem := range data {
		incident := IncidentData{
			Topic:  elem.Topic,
			Status: elem.Status,
		}
		dataStore = append(dataStore, incident)
	}

	return dataStore

}

func getIncidentData() []*IncidentData {
	dataStore := make([]*IncidentData, 0)
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
		var str []*IncidentData
		if errJson := json.Unmarshal(body, &str); errJson != nil {

			return nil
		}

		dataStore = append(dataStore, str...)
		sort.Slice(dataStore, func(i, j int) bool {
			return dataStore[i].Status == "active"
		})
		return dataStore
	}

	return nil

}
