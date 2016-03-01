package gothreat

import (
	"encoding/json"
)

type DomainData struct {
	ResponseCode string `json:"response_code"`
	Resolutions  []struct {
		LastResolved string `json:"last_resolved"`
		IPAddress    string `json:"ip_address"`
	} `json:"resolutions"`
	Hashes     []string `json:"hashes"`
	Emails     []string `json:"emails"`
	Subdomains []string `json:"subdomains"`
	References []string `json:"references"`
	Permalink  string   `json:"permalink"`
}

func DomainReportRaw(domain string) ([]byte, error) {
	return process_report("domain", domain)
}

func DomainReport(domain string) (DomainData, error) {
	var domainData DomainData
	data, err := DomainReportRaw(domain)

	if err != nil {
		return domainData, err
	}

	json.Unmarshal(data, &domainData)

	return domainData, nil
}
