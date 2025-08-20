# aimlapi-go

Go client for the [AI/ML API](https://api.aimlapi.com/v1) which provides access to
hundreds of models such as DeepSeek, Gemini, and ChatGPT.

## Usage

```go
package main

import "github.com/casibase/aimlapi-go"

func main() {
    client, _ := aimlapi.NewClient("API_KEY", "MyApp", "https://example.com")
    _ = client // use the client
}
```

