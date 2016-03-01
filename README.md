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

    fmt.Printf("%s\n", data)
}
```
