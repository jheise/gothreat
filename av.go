package gothreat

/*
ThreatCrowd antivirus report, based on
https://www.threatcrowd.org/searchApi/v2/antivirus/report/?antivirus=plugx
*/

import (
	"encoding/json"
)

type AntiVirusData struct {
	ResponseCode string   `json:"response_code"`
	Hashes       []string `json:"hashes"`
	References   []string `json:"references"`
	Permalink    string   `json:"permalink"`
}

func AntiVirusReportRaw(av string) ([]byte, error) {
	return process_report("antivirus", "antivirus", av)
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
