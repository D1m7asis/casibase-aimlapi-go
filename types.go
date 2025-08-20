package aimlapi

const (
        ModelDeepSeekChat = "deepseek-chat"
        ModelGeminiPro    = "gemini-pro"
        ModelGPT35Turbo   = "gpt-3.5-turbo"
        ModelGPT4         = "gpt-4"
)

type ChatCompletionMessage struct {
        Role    string `json:"role"`
        Content string `json:"content"`
}

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
        Model       string                  `json:"model"`
        Messages    []ChatCompletionMessage `json:"messages"`
        MaxTokens   int                     `json:"max_tokens,omitempty"`
        Stream      bool                    `json:"stream,omitempty"`
        Temperature *float32                `json:"temperature,omitempty"`
        TopP        *float32                `json:"top_p,omitempty"`
        TopK        *uint                   `json:"top_k,omitempty"`
}

type Index struct {
        Role    string `json:"role"`
        Content string `json:"content"`
}

type ChatCompletionChoice struct {
        Message      Index  `json:"message,omitempty"`
        FinishReason string `json:"finish_reason,omitempty"`
        Delta        Index  `json:"delta,omitempty"`
        Index        uint   `json:"index,omitempty"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
        ID      string                 `json:"id,omitempty"`
        Object  string                 `json:"object,omitempty"`
        Created int64                  `json:"created,omitempty"`
        Model   string                 `json:"model"`
        Choices []ChatCompletionChoice `json:"choices"`
        //Usage   Usage                  `json:"usage,omitempty"`
}

type Usage struct {
        CompletionTokens int `json:"completion_tokens"`
        TotalTokens      int `json:"total_tokens"`
}

