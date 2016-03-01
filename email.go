package gothreat

import (
    "encoding/json"
)

type EmailData struct {
    Permalink string
    References []string
    Domains []string
    ResponseCode int
}

func EmailReportRaw(email string) ([]byte, error) {
    return process_report("email", email)
}

func EmailReport(email string) (EmailData, error){
    var email_data EmailData
    data, err := EmailReportRaw(email)

    if err != nil {
        return email_data, err
    }

    json.Unmarshal(data, &email_data)

    return email_data, nil
}
