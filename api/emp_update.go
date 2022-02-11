package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// func FxPush(path, token string, body interface{}) (int, error) {
// 	payload := &MoPayload{
// 		PeriodStart: time.Now().Format("2006-01-02"),
// 		PeriodEnd:   time.Now().Format("2006-01-02"),
// 		PeriodKey:   "month",
// 	}
// 	return 0, nil
// }

func EmployeeUpdate(path, token string, moID string, fieldName, fieldValue string) (int, error) {
	payload := &SaveIndicatorInstanceField{
		AuthUserId:                 "4",
		PeriodStart:          time.Now().Format("2006-01-02"),
		PeriodEnd:            time.Now().Format("2006-01-02"),
		PeriodKey:            "month",
		IndicatorToMoId:      moID,
		ApplyToFuturePeriods: true,
		FieldName:            fieldName,
		FieldValue:           fieldValue,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return -1, errors.New("Body marshalling error")
	}

	reqBody := bytes.NewReader(b)
	req, _ := http.NewRequest(http.MethodPost, path, reqBody)

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
		return -1, errProcessing()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("response status: %v", resp.StatusCode)
		return -1, errors.New("Unauthorized")
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("respons body reading error: %v", err)
		return -1, errProcessing()
	}

	var respBody struct {
		Data struct {
			IndicatorToMoHstID int `json:"indicator_to_mo_hst_id"`
		}
	}
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		log.Printf("body unmarshalling error: %v\n", err)
		return -1, errProcessing()
	}

	return respBody.Data.IndicatorToMoHstID, nil
}
