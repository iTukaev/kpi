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

func errProcessing() error {
	return errors.New("Processing error")
}

func GetAllMo(path, token string) ([]*Mo, error) {
	moPayload := &MoPayload{
		PeriodStart:  time.Now().Format("2006-01-02"),
		PeriodEnd:    time.Now().Format("2006-01-02"),
		PeriodKey:    "month",
		MoChatFilter: false,
	}

	b, err := json.Marshal(moPayload)
	if err != nil {
		return nil, errors.New("Body marshalling error")
	}

	reqBody := bytes.NewReader(b)
	req, _ := http.NewRequest(http.MethodPost, path, reqBody)

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
		return nil, errProcessing()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unauthorized")
	}
	
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("respons body reading error: %v", err)
		return nil, errProcessing()
	}

	respBody := &Body{}
	err = json.Unmarshal(respBodyBytes, respBody)
	if err != nil {
		log.Printf("body unmarshalling error: %v\n", err)
		return nil, errProcessing()
	}

	return respBody.Data.Rows, nil
}
