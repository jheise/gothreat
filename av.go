package gothreat

import (
    "encoding/json"
)

type AntiVirusData struct {
    Permalink string
    References []string
    Hashes []string
    ResponseCode int
}

func AntiVirusReportRaw(av string) ([]byte, error) {
    return process_report("antivirus", av)
}

func AntiVirusReport(av string) (AntiVirusData, error){
    var av_data AntiVirusData
    data, err := AntiVirusReportRaw(av)

    if err != nil {
        return av_data, err
    }

    json.Unmarshal(data, &av_data)

    return av_data, nil
}
