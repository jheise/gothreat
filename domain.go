package gothreat

import (
    "encoding/json"
)

type DomainData struct {
    Permalink string
    References []string
    Emails []string
    Subdomains []string
    Hashes []string
    Resolutions []map[string]string
    ResponseCode int
}

func DomainReportRaw(domain string) ([]byte, error) {
    return process_report("domain", domain)
}

func DomainReport(domain string) (DomainData, error){
    var domain_data DomainData
    data, err := DomainReportRaw(domain)

    if err != nil {
        return domain_data, err
    }

    json.Unmarshal(data, &domain_data)

    return domain_data, nil
}
