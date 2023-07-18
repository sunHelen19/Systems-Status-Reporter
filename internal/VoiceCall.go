package internal

import (
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func PrepareVoiceCallData() []VoiceCallData {
	data := getVoiceCallData()
	if len(data) == 0 {
		return nil
	}
	dataStore := make([]VoiceCallData, 0, len(data))
	for _, elem := range data {
		voiceCall := VoiceCallData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
			elem.ConnectionStability,
			elem.TTFB,
			elem.VoicePurity,
			elem.MedianOfCallsTime,
		}
		dataStore = append(dataStore, voiceCall)
	}

	return dataStore

}

func getVoiceCallData() []*VoiceCallData {
	dataStore := make([]*VoiceCallData, 0)
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

		str := VoiceCallData{
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
