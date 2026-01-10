package types

// StreamChunk represents a chunk of data in a streaming response.
type StreamChunk interface {
	// GetID returns the unique identifier for the stream.
	GetID() string

	// GetModel returns the model used for generation.
	GetModel() string

	// GetChoices returns the choices in this chunk.
	GetChoices() []*StreamChoice

	// IsComplete returns true if this is the final chunk.
	IsComplete() bool
}

// ChatStreamChunk represents a chunk in a chat completion stream.
type ChatStreamChunk struct {
	// ID is a unique identifier for the stream.
	ID string `json:"id"`

	// Object is the object type (e.g., "chat.completion.chunk").
	Object string `json:"object"`

	// Created is the Unix timestamp when the chunk was created.
	Created int64 `json:"created"`

	// Model is the model used to generate the response.
	Model string `json:"model"`

	// Choices contains the delta choices for this chunk.
	Choices []*StreamChoice `json:"choices"`

	// Usage contains token usage information (typically only in final chunk).
	Usage *Usage `json:"usage,omitempty"`

	// SystemFingerprint is a fingerprint of the system configuration.
	SystemFingerprint string `json:"system_fingerprint,omitempty"`
}

// GetID returns the stream ID.
func (c *ChatStreamChunk) GetID() string {
	return c.ID
}

// GetModel returns the model name.
func (c *ChatStreamChunk) GetModel() string {
	return c.Model
}

// GetChoices returns the stream choices.
func (c *ChatStreamChunk) GetChoices() []*StreamChoice {
	return c.Choices
}

// IsComplete returns true if this chunk marks the end of the stream.
func (c *ChatStreamChunk) IsComplete() bool {
	if len(c.Choices) == 0 {
		return false
	}
	// Check if any choice has a finish reason
	for _, choice := range c.Choices {
		if choice.FinishReason != "" && choice.FinishReason != FinishReasonNull {
			return true
		}
	}
	return false
}

// GetFirstDelta returns the delta from the first choice.
func (c *ChatStreamChunk) GetFirstDelta() *MessageDelta {
	if len(c.Choices) > 0 {
		return c.Choices[0].Delta
	}
	return nil
}

// GetFirstContent returns the content from the first delta (if available).
func (c *ChatStreamChunk) GetFirstContent() string {
	if delta := c.GetFirstDelta(); delta != nil {
		return delta.Content
	}
	return ""
}

// StreamChoice represents a choice in a streaming response.
type StreamChoice struct {
	// Index is the index of this choice in the list.
	Index int `json:"index"`

	// Delta contains the incremental changes to the message.
	Delta *MessageDelta `json:"delta"`

	// FinishReason indicates why the model stopped generating.
	// Empty string means streaming is still in progress.
	FinishReason FinishReason `json:"finish_reason"`

	// LogProbs contains log probability information (if requested).
	LogProbs *LogProbability `json:"logprobs,omitempty"`
}

// MessageDelta represents incremental changes to a message during streaming.
type MessageDelta struct {
	// Role is the role of the message (typically only in first chunk).
	Role Role `json:"role,omitempty"`

	// Content is the incremental text content.
	Content string `json:"content,omitempty"`

	// ToolCalls contains incremental tool call updates.
	ToolCalls []*ToolCallDelta `json:"tool_calls,omitempty"`

	// FunctionCall contains incremental function call updates (legacy).
	FunctionCall *FunctionCallDelta `json:"function_call,omitempty"`
}

// ToolCallDelta represents incremental updates to a tool call.
type ToolCallDelta struct {
	// Index is the index of this tool call.
	Index int `json:"index"`

	// ID is the tool call ID (typically only in first chunk).
	ID string `json:"id,omitempty"`

	// Type is the tool type (typically only in first chunk).
	Type ToolType `json:"type,omitempty"`

	// Function contains incremental function call updates.
	Function *FunctionCallDelta `json:"function,omitempty"`
}

// FunctionCallDelta represents incremental updates to a function call.
type FunctionCallDelta struct {
	// Name is the function name (typically only in first chunk).
	Name string `json:"name,omitempty"`

	// Arguments contains incremental argument JSON.
	Arguments string `json:"arguments,omitempty"`
}

// StreamEvent represents an event in a server-sent events stream.
type StreamEvent struct {
	// Event is the event type (e.g., "message", "content_block_delta", "done").
	Event string `json:"event"`

	// Data contains the event data.
	Data interface{} `json:"data,omitempty"`

	// Error contains error information (if Event is "error").
	Error *ProviderError `json:"error,omitempty"`
}

// StreamError represents an error that occurred during streaming.
type StreamError struct {
	// Message is the error message.
	Message string `json:"message"`

	// Type is the error type.
	Type ErrorType `json:"type"`

	// Code is the error code.
	Code string `json:"code,omitempty"`
}

// Error implements the error interface.
func (e *StreamError) Error() string {
	if e.Code != "" {
		return e.Type.String() + ": " + e.Message + " (" + e.Code + ")"
	}
	return e.Type.String() + ": " + e.Message
}

