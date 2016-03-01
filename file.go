package gothreat

import (
	"encoding/json"
)

type FileData struct {
	ResponseCode string   `json:"response_code"`
    Md5          string   `json:"md5"`
    Sha1         stinrg   `json:"sha1"`
	Scans        []string `json:"scans"`
    Ips          []string `json:"ips"`
    Domains      []string `json:"domains"`
	References   []string `json:"references"`
	Permalink    string   `json:"permalink"`
}

func FileReportRaw(file string) ([]byte, error) {
	return process_report("file", "resource", file)
}

func FileReport(file string) (AntiVirusData, error) {
	var fileData FileData
	data, err := FileReportRaw(file)

	if err != nil {
		return avData, err
	}

	json.Unmarshal(data, &fileData)

	return fileData, nil
}
