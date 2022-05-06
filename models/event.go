package models

import (
	"encoding/json"
)

type JSON map[string]interface{}

type Event struct {
	Id              string `json:"id"`
	SpecVersion     string `json:"specVersion"`
	Source          string `json:"source"`
	Type            string `json:"type"`
	Subject         string `json:"subject"`
	Time            string `json:"time"`
	CorrelationID   string `json:"correlationId"`
	DataContentType string `json:"dataContentType"`
	Data            JSON   `json:"data"`
}

var Events = []Event{
	{Id: "1dbd06d7-fece-43b4-bbab-072f8a0e80c2", SpecVersion: "1.0", Source: "/product/domain/subdomain/service", Type: "br.com.example.correctTopic", Time: "2022-03-22T17:41:02", CorrelationID: "", DataContentType: "application/json"},
	{Id: "df8cc9e0-ecf2-429f-91b6-8978d7cd1985", SpecVersion: "1.0", Source: "/product/domain/subdomain/service", Type: "br.com.example.correctTopic", Time: "2022-03-22T17:41:02", CorrelationID: "", DataContentType: "application/json"},
	{Id: "5df75212-d861-4ba5-88cd-5abf95685f76", SpecVersion: "1.0", Source: "/product/domain/subdomain/service", Type: "br.com.example.correctTopic", Time: "2022-03-22T17:41:02", CorrelationID: "", DataContentType: "application/json"},
	{Id: "cc654878-c180-4544-b6fd-930efe6e9f5a", SpecVersion: "1.0", Source: "/product/domain/subdomain/service", Type: "br.com.example.correctTopic", Time: "2022-03-22T17:41:02", CorrelationID: "", DataContentType: "application/json"},
}

func EventToJson(event Event) []byte {
	jsonEvent, err := json.Marshal(&event)

	if err != nil {
		panic("Error Enconding EventToJson")
	}

	return jsonEvent
}
