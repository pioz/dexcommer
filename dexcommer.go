package dexcommer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const endpoint = "https://shareous1.dexcom.com/ShareWebServices/Services"

var dateRegexp = regexp.MustCompile(`^Date\((\d+)\)$`)

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

type glucoseValueResponse struct {
	Date  string `json:"WT"`
	Value int
	Trend string
}

type GlucoseValue struct {
	Date  time.Time
	Value int
	Trend string
}

func getLastestGlucoseValues(sessionId string, minutes int, count int) []GlucoseValue {
	url := fmt.Sprintf("/Publisher/ReadPublisherLatestGlucoseValues?sessionID=%s&minutes=%d&maxCount=%d", sessionId, minutes, count)
	body, err := post(url, nil)
	if err != nil {
		panic(err)
	}

	rawGlucoseValues := make([]glucoseValueResponse, 0, count)
	json.Unmarshal(body, &rawGlucoseValues)

	glucoseValues := make([]GlucoseValue, count)
	for i, value := range rawGlucoseValues {
		captures := dateRegexp.FindStringSubmatch(value.Date)
		if len(captures) != 2 {
			panic("Invalid glucoseValueResponse Date field")
		}
		unixTimestamp, err := strconv.Atoi(captures[1])
		if err != nil {
			panic("Invalid glucoseValueResponse Date field")
		}
		glucoseValues[i].Date = time.UnixMilli(int64(unixTimestamp))
		glucoseValues[i].Value = value.Value
		glucoseValues[i].Trend = value.Trend
	}

	return glucoseValues
}

type Session struct {
	accountId string
	sessionId string
}

func NewSession(username, password, applicationId string) *Session {
	accountId := getAccoundId(username, password, applicationId)
	sessionId := getSessionId(accountId, password, applicationId)
	return &Session{accountId, sessionId}
}

func (session *Session) ReadLastestGlucoseValues(minutes, count int) []GlucoseValue {
	return getLastestGlucoseValues(session.sessionId, minutes, count)
}
