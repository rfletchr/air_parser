# AIR Parser
A Go parser for [Civitai AIR](https://developer.civitai.com/docs/getting-started/ai-resource-identifier) URNs.

e.g. 
```
urn:air:sd1:model:civitai:2421@43533
```


## Usage 
```go
package main

import (
    "fmt"
    "github.com/rfletchr/air_parser"
)

func main() {
    resource, err := air_parser.NewAirFromString("urn:air:sd1:model:civitai:2421@43533")
    if err != nil {
        panic(err)
    }
    fmt.Printf("result: %v\n", resource)
}
```



