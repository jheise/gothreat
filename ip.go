package gothreat

import (
    "encoding/json"
)

type IPData struct {
    Permalink string
    Response_code int
    References []string
    Resolutions []map[string]string
    Hashes []string
}

func IPReportRaw(ipaddr string) ([]byte, error){
    return process_report("ip", ipaddr)
}

func IPReport(ipaddr string) (IPData, error){
    var ip_data IPData
    data, err := IPReportRaw(ipaddr)

    if err != nil {
        return ip_data, err
    }

    json.Unmarshal(data, &ip_data)

    return ip_data, nil
}
