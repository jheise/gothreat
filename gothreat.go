package gothreat

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func process_report(report_type string, query string) ([]byte, error){
    endpoint := fmt.Sprintf("https://threatcrowd.org/searchApi/v2/%s/report/?%s=%s",report_type, report_type, query)
    resp, err := http.Get(endpoint)

    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, err
}
