# aimlapi-go

Go client for the [AI/ML API](https://aimlapi.com/) which provides access to
hundreds of models such as DeepSeek, Gemini, and ChatGPT.

## Usage

```go
package main

import "github.com/D1m7asis/casibase-aimlapi-go"

func main() {
    client, _ := aimlapi.NewClient("API_KEY", "MyApp", "https://example.com")
    _ = client // use the client
}
```

