package aimlapi

import (
        "context"
        "errors"
        "net/http"
)

// Chat message role defined by the API.

type ModelName string

const (
        ChatMessageRoleUser      = "user"
        ChatMessageRoleSystem    = "system"
        ChatMessageRoleAssistant = "assistant"
)

var (
        ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream") //nolint:lll
)

// CreateChatCompletion — API call to create a completion for the chat message.
func (c *Client) CreateChatCompletion(
        ctx context.Context,
        request *ChatCompletionRequest,
) (response *ChatCompletionResponse, err error) {
        if request.Stream {
                err = ErrChatCompletionStreamNotSupported
                return
        }

        urlSuffix := "/chat/completions"

        req, err := c.requestBuilder.Build(ctx, http.MethodPost, c.fullURL(urlSuffix), request)
        if err != nil {
                return
        }

        err = c.sendRequest(req, &response)
        return
}

