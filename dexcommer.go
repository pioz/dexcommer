package dexcommer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const endpoint = "https://shareous1.dexcom.com/ShareWebServices/Services"

func post(path string, params map[string]string) ([]byte, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", endpoint+path, bytes.NewBuffer(jsonParams))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "share2nightscout-bridge/0.2.8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status + ": " + string(body))
	}
	return body, nil
}

func getAccoundId(username, password, applicationId string) string {
	params := map[string]string{
		"accountName":   username,
		"password":      password,
		"applicationId": applicationId,
	}
	body, err := post("/General/AuthenticatePublisherAccount", params)
	if err != nil {
		panic(err)
	}
	s := string(body)
	return s[1 : len(s)-1]
}

func getSessionId(accountId, password, applicationId string) string {
	params := map[string]string{
		"accountId":     accountId,
		"password":      password,
		"applicationId": applicationId,
	}
	body, err := post("/General/LoginPublisherAccountById", params)
	if err != nil {
		panic(err)
	}
	s := string(body)
	return s[1 : len(s)-1]
}

func getLastestGlucoseValues(sessionId string) string {
	body, err := post("/Publisher/ReadPublisherLatestGlucoseValues?sessionID="+sessionId+"&minutes=1440&maxCount=6", nil)
	if err != nil {
		panic(err)
	}
	s := string(body)
	return s
}

// ReadLastestGlucoseValues fetch the latest glucose values using your Dexcom
// Follow credentials. If you do not have an applicationId you can use
// Nightscout's applicationId which is "d89443d2-327c-4a6f-89e5-496bbb0317db".
func ReadLastestGlucoseValues(username, password, applicationId string) string {
	accountId := getAccoundId(username, password, applicationId)
	sessionId := getSessionId(accountId, password, applicationId)
	return getLastestGlucoseValues(sessionId)
}
