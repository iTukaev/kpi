package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// var text = "{\"auth_user_id\":\"4\",\"indicator_interpretation_id\":\"0\",\"indicator_interpretation_area_id\":\"2\","+
// 	"\"protected\":\"protected\",\"scale\":false,\"name\":\"some_name_probe\",\"indicator_interpretation_point_id\":0,"+
// 	"\"point_x\":0,\"point_y\":10,\"indicator_interpretation_point_id\":0,\"point_x\":50,\"point_y\":60,"+
// 	"\"indicator_interpretation_point_id\":0,\"point_x\":100,\"point_y\":110}"

type SalaryFunc struct {
	X  []string
	Fx []string
}
type Dependens struct {
	Indicator string `json:"indicator_interpretation_point_id"`
	X         string `json:"point_x"`
	Y         string `json:"point_y"`
}

type Construct struct {
	AuthUserId                    string `json:"auth_user_id"`
	IndicatorInterpretationID     string `json:"indicator_interpretation_id"`
	IndicatorInterpretationAreaID string `json:"indicator_interpretation_area_id"`
	Protected                     string `json:"protected"`
	Scale                         bool   `json:"scale"`
	Name                          string `json:"name"`
	Dependencies                  []Dependens
}

func SetFunc(path, token string) {
	sf := SalaryFunc{
		X:  []string{"0", "50", "100", "150", "200"},
		Fx: []string{"10", "60", "110", "160", "210"},
	}

	dependencies := []Dependens{}

	for i := range sf.X {
		dep := Dependens{
			Indicator: "0",
			X:         sf.X[i],
			Y:         sf.Fx[i],
		}
		dependencies = append(dependencies, dep)
	}

	construct := Construct{
		AuthUserId:                    "4",
		IndicatorInterpretationID:     "0",
		IndicatorInterpretationAreaID: "2",
		Protected:                     "protected",
		Scale:                         false,
		Name:                          "some_name_probe",
		Dependencies:                  dependencies,
	}

	body, _ := json.Marshal(construct)

	reqBody := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, path, reqBody)

	fmt.Println(string(body))

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	client := NewClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("response status: %v", resp.StatusCode)
		return
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("respons body reading error: %v", err)
		return
	}

	var respBody struct {
		IndicatorToMoHstID int `json:"indicator_interpretation_id"`
	}

	fmt.Println(string(respBodyBytes), resp.Status)
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		log.Printf("body unmarshalling error: %v\n", err)
		return
	}

	fmt.Println(respBody.IndicatorToMoHstID)
}

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
