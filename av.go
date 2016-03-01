package gothreat

import (
	"encoding/json"
)

type AntiVirusData struct {
	ResponseCode string   `json:"response_code"`
	Md5          string   `json:"md5"`
	Sha1         string   `json:"sha1"`
	Scans        []string `json:"scans"`
	Ips          []string `json:"ips"`
	Domains      []string `json:"domains"`
	References   []string `json:"references"`
	Permalink    string   `json:"permalink"`
}

func AntiVirusReportRaw(av string) ([]byte, error) {
	return process_report("antivirus", av)
}

func AntiVirusReport(av string) (AntiVirusData, error) {
	var avData AntiVirusData
	data, err := AntiVirusReportRaw(av)

	if err != nil {
		return avData, err
	}

	json.Unmarshal(data, &avData)

	return avData, nil
}
