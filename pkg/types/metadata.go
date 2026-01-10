package types

import "time"

// Usage represents token usage information for API requests.
type Usage struct {
	// PromptTokens is the number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`

	// CompletionTokens is the number of tokens in the completion.
	CompletionTokens int `json:"completion_tokens"`

	// TotalTokens is the total number of tokens used (prompt + completion).
	TotalTokens int `json:"total_tokens"`

	// CachedTokens is the number of tokens served from cache (if applicable).
	// Some providers like Anthropic support prompt caching.
	CachedTokens int `json:"cached_tokens,omitempty"`

	// ReasoningTokens is the number of tokens used for reasoning (if applicable).
	// Some models like o1 use separate reasoning tokens.
	ReasoningTokens int `json:"reasoning_tokens,omitempty"`
}

// Add adds usage statistics from another Usage instance.
func (u *Usage) Add(other *Usage) {
	if other == nil {
		return
	}
	u.PromptTokens += other.PromptTokens
	u.CompletionTokens += other.CompletionTokens
	u.TotalTokens += other.TotalTokens
	u.CachedTokens += other.CachedTokens
	u.ReasoningTokens += other.ReasoningTokens
}

// MessageMetadata contains metadata about a message.
type MessageMetadata struct {
	// ID is a unique identifier for the message.
	ID string `json:"id,omitempty"`

	// Timestamp is when the message was created.
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Model is the model that generated the message (for assistant messages).
	Model string `json:"model,omitempty"`

	// Custom holds custom metadata fields.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// ResponseMetadata contains metadata about an API response.
type ResponseMetadata struct {
	// ID is a unique identifier for the response.
	ID string `json:"id"`

	// Created is the Unix timestamp when the response was created.
	Created int64 `json:"created"`

	// Model is the model used to generate the response.
	Model string `json:"model"`

	// Provider is the AI provider that handled the request.
	Provider Provider `json:"provider,omitempty"`

	// SystemFingerprint is a fingerprint of the system configuration.
	// Used to track backend changes.
	SystemFingerprint string `json:"system_fingerprint,omitempty"`

	// Custom holds custom metadata fields.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// RequestMetadata contains metadata about an API request.
type RequestMetadata struct {
	// ID is a unique identifier for the request.
	ID string `json:"id,omitempty"`

	// Timestamp is when the request was created.
	Timestamp time.Time `json:"timestamp,omitempty"`

	// UserID is the ID of the user making the request.
	UserID string `json:"user_id,omitempty"`

	// SessionID is the ID of the session this request belongs to.
	SessionID string `json:"session_id,omitempty"`

	// TraceID is a distributed tracing ID for request correlation.
	TraceID string `json:"trace_id,omitempty"`

	// Custom holds custom metadata fields.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// ModelInfo contains information about an AI model.
type ModelInfo struct {
	// ID is the model identifier.
	ID string `json:"id"`

	// Name is the human-readable model name.
	Name string `json:"name,omitempty"`

	// Provider is the provider that offers this model.
	Provider Provider `json:"provider"`

	// Capabilities lists the capabilities this model supports.
	Capabilities []ModelCapability `json:"capabilities,omitempty"`

	// MaxTokens is the maximum number of tokens the model supports.
	MaxTokens int `json:"max_tokens,omitempty"`

	// MaxInputTokens is the maximum input tokens (if different from MaxTokens).
	MaxInputTokens int `json:"max_input_tokens,omitempty"`

	// MaxOutputTokens is the maximum output tokens.
	MaxOutputTokens int `json:"max_output_tokens,omitempty"`

	// SupportsSystemMessages indicates if the model supports system messages.
	SupportsSystemMessages bool `json:"supports_system_messages,omitempty"`

	// SupportsImages indicates if the model supports image inputs.
	SupportsImages bool `json:"supports_images,omitempty"`

	// SupportsAudio indicates if the model supports audio inputs.
	SupportsAudio bool `json:"supports_audio,omitempty"`

	// Description is a description of the model.
	Description string `json:"description,omitempty"`
}

// HasCapability checks if the model has a specific capability.
func (m *ModelInfo) HasCapability(capability ModelCapability) bool {
	for _, c := range m.Capabilities {
		if c == capability {
			return true
		}
	}
	return false
}
