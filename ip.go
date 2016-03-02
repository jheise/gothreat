package gothreat

/*
ThreatCrowd ip report, based on
https://www.threatcrowd.org/searchApi/v2/ip/report/?ip=188.40.75.132
*/

import (
	"encoding/json"
)

type IPData struct {
	ResponseCode string `json:"response_code"`
	Resolutions  []struct {
		LastResolved string `json:"last_resolved"`
		Domain       string `json:"domain"`
	} `json:"resolutions"`
	Hashes     []string `json:"hashes"`
	References []string `json:"references"`
	Permalink  string   `json:"permalink"`
}

func IPReportRaw(ipaddr string) ([]byte, error) {
	return process_report("ip", "ip", ipaddr)
}

func IPReport(ipaddr string) (IPData, error) {
	var ipData IPData
	data, err := IPReportRaw(ipaddr)

	if err != nil {
		return ipData, err
	}

	json.Unmarshal(data, &ipData)

	return ipData, nil
}
