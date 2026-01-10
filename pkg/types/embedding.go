package types

// EmbeddingRequest represents a request to generate embeddings.
type EmbeddingRequest struct {
	// Model is the ID of the model to use for embeddings.
	Model string `json:"model"`

	// Input is the text or texts to generate embeddings for.
	// Can be a string or array of strings.
	Input interface{} `json:"input"`

	// EncodingFormat specifies the format for the embeddings.
	// Supported values: "float", "base64".
	EncodingFormat string `json:"encoding_format,omitempty"`

	// Dimensions is the number of dimensions for the embedding (if supported).
	// Some models support reducing the embedding dimensionality.
	Dimensions int `json:"dimensions,omitempty"`

	// User is a unique identifier for the end-user.
	User string `json:"user,omitempty"`

	// Metadata contains additional request metadata.
	Metadata *RequestMetadata `json:"metadata,omitempty"`
}

// EmbeddingResponse represents a response from embeddings generation.
type EmbeddingResponse struct {
	// Object is the object type (e.g., "list").
	Object string `json:"object"`

	// Data contains the list of embeddings.
	Data []*Embedding `json:"data"`

	// Model is the model used to generate the embeddings.
	Model string `json:"model"`

	// Usage contains token usage information.
	Usage *Usage `json:"usage,omitempty"`

	// Metadata contains additional response metadata.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
}

// Embedding represents a single embedding vector.
type Embedding struct {
	// Object is the object type (e.g., "embedding").
	Object string `json:"object"`

	// Index is the index of this embedding in the list.
	Index int `json:"index"`

	// Embedding is the embedding vector.
	// Usually a slice of float64, but can be base64 encoded.
	Embedding interface{} `json:"embedding"`

	// Dimensions is the number of dimensions in the embedding.
	Dimensions int `json:"dimensions,omitempty"`
}

// AsFloatVector returns the embedding as a float64 slice.
// Returns nil if the embedding is not in float format.
func (e *Embedding) AsFloatVector() []float64 {
	switch v := e.Embedding.(type) {
	case []float64:
		return v
	case []interface{}:
		// Convert []interface{} to []float64
		result := make([]float64, len(v))
		for i, val := range v {
			if f, ok := val.(float64); ok {
				result[i] = f
			}
		}
		return result
	default:
		return nil
	}
}

// AsBase64 returns the embedding as a base64 encoded string.
// Returns empty string if the embedding is not in base64 format.
func (e *Embedding) AsBase64() string {
	if s, ok := e.Embedding.(string); ok {
		return s
	}
	return ""
}

// Helper functions for creating requests

// NewEmbeddingRequest creates a new EmbeddingRequest.
func NewEmbeddingRequest(model string, input interface{}) *EmbeddingRequest {
	return &EmbeddingRequest{
		Model: model,
		Input: input,
	}
}

// NewEmbeddingRequestFromString creates an embedding request for a single string.
func NewEmbeddingRequestFromString(model, input string) *EmbeddingRequest {
	return &EmbeddingRequest{
		Model: model,
		Input: input,
	}
}

// NewEmbeddingRequestFromStrings creates an embedding request for multiple strings.
func NewEmbeddingRequestFromStrings(model string, inputs []string) *EmbeddingRequest {
	return &EmbeddingRequest{
		Model: model,
		Input: inputs,
	}
}

// WithEncodingFormat sets the encoding format for the request.
func (r *EmbeddingRequest) WithEncodingFormat(format string) *EmbeddingRequest {
	r.EncodingFormat = format
	return r
}

// WithDimensions sets the dimensions for the request.
func (r *EmbeddingRequest) WithDimensions(dimensions int) *EmbeddingRequest {
	r.Dimensions = dimensions
	return r
}

// WithUser sets the user ID for the request.
func (r *EmbeddingRequest) WithUser(user string) *EmbeddingRequest {
	r.User = user
	return r
}

// GetInputAsString returns the input as a single string.
// Returns empty string if input is not a string.
func (r *EmbeddingRequest) GetInputAsString() string {
	if s, ok := r.Input.(string); ok {
		return s
	}
	return ""
}

// GetInputAsStrings returns the input as a slice of strings.
// Returns nil if input cannot be converted to string slice.
func (r *EmbeddingRequest) GetInputAsStrings() []string {
	switch v := r.Input.(type) {
	case string:
		return []string{v}
	case []string:
		return v
	case []interface{}:
		result := make([]string, 0, len(v))
		for _, item := range v {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	default:
		return nil
	}
}

// GetFirstEmbedding returns the first embedding from the response.
func (r *EmbeddingResponse) GetFirstEmbedding() *Embedding {
	if len(r.Data) > 0 {
		return r.Data[0]
	}
	return nil
}

// GetFirstVector returns the first embedding as a float vector.
func (r *EmbeddingResponse) GetFirstVector() []float64 {
	if emb := r.GetFirstEmbedding(); emb != nil {
		return emb.AsFloatVector()
	}
	return nil
}

// GetAllVectors returns all embeddings as float vectors.
func (r *EmbeddingResponse) GetAllVectors() [][]float64 {
	vectors := make([][]float64, 0, len(r.Data))
	for _, emb := range r.Data {
		if vec := emb.AsFloatVector(); vec != nil {
			vectors = append(vectors, vec)
		}
	}
	return vectors
}
