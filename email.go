package gothreat

import (
	"encoding/json"
)

type EmailData struct {
	ResponseCode string   `json:"response_code"`
	Domains      []string `json:"domains"`
	References   []string `json:"references"`
	Permalink    string   `json:"permalink"`
}

func EmailReportRaw(email string) ([]byte, error) {
	return process_report("email", email)
}

func EmailReport(email string) (EmailData, error) {
	var emailData EmailData
	data, err := EmailReportRaw(email)

	if err != nil {
		return emailData, err
	}

	json.Unmarshal(data, &emailData)

	return emailData, nil
}
