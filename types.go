package aimlapi

const (
    // Alibaba Cloud
    ModelQwen3_32B              = "alibaba/qwen3-32b"
    ModelQwen3Coder480bA35b     = "alibaba/qwen3-coder-480b-a35b-instruct"
    ModelQwen3Thinking2507      = "alibaba/qwen3-235b-a22b-thinking-2507"
    ModelQwenMax                = "alibaba/qwen-max"
    ModelQwenMax20250125        = "alibaba/qwen-max-2025-01-25"
    ModelQwenPlus               = "alibaba/qwen-plus"
    ModelQwenTurbo              = "alibaba/qwen-turbo"
    ModelQwen3_235B_A22B        = "Qwen/Qwen3-235B-A22B-fp8-tput"

    // Anthracite
    ModelMagnumV4_72B           = "anthracite-org/magnum-v4-72b"

    // Anthropic
    ModelClaude3Haiku20240307   = "claude-3-haiku-20240307"
    ModelClaude3Opus20240229    = "claude-3-opus-20240229"
    ModelClaude35Haiku20241022  = "claude-3-5-haiku-20241022"
    ModelClaude35Sonnet20240620 = "claude-3-5-sonnet-20240620"
    ModelClaude35Sonnet20241022 = "claude-3-5-sonnet-20241022"
    ModelClaude37Sonnet20250219 = "claude-3-7-sonnet-20250219"
    ModelClaude4Opus20250514    = "claude-opus-4-20250514"
    ModelClaude4Sonnet20250514  = "claude-sonnet-4-20250514"
    ModelClaudeOpus41_20250805  = "claude-opus-4-1-20250805"

    // Bagoodex
    ModelBagoodexSearchV1       = "bagoodex/bagoodex-search-v1"

    // Cohere
    ModelCommandA               = "cohere/command-a"
    ModelCommandRPlus           = "cohere/command-r-plus"

    // DeepSeek AI
    ModelDeepSeekProverV2       = "deepseek/deepseek-prover-v2"
    ModelDeepSeekR1             = "deepseek-reasoner"
    ModelDeepSeekChat           = "deepseek-chat"

    // Google
    ModelGemini20Flash          = "google/gemini-2.0-flash"
    ModelGemini20FlashExp       = "google/gemini-2.0-flash-exp"
    ModelGemini25Flash          = "google/gemini-2.5-flash"
    ModelGemini25FlashLite      = "google/gemini-2.5-flash-lite-preview"
    ModelGemini25Pro            = "google/gemini-2.5-pro"
    ModelGemma227BIt            = "google/gemma-2-27b-it"
    ModelGemma312BIt            = "google/gemma-3-12b-it"
    ModelGemma327BIt            = "google/gemma-3-27b-it"
    ModelGemma34BIt             = "google/gemma-3-4b-it"
    ModelGemma3n4BIt            = "google/gemma-3n-e4b-it"

    // Meta
    ModelLlama370BChatHF        = "meta-llama/Llama-3-70b-chat-hf"
    ModelLlama38BChatHF         = "meta-llama/Llama-3-8b-chat-hf"
    ModelLlama32_11BVision      = "meta-llama/Llama-3.2-11B-Vision-Instruct-Turbo"
    ModelLlama32_3BInstruct     = "meta-llama/Llama-3.2-3B-Instruct-Turbo"
    ModelLlama32_90BVision      = "meta-llama/Llama-3.2-90B-Vision-Instruct-Turbo"
    ModelLlama33_70BInstruct    = "meta-llama/Llama-3.3-70B-Instruct-Turbo"
    ModelLlama4Maverick         = "meta-llama/llama-4-maverick"
    ModelLlama4Scout            = "meta-llama/llama-4-scout"
    ModelLlamaGuard311BVision   = "meta-llama/Llama-Guard-3-11B-Vision-Turbo"
    ModelLlamaVisionFree        = "meta-llama/Llama-Vision-Free"
    ModelLlamaGuard28B          = "meta-llama/LlamaGuard-2-8b"
    ModelMetaLlama38BLite       = "meta-llama/Meta-Llama-3-8B-Instruct-Lite"
    ModelMetaLlama31_405B       = "meta-llama/Meta-Llama-3.1-405B-Instruct-Turbo"
    ModelMetaLlama31_70B        = "meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo"
    ModelMetaLlama31_8B         = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
    ModelMetaLlamaGuard38B      = "meta-llama/Meta-Llama-Guard-3-8B"

    // Minimax AI
    ModelMinimaxM1              = "minimax/m1"
    ModelMiniMaxText01          = "MiniMax-Text-01"

    // Mistral AI
    ModelCodestral2501          = "mistralai/codestral-2501"
    ModelMistral7Bv01           = "mistralai/Mistral-7B-Instruct-v0.1"
    ModelMistral7Bv02           = "mistralai/Mistral-7B-Instruct-v0.2"
    ModelMistral7Bv03           = "mistralai/Mistral-7B-Instruct-v0.3"
    ModelMistralNemo            = "mistralai/mistral-nemo"
    ModelMistralTiny            = "mistralai/mistral-tiny"
    ModelMixtral8x7Bv01         = "mistralai/Mixtral-8x7B-Instruct-v0.1"

    // Moonshot
    ModelKimiK2Preview          = "moonshot/kimi-k2-preview"

    // NousResearch
    ModelNousHermes2Mixtral     = "NousResearch/Nous-Hermes-2-Mixtral-8x7B-DPO"

    // Nvidia
    ModelLlama31Nemotron70B     = "nvidia/llama-3.1-nemotron-70b-instruct"
    ModelLlama31Nemotron70BHF   = "nvidia/Llama-3.1-Nemotron-70B-Instruct-HF"

    // OpenAI
    ModelGPT41_20250414         = "openai/gpt-4.1-2025-04-14"
    ModelGPT41Mini20250414      = "openai/gpt-4.1-mini-2025-04-14"
    ModelGPT41Nano20250414      = "openai/gpt-4.1-nano-2025-04-14"
    ModelGPT4oAudioPreview      = "gpt-4o-audio-preview"
    ModelGPT4oMiniAudioPreview  = "gpt-4o-mini-audio-preview"
    ModelGPT4oMiniSearchPreview = "gpt-4o-mini-search-preview"
    ModelGPT4oSearchPreview     = "gpt-4o-search-preview"
    ModelChatGPT4oLatest        = "chatgpt-4o-latest"
    ModelGPT35Turbo             = "gpt-3.5-turbo"
    ModelGPT35Turbo0125         = "gpt-3.5-turbo-0125"
    ModelGPT35Turbo1106         = "gpt-3.5-turbo-1106"
    ModelGPT4                   = "gpt-4"
    ModelGPT40125Preview        = "gpt-4-0125-preview"
    ModelGPT41106Preview        = "gpt-4-1106-preview"
    ModelGPT4Turbo              = "gpt-4-turbo"
    ModelGPT4Turbo20240409      = "gpt-4-turbo-2024-04-09"
    ModelGPT4o                  = "openai/gpt-4o"
    ModelGPT4o20240513          = "gpt-4o-2024-05-13"
    ModelGPT4o20240806          = "gpt-4o-2024-08-06"
    ModelGPT4oMini              = "gpt-4o-mini"
    ModelGPT4oMini20240718      = "gpt-4o-mini-2024-07-18"
    ModelGPTOss120B             = "openai/gpt-oss-120b"
    ModelGPTOss20B              = "openai/gpt-oss-20b"
    ModelGPT5_20250807          = "openai/gpt-5-2025-08-07"
    ModelGPT5ChatLatest         = "openai/gpt-5-chat-latest"
    ModelGPT5Mini20250807       = "openai/gpt-5-mini-2025-08-07"
    ModelGPT5Nano20250807       = "openai/gpt-5-nano-2025-08-07"
    ModelO1                     = "o1"
    ModelO1Mini                 = "o1-mini"
    ModelO1Mini20240912         = "o1-mini-2024-09-12"
    ModelO3_20250416            = "openai/o3-2025-04-16"
    ModelO3Mini                 = "o3-mini"
    ModelO3Pro                  = "openai/o3-pro"
    ModelO4Mini20250416         = "openai/o4-mini-2025-04-16"

    // Perplexity
    ModelSonar                  = "perplexity/sonar"
    ModelSonarPro               = "perplexity/sonar-pro"

    // Qwen
    ModelQwQ32B                 = "Qwen/QwQ-32B"
    ModelQwen272BInstruct       = "Qwen/Qwen2-72B-Instruct"
    ModelQwen25_72BInstruct     = "Qwen/Qwen2.5-72B-Instruct-Turbo"
    ModelQwen25_7BInstruct      = "Qwen/Qwen2.5-7B-Instruct-Turbo"
    ModelQwen25Coder32B         = "Qwen/Qwen2.5-Coder-32B-Instruct"

    // X AI
    ModelGrok3Beta              = "x-ai/grok-3-beta"
    ModelGrok3MiniBeta          = "x-ai/grok-3-mini-beta"
    ModelGrok4_0709             = "x-ai/grok-4-07-09"

    // Zhipu AI
    ModelGLM45                  = "zhipu/glm-4.5"
    ModelGLM45Air               = "zhipu/glm-4.5-air"
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

