package interfaces

import (
	"context"

	"github.com/zacw/go-ai-types/pkg/types"
)

// Provider represents an AI service provider (e.g., OpenAI, Anthropic, Google).
//
// A Provider acts as a factory for service interfaces and provides metadata
// about the provider's capabilities. Implementations of this interface handle
// provider-specific authentication, configuration, and service instantiation.
//
// Example usage:
//
//	provider := openai.NewProvider(apiKey)
//	chatService := provider.ChatService()
//	resp, err := chatService.CreateCompletion(ctx, req)
type Provider interface {
	// Name returns the provider's name (e.g., "openai", "anthropic").
	Name() types.Provider

	// Capabilities returns the list of capabilities supported by this provider.
	// This allows clients to check for feature support before making requests.
	//
	// Example:
	//   caps := provider.Capabilities()
	//   if slices.Contains(caps, types.CapabilityStreaming) {
	//       // Provider supports streaming
	//   }
	Capabilities() []types.ModelCapability

	// Models returns the list of model IDs available from this provider.
	// The returned model IDs can be used in ChatRequest.Model or EmbeddingRequest.Model.
	//
	// Note: For large model catalogs, implementations may return a subset
	// or require separate API calls to fetch the full list.
	Models() []string

	// ChatService returns the chat completion service for this provider.
	// Returns nil if the provider does not support chat completions.
	ChatService() ChatService

	// EmbeddingService returns the embedding generation service for this provider.
	// Returns nil if the provider does not support embeddings.
	EmbeddingService() EmbeddingService
}

// ProviderConfig contains configuration options for initializing a provider.
//
// This struct is designed to be embedded in provider-specific configuration
// structs, allowing each provider to add its own custom fields while maintaining
// a common set of base configuration options.
type ProviderConfig struct {
	// APIKey is the authentication key for the provider.
	APIKey string

	// BaseURL is the base URL for API requests.
	// If empty, the provider's default URL is used.
	BaseURL string

	// Organization is the organization ID for the provider (if applicable).
	Organization string

	// DefaultModel is the default model to use if not specified in requests.
	DefaultModel string

	// Timeout is the default timeout for API requests.
	// If zero, a provider-specific default is used.
	Timeout int

	// MaxRetries is the maximum number of retries for failed requests.
	// If zero, retries are disabled.
	MaxRetries int

	// UserAgent is the User-Agent header to send with requests.
	// If empty, a default user agent is used.
	UserAgent string

	// Custom contains provider-specific configuration options.
	// Implementations can use this for any additional settings.
	Custom map[string]interface{}
}

// ProviderInfo contains metadata about a provider.
//
// This struct is returned by provider discovery and registry functions
// to provide information about available providers without instantiating them.
type ProviderInfo struct {
	// Name is the provider's identifier.
	Name types.Provider

	// DisplayName is the human-readable name of the provider.
	DisplayName string

	// Description provides a brief description of the provider.
	Description string

	// Capabilities lists the capabilities supported by this provider.
	Capabilities []types.ModelCapability

	// DefaultBaseURL is the default base URL for API requests.
	DefaultBaseURL string

	// RequiresAPIKey indicates whether this provider requires an API key.
	RequiresAPIKey bool

	// SupportsOrganization indicates whether this provider supports organization IDs.
	SupportsOrganization bool

	// DocumentationURL is the URL to the provider's API documentation.
	DocumentationURL string

	// Models is a list of popular or recommended models for this provider.
	// For the full model list, use Provider.Models().
	Models []string
}

// HealthChecker is an optional interface that providers can implement
// to support health checks.
//
// Health checks are useful for monitoring and ensuring that the provider
// is accessible and functioning correctly before making actual requests.
type HealthChecker interface {
	// Health checks the provider's health by making a lightweight API call.
	// Returns nil if the provider is healthy, or an error describing the problem.
	//
	// Implementations should use a simple, low-cost operation like listing models
	// or validating credentials.
	Health(ctx context.Context) error
}

// ModelLister is an optional interface that providers can implement
// to support dynamic model discovery.
//
// This is useful for providers that have frequently changing model catalogs
// or for fetching detailed model information beyond just model IDs.
type ModelLister interface {
	// ListModels returns detailed information about all available models.
	// This may make an API call to fetch the latest model information.
	ListModels(ctx context.Context) ([]*types.ModelInfo, error)

	// GetModel returns detailed information about a specific model.
	// Returns an error if the model doesn't exist or can't be retrieved.
	GetModel(ctx context.Context, modelID string) (*types.ModelInfo, error)
}

// ProviderFactory is a function that creates a new Provider instance.
// This type is used by provider registries to enable dynamic provider creation.
type ProviderFactory func(config *ProviderConfig) (Provider, error)
