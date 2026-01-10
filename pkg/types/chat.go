package types

// ChatRequest represents a request for chat completion.
type ChatRequest struct {
	// Model is the ID of the model to use.
	Model string `json:"model"`

	// Messages is the list of messages in the conversation.
	Messages []*Message `json:"messages"`

	// Temperature controls randomness in the response (0.0 to 2.0).
	// Lower values make output more focused and deterministic.
	Temperature *float64 `json:"temperature,omitempty"`

	// MaxTokens is the maximum number of tokens to generate.
	MaxTokens int `json:"max_tokens,omitempty"`

	// TopP controls diversity via nucleus sampling (0.0 to 1.0).
	// Alternative to temperature.
	TopP *float64 `json:"top_p,omitempty"`

	// TopK controls diversity by limiting token choices to top K options.
	TopK int `json:"top_k,omitempty"`

	// N is the number of completions to generate.
	N int `json:"n,omitempty"`

	// Stream enables streaming of partial results.
	Stream bool `json:"stream,omitempty"`

	// Stop is a list of sequences where the API will stop generating.
	Stop []string `json:"stop,omitempty"`

	// PresencePenalty penalizes new tokens based on whether they appear in the text (-2.0 to 2.0).
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`

	// FrequencyPenalty penalizes new tokens based on their frequency in the text (-2.0 to 2.0).
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// LogitBias modifies the likelihood of specified tokens appearing.
	LogitBias map[string]float64 `json:"logit_bias,omitempty"`

	// User is a unique identifier for the end-user (for abuse monitoring).
	User string `json:"user,omitempty"`

	// Tools is a list of tools the model can call.
	Tools []*ToolDefinition `json:"tools,omitempty"`

	// ToolChoice controls which (if any) tool is called by the model.
	ToolChoice interface{} `json:"tool_choice,omitempty"` // Can be string or object

	// Functions is a list of functions the model can call (legacy, use Tools).
	Functions []*FunctionDefinition `json:"functions,omitempty"`

	// FunctionCall controls which (if any) function is called (legacy, use ToolChoice).
	FunctionCall interface{} `json:"function_call,omitempty"` // Can be string or object

	// ResponseFormat specifies the format of the response.
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`

	// Seed for deterministic sampling (if supported by provider).
	Seed *int `json:"seed,omitempty"`

	// LogProbs returns log probabilities of output tokens (if supported).
	LogProbs bool `json:"logprobs,omitempty"`

	// TopLogProbs is the number of most likely tokens to return at each position.
	TopLogProbs int `json:"top_logprobs,omitempty"`

	// Metadata contains additional request metadata.
	Metadata *RequestMetadata `json:"metadata,omitempty"`
}

// ChatResponse represents a response from chat completion.
type ChatResponse struct {
	// ID is a unique identifier for the response.
	ID string `json:"id"`

	// Object is the object type (e.g., "chat.completion").
	Object string `json:"object"`

	// Created is the Unix timestamp when the response was created.
	Created int64 `json:"created"`

	// Model is the model used to generate the response.
	Model string `json:"model"`

	// Choices contains the generated completions.
	Choices []*Choice `json:"choices"`

	// Usage contains token usage information.
	Usage *Usage `json:"usage,omitempty"`

	// SystemFingerprint is a fingerprint of the system configuration.
	SystemFingerprint string `json:"system_fingerprint,omitempty"`

	// Metadata contains additional response metadata.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
}

// Choice represents a single completion choice in a chat response.
type Choice struct {
	// Index is the index of this choice in the list.
	Index int `json:"index"`

	// Message is the generated message.
	Message *Message `json:"message"`

	// Delta contains the message delta (used in streaming).
	Delta *Message `json:"delta,omitempty"`

	// FinishReason indicates why the model stopped generating.
	FinishReason FinishReason `json:"finish_reason"`

	// LogProbs contains log probability information (if requested).
	LogProbs *LogProbability `json:"logprobs,omitempty"`
}

// LogProbability contains log probability information for tokens.
type LogProbability struct {
	// Content contains log probabilities for each token in the content.
	Content []*TokenLogProb `json:"content,omitempty"`
}

// TokenLogProb contains log probability information for a single token.
type TokenLogProb struct {
	// Token is the token string.
	Token string `json:"token"`

	// LogProb is the log probability of the token.
	LogProb float64 `json:"logprob"`

	// Bytes is the byte representation of the token (if available).
	Bytes []int `json:"bytes,omitempty"`

	// TopLogProbs contains the most likely tokens at this position.
	TopLogProbs []*TopLogProb `json:"top_logprobs,omitempty"`
}

// TopLogProb represents a top log probability candidate.
type TopLogProb struct {
	// Token is the token string.
	Token string `json:"token"`

	// LogProb is the log probability of the token.
	LogProb float64 `json:"logprob"`

	// Bytes is the byte representation of the token (if available).
	Bytes []int `json:"bytes,omitempty"`
}

// GetFirstMessage returns the message from the first choice, or nil if no choices.
func (r *ChatResponse) GetFirstMessage() *Message {
	if len(r.Choices) > 0 {
		return r.Choices[0].Message
	}
	return nil
}

// GetFirstContent returns the content from the first choice as a string.
func (r *ChatResponse) GetFirstContent() string {
	if msg := r.GetFirstMessage(); msg != nil && msg.Content != nil {
		return msg.Content.String()
	}
	return ""
}

// HasToolCalls returns true if the first choice contains tool calls.
func (r *ChatResponse) HasToolCalls() bool {
	if msg := r.GetFirstMessage(); msg != nil {
		return len(msg.ToolCalls) > 0
	}
	return false
}

// GetToolCalls returns the tool calls from the first choice.
func (r *ChatResponse) GetToolCalls() []*ToolCall {
	if msg := r.GetFirstMessage(); msg != nil {
		return msg.ToolCalls
	}
	return nil
}

// Helper functions for creating requests

// NewChatRequest creates a new ChatRequest with required fields.
func NewChatRequest(model string, messages []*Message) *ChatRequest {
	return &ChatRequest{
		Model:    model,
		Messages: messages,
	}
}

// WithTemperature sets the temperature for the request.
func (r *ChatRequest) WithTemperature(temp float64) *ChatRequest {
	r.Temperature = &temp
	return r
}

// WithMaxTokens sets the max tokens for the request.
func (r *ChatRequest) WithMaxTokens(maxTokens int) *ChatRequest {
	r.MaxTokens = maxTokens
	return r
}

// WithStream enables streaming for the request.
func (r *ChatRequest) WithStream(stream bool) *ChatRequest {
	r.Stream = stream
	return r
}

// WithTools sets the tools for the request.
func (r *ChatRequest) WithTools(tools []*ToolDefinition) *ChatRequest {
	r.Tools = tools
	return r
}

// WithToolChoice sets the tool choice for the request.
func (r *ChatRequest) WithToolChoice(choice ToolChoice) *ChatRequest {
	r.ToolChoice = choice.String()
	return r
}

// WithResponseFormat sets the response format for the request.
func (r *ChatRequest) WithResponseFormat(format *ResponseFormat) *ChatRequest {
	r.ResponseFormat = format
	return r
}

// AddMessage adds a message to the request.
func (r *ChatRequest) AddMessage(message *Message) *ChatRequest {
	r.Messages = append(r.Messages, message)
	return r
}