// StreamAccumulator helps accumulate streaming chunks into a complete response.
type StreamAccumulator struct {
	// ID is the stream ID.
	ID string

	// Model is the model name.
	Model string

	// Created is the creation timestamp.
	Created int64

	// Choices accumulates the choices.
	Choices map[int]*AccumulatedChoice

	// Usage accumulates usage information.
	Usage *Usage

	// SystemFingerprint is the system fingerprint.
	SystemFingerprint string
}

// AccumulatedChoice represents an accumulated choice during streaming.
type AccumulatedChoice struct {
	// Index is the choice index.
	Index int

	// Role is the message role.
	Role Role

	// Content accumulates the text content.
	Content string

	// ToolCalls accumulates tool calls.
	ToolCalls map[int]*AccumulatedToolCall

	// FinishReason is the finish reason (set in final chunk).
	FinishReason FinishReason
}

// AccumulatedToolCall represents an accumulated tool call.
type AccumulatedToolCall struct {
	// Index is the tool call index.
	Index int

	// ID is the tool call ID.
	ID string

	// Type is the tool type.
	Type ToolType

	// FunctionName is the function name.
	FunctionName string

	// Arguments accumulates the function arguments JSON.
	Arguments string
}

// NewStreamAccumulator creates a new StreamAccumulator.
func NewStreamAccumulator() *StreamAccumulator {
	return &StreamAccumulator{
		Choices: make(map[int]*AccumulatedChoice),
	}
}

// Add processes a stream chunk and updates the accumulator.
func (a *StreamAccumulator) Add(chunk StreamChunk) {
	if c, ok := chunk.(*ChatStreamChunk); ok {
		a.ID = c.ID
		a.Model = c.Model
		a.Created = c.Created
		a.SystemFingerprint = c.SystemFingerprint

		if c.Usage != nil {
			if a.Usage == nil {
				a.Usage = &Usage{}
			}
			a.Usage.Add(c.Usage)
		}

		for _, choice := range c.Choices {
			if choice.Delta == nil {
				continue
			}

			idx := choice.Index
			if _, exists := a.Choices[idx]; !exists {
				a.Choices[idx] = &AccumulatedChoice{
					Index:     idx,
					ToolCalls: make(map[int]*AccumulatedToolCall),
				}
			}

			accChoice := a.Choices[idx]

			// Accumulate role (typically only in first chunk)
			if choice.Delta.Role != "" {
				accChoice.Role = choice.Delta.Role
			}

			// Accumulate content
			if choice.Delta.Content != "" {
				accChoice.Content += choice.Delta.Content
			}

			// Accumulate tool calls
			for _, toolCallDelta := range choice.Delta.ToolCalls {
				if _, exists := accChoice.ToolCalls[toolCallDelta.Index]; !exists {
					accChoice.ToolCalls[toolCallDelta.Index] = &AccumulatedToolCall{
						Index: toolCallDelta.Index,
					}
				}

				accTool := accChoice.ToolCalls[toolCallDelta.Index]

				if toolCallDelta.ID != "" {
					accTool.ID = toolCallDelta.ID
				}
				if toolCallDelta.Type != "" {
					accTool.Type = toolCallDelta.Type
				}
				if toolCallDelta.Function != nil {
					if toolCallDelta.Function.Name != "" {
						accTool.FunctionName = toolCallDelta.Function.Name
					}
					if toolCallDelta.Function.Arguments != "" {
						accTool.Arguments += toolCallDelta.Function.Arguments
					}
				}
			}

			// Set finish reason
			if choice.FinishReason != "" && choice.FinishReason != FinishReasonNull {
				accChoice.FinishReason = choice.FinishReason
			}
		}
	}
}

// ToChatResponse converts the accumulated data to a ChatResponse.
func (a *StreamAccumulator) ToChatResponse() *ChatResponse {
	choices := make([]*Choice, 0, len(a.Choices))
	for _, accChoice := range a.Choices {
		message := &Message{
			Role:    accChoice.Role,
			Content: NewTextContent(accChoice.Content),
		}

		// Convert tool calls
		if len(accChoice.ToolCalls) > 0 {
			toolCalls := make([]*ToolCall, 0, len(accChoice.ToolCalls))
			for _, accTool := range accChoice.ToolCalls {
				toolCalls = append(toolCalls, &ToolCall{
					ID:   accTool.ID,
					Type: accTool.Type,
					Function: FunctionCall{
						Name:      accTool.FunctionName,
						Arguments: accTool.Arguments,
					},
					Index: accTool.Index,
				})
			}
			message.ToolCalls = toolCalls
		}

		choices = append(choices, &Choice{
			Index:        accChoice.Index,
			Message:      message,
			FinishReason: accChoice.FinishReason,
		})
	}

	return &ChatResponse{
		ID:                a.ID,
		Object:            "chat.completion",
		Created:           a.Created,
		Model:             a.Model,
		Choices:           choices,
		Usage:             a.Usage,
		SystemFingerprint: a.SystemFingerprint,
	}
}
