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
package main

import (
	"fmt"
	"github.com/jheise/gothreat"
)

func main() {
	ipData, err := gothreat.IPReport("4.2.2.1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Permalink: %s\n", ipData.Permalink)
	fmt.Printf("DNS Resolutions:\n")
	for _, resolve := range ipData.Resolutions {
		fmt.Printf("\t%v -> %v\n", resolve.LastResolved, resolve.Domain)
	}
	fmt.Printf("References:\n")
	for _, reference := range ipData.References {
		fmt.Printf("\t%s\n", reference)
	}
	fmt.Printf("Hashes:\n")
	for _, hash := range ipData.Hashes {
		fmt.Printf("\t%s\n", hash)
	}
}
```
