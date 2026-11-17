package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PlayByPlayData struct {
	Events []Event
}

type Event struct {
	Date     string `json:"date"`
	LastPlay struct {
		Text string `json:"text"`
	}
}

func LatestPlayClient() ([]Event, error) {
	currentDate := time.Now().Format("20060102")
	u, err := url.Parse("https://site.web.api.espn.com/apis/fantasy/v2/games/ffl/games")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	q := u.Query()

	q.Set("dates", fmt.Sprintf("%s-%s", currentDate, currentDate))
	q.Set("pbpOnly", "true")

	u.RawQuery = q.Encode()

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var playByPlayData PlayByPlayData

	err = json.NewDecoder(res.Body).Decode(&playByPlayData)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	events := playByPlayData.Events

	return events, nil
}
