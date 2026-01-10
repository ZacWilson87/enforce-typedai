package types

import "encoding/json"

// Content represents the content of a message.
// Different implementations support text, images, audio, etc.
type Content interface {
	// Type returns the type of content.
	Type() ContentType

	// String returns a string representation of the content.
	String() string
}

// Message represents a single message in a conversation.
type Message struct {
	// Role is the role of the message sender.
	Role Role `json:"role"`

	// Content is the message content. Can be simple text or multi-modal content.
	Content Content `json:"content"`

	// Name is an optional name for the message sender.
	// Useful for distinguishing between multiple users or tools.
	Name string `json:"name,omitempty"`

	// ToolCallID is the ID of the tool call this message is responding to.
	// Required when Role is RoleTool.
	ToolCallID string `json:"tool_call_id,omitempty"`

	// ToolCalls contains tool calls made by the assistant.
	// Only present when Role is RoleAssistant and the model wants to call tools.
	ToolCalls []*ToolCall `json:"tool_calls,omitempty"`

	// FunctionCall contains a function call (legacy, use ToolCalls).
	FunctionCall *FunctionCall `json:"function_call,omitempty"`

	// Metadata contains additional message metadata.
	Metadata *MessageMetadata `json:"metadata,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for Message.
// This handles the Content interface serialization.
func (m *Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	aux := &struct {
		Content interface{} `json:"content"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	// Handle content serialization
	if m.Content != nil {
		switch c := m.Content.(type) {
		case *TextContent:
			aux.Content = c.Text
		case *MultiContent:
			aux.Content = c.Parts
		default:
			aux.Content = m.Content
		}
	}

	return json.Marshal(aux)
}

// UnmarshalJSON implements custom JSON unmarshaling for Message.
func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		Content json.RawMessage `json:"content"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Try to unmarshal content as string first (simple text)
	var text string
	if err := json.Unmarshal(aux.Content, &text); err == nil {
		m.Content = NewTextContent(text)
		return nil
	}

	// Try to unmarshal as array (multi-content)
	var parts []ContentPart
	if err := json.Unmarshal(aux.Content, &parts); err == nil {
		m.Content = &MultiContent{Parts: parts}
		return nil
	}

	return nil
}

// TextContent represents simple text content.
type TextContent struct {
	Text string `json:"text"`
}

// Type returns the content type.
func (t *TextContent) Type() ContentType {
	return ContentTypeText
}

// String returns the text content.
func (t *TextContent) String() string {
	return t.Text
}

// NewTextContent creates a new TextContent.
func NewTextContent(text string) *TextContent {
	return &TextContent{Text: text}
}

// ImageContent represents image content.
type ImageContent struct {
	// URL is the URL of the image.
	URL string `json:"url,omitempty"`

	// Data is the base64-encoded image data.
	Data string `json:"data,omitempty"`

	// Detail is the level of detail for image processing.
	Detail ImageDetail `json:"detail,omitempty"`

	// MimeType is the MIME type of the image (e.g., "image/png").
	MimeType string `json:"mime_type,omitempty"`
}

// Type returns the content type.
func (i *ImageContent) Type() ContentType {
	return ContentTypeImage
}

// String returns a string representation of the image content.
func (i *ImageContent) String() string {
	if i.URL != "" {
		return "[Image: " + i.URL + "]"
	}
	return "[Image: base64 data]"
}

// NewImageContentFromURL creates ImageContent from a URL.
func NewImageContentFromURL(url string, detail ImageDetail) *ImageContent {
	return &ImageContent{
		URL:    url,
		Detail: detail,
	}
}

// NewImageContentFromData creates ImageContent from base64 data.
func NewImageContentFromData(data, mimeType string, detail ImageDetail) *ImageContent {
	return &ImageContent{
		Data:     data,
		MimeType: mimeType,
		Detail:   detail,
	}
}

// AudioContent represents audio content.
type AudioContent struct {
	// URL is the URL of the audio.
	URL string `json:"url,omitempty"`

	// Data is the base64-encoded audio data.
	Data string `json:"data,omitempty"`

	// MimeType is the MIME type of the audio (e.g., "audio/mp3").
	MimeType string `json:"mime_type,omitempty"`

	// Transcript is the text transcript of the audio (if available).
	Transcript string `json:"transcript,omitempty"`
}

// Type returns the content type.
func (a *AudioContent) Type() ContentType {
	return ContentTypeAudio
}

// String returns a string representation of the audio content.
func (a *AudioContent) String() string {
	if a.Transcript != "" {
		return "[Audio: " + a.Transcript + "]"
	}
	if a.URL != "" {
		return "[Audio: " + a.URL + "]"
	}
	return "[Audio: base64 data]"
}

// ContentPart represents a single part of multi-modal content.
type ContentPart struct {
	// Type is the type of this content part.
	Type ContentType `json:"type"`

	// Text is present when Type is ContentTypeText.
	Text string `json:"text,omitempty"`

	// ImageURL contains image information when Type is ContentTypeImage.
	ImageURL *struct {
		URL    string      `json:"url"`
		Detail ImageDetail `json:"detail,omitempty"`
	} `json:"image_url,omitempty"`

	// Audio contains audio information when Type is ContentTypeAudio.
	Audio *AudioContent `json:"audio,omitempty"`
}

// MultiContent represents content with multiple parts (text, images, etc.).
type MultiContent struct {
	Parts []ContentPart `json:"parts"`
}

// Type returns the content type.
func (m *MultiContent) Type() ContentType {
	return ContentTypeText // Multi-content is still considered text-based
}

// String returns a string representation combining all parts.
func (m *MultiContent) String() string {
	result := ""
	for i, part := range m.Parts {
		if i > 0 {
			result += " "
		}
		switch part.Type {
		case ContentTypeText:
			result += part.Text
		case ContentTypeImage, ContentTypeImageURL:
			if part.ImageURL != nil {
				result += "[Image: " + part.ImageURL.URL + "]"
			} else {
				result += "[Image]"
			}
		case ContentTypeAudio:
			if part.Audio != nil {
				result += part.Audio.String()
			} else {
				result += "[Audio]"
			}
		}
	}
	return result
}

// NewMultiContent creates a new MultiContent from parts.
func NewMultiContent(parts ...ContentPart) *MultiContent {
	return &MultiContent{Parts: parts}
}

// NewTextPart creates a text content part.
func NewTextPart(text string) ContentPart {
	return ContentPart{
		Type: ContentTypeText,
		Text: text,
	}
}

// NewImagePart creates an image content part from URL.
func NewImagePart(url string, detail ImageDetail) ContentPart {
	return ContentPart{
		Type: ContentTypeImageURL,
		ImageURL: &struct {
			URL    string      `json:"url"`
			Detail ImageDetail `json:"detail,omitempty"`
		}{
			URL:    url,
			Detail: detail,
		},
	}
}

// NewAudioPart creates an audio content part.
func NewAudioPart(audio *AudioContent) ContentPart {
	return ContentPart{
		Type:  ContentTypeAudio,
		Audio: audio,
	}
}
