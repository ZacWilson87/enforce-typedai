package types

// Role represents the role of a message sender in a conversation.
type Role string

const (
	// RoleUser represents a message from the user/human.
	RoleUser Role = "user"

	// RoleAssistant represents a message from the AI assistant.
	RoleAssistant Role = "assistant"

	// RoleSystem represents a system message that sets context or instructions.
	RoleSystem Role = "system"

	// RoleTool represents a message from a tool/function execution.
	RoleTool Role = "tool"

	// RoleFunction represents a message from a function call (legacy, use RoleTool).
	RoleFunction Role = "function"
)

// String returns the string representation of the Role.
func (r Role) String() string {
	return string(r)
}

// IsValid returns true if the Role is one of the defined constants.
func (r Role) IsValid() bool {
	switch r {
	case RoleUser, RoleAssistant, RoleSystem, RoleTool, RoleFunction:
		return true
	default:
		return false
	}
}

// ContentType represents the type of content in a message.
type ContentType string

const (
	// ContentTypeText represents plain text content.
	ContentTypeText ContentType = "text"

	// ContentTypeImage represents image content (URL or base64).
	ContentTypeImage ContentType = "image"

	// ContentTypeImageURL represents an image provided via URL.
	ContentTypeImageURL ContentType = "image_url"

	// ContentTypeAudio represents audio content.
	ContentTypeAudio ContentType = "audio"

	// ContentTypeVideo represents video content.
	ContentTypeVideo ContentType = "video"

	// ContentTypeFile represents file content.
	ContentTypeFile ContentType = "file"
)

// String returns the string representation of the ContentType.
func (c ContentType) String() string {
	return string(c)
}

// IsValid returns true if the ContentType is one of the defined constants.
func (c ContentType) IsValid() bool {
	switch c {
	case ContentTypeText, ContentTypeImage, ContentTypeImageURL,
		ContentTypeAudio, ContentTypeVideo, ContentTypeFile:
		return true
	default:
		return false
	}
}

// FinishReason represents why a model stopped generating tokens.
type FinishReason string

const (
	// FinishReasonStop indicates natural completion.
	FinishReasonStop FinishReason = "stop"

	// FinishReasonLength indicates max token limit reached.
	FinishReasonLength FinishReason = "length"

	// FinishReasonToolCalls indicates model wants to call a tool/function.
	FinishReasonToolCalls FinishReason = "tool_calls"

	// FinishReasonFunctionCall indicates model wants to call a function (legacy).
	FinishReasonFunctionCall FinishReason = "function_call"

	// FinishReasonContentFilter indicates content was filtered.
	FinishReasonContentFilter FinishReason = "content_filter"

	// FinishReasonError indicates an error occurred.
	FinishReasonError FinishReason = "error"

	// FinishReasonNull indicates no finish reason (streaming in progress).
	FinishReasonNull FinishReason = "null"
)

// String returns the string representation of the FinishReason.
func (f FinishReason) String() string {
	return string(f)
}

// IsValid returns true if the FinishReason is one of the defined constants.
func (f FinishReason) IsValid() bool {
	switch f {
	case FinishReasonStop, FinishReasonLength, FinishReasonToolCalls,
		FinishReasonFunctionCall, FinishReasonContentFilter,
		FinishReasonError, FinishReasonNull:
		return true
	default:
		return false
	}
}

// ModelCapability represents capabilities that a model can support.
type ModelCapability string

const (
	// CapabilityChat indicates the model supports chat completions.
	CapabilityChat ModelCapability = "chat"

	// CapabilityCompletion indicates the model supports text completions.
	CapabilityCompletion ModelCapability = "completion"

	// CapabilityEmbedding indicates the model supports embeddings generation.
	CapabilityEmbedding ModelCapability = "embedding"

	// CapabilityStreaming indicates the model supports streaming responses.
	CapabilityStreaming ModelCapability = "streaming"

	// CapabilityFunctionCalling indicates the model supports function calling.
	CapabilityFunctionCalling ModelCapability = "function_calling"

	// CapabilityToolCalling indicates the model supports tool calling.
	CapabilityToolCalling ModelCapability = "tool_calling"

	// CapabilityVision indicates the model supports image understanding.
	CapabilityVision ModelCapability = "vision"

	// CapabilityAudio indicates the model supports audio processing.
	CapabilityAudio ModelCapability = "audio"

	// CapabilityImageGeneration indicates the model can generate images.
	CapabilityImageGeneration ModelCapability = "image_generation"

	// CapabilityJSONMode indicates the model supports JSON output mode.
	CapabilityJSONMode ModelCapability = "json_mode"
)

// String returns the string representation of the ModelCapability.
func (m ModelCapability) String() string {
	return string(m)
}

// Provider represents an AI service provider.
type Provider string

const (
	// ProviderOpenAI represents OpenAI (GPT models).
	ProviderOpenAI Provider = "openai"

	// ProviderAnthropic represents Anthropic (Claude models).
	ProviderAnthropic Provider = "anthropic"

	// ProviderGoogle represents Google (Gemini, PaLM models).
	ProviderGoogle Provider = "google"

	// ProviderCohere represents Cohere.
	ProviderCohere Provider = "cohere"

	// ProviderMistral represents Mistral AI.
	ProviderMistral Provider = "mistral"

	// ProviderHuggingFace represents Hugging Face.
	ProviderHuggingFace Provider = "huggingface"

	// ProviderAzure represents Azure OpenAI.
	ProviderAzure Provider = "azure"

	// ProviderCustom represents a custom provider.
	ProviderCustom Provider = "custom"
)

// String returns the string representation of the Provider.
func (p Provider) String() string {
	return string(p)
}

// IsValid returns true if the Provider is one of the defined constants.
func (p Provider) IsValid() bool {
	switch p {
	case ProviderOpenAI, ProviderAnthropic, ProviderGoogle,
		ProviderCohere, ProviderMistral, ProviderHuggingFace,
		ProviderAzure, ProviderCustom:
		return true
	default:
		return false
	}
}

// ImageDetail represents the level of detail for image processing.
type ImageDetail string

const (
	// ImageDetailAuto lets the model choose the detail level.
	ImageDetailAuto ImageDetail = "auto"

	// ImageDetailLow uses lower resolution for faster processing.
	ImageDetailLow ImageDetail = "low"

	// ImageDetailHigh uses higher resolution for more detailed analysis.
	ImageDetailHigh ImageDetail = "high"
)

// String returns the string representation of the ImageDetail.
func (i ImageDetail) String() string {
	return string(i)
}
