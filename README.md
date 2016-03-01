# gothreat
---

Go package for ThreatCrowd.org

Functions and Structs for:

- email

- ip

- domain

- antivirus

---

Example

```go
import (
    "github.com/jheise/gothreat"
    "fmt"
)

func main(){
    data, err := gothreat.IPReport("4.2.2.1")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Permalink: %s\n", ip_data.Permalink)
    fmt.Printf("DNS Resolutions:\n")
    for _, resolve := range ip_data.Resolutions {
        fmt.Printf("\t%s -> %s\n", resolve["last_resolved"], resolve["domain"])
    }
    fmt.Printf("References:\n")
    for _, reference := range ip_data.References {
        fmt.Printf("\t%s\n", reference)
    }
    fmt.Printf("Hashes:\n")
    for _, hash := range ip_data.Hashes {
        fmt.Printf("\t%s\n", hash)
    }
}
```
