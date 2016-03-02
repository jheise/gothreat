package gothreat

/*
File reports based on
https://www.threatcrowd.org/searchApi/v2/file/report/?resource=ec8c89aa5e521572c74e2dd02a4daf78
*/

import (
	"encoding/json"
)

type FileData struct {
    ResponseCode string   `json:"response_code"`
    Md5          string   `json:"md5"`
    Sha1         string   `json:"sha1"`
    Scans        []string `json:"scans"`
    Ips          []string `json:"ips"`
    Domains      []string `json:"domains"`
    References   []string `json:"references"`
    Permalink    string   `json:"permalink"`
}

func FileReportRaw(file string) ([]byte, error) {
	return process_report("file", "resource", file)
}

func FileReport(file string) (FileData, error) {
	var fileData FileData
	data, err := FileReportRaw(file)

	if err != nil {
		return fileData, err
	}

	json.Unmarshal(data, &fileData)

	return fileData, nil
}
