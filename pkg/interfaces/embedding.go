package interfaces

import (
	"context"

	"github.com/zacw/go-ai-types/pkg/types"
)

// EmbeddingService provides methods for generating embeddings.
//
// Embeddings are vector representations of text that capture semantic meaning,
// useful for tasks like semantic search, clustering, and similarity comparison.
//
// This interface abstracts embedding generation across different providers,
// allowing applications to work with multiple providers using a unified API.
//
// Example usage:
//
//	embeddingService := provider.EmbeddingService()
//	req := &types.EmbeddingRequest{
//	    Model: "text-embedding-ada-002",
//	    Input: "Hello, world!",
//	}
//	resp, err := embeddingService.CreateEmbedding(ctx, req)
type EmbeddingService interface {
	// CreateEmbedding generates embeddings for the given input.
	//
	// The input can be a single string or an array of strings:
	//   req.Input = "Hello, world!"          // Single input
	//   req.Input = []string{"Hello", "Hi"}  // Multiple inputs
	//
	// The response contains an array of embeddings, one for each input.
	// The embeddings are returned in the same order as the inputs.
	//
	// Example:
	//   req := &types.EmbeddingRequest{
	//       Model: "text-embedding-ada-002",
	//       Input: []string{"cat", "dog", "car"},
	//   }
	//   resp, err := service.CreateEmbedding(ctx, req)
	//   if err != nil {
	//       return err
	//   }
	//
	//   for _, emb := range resp.Data {
	//       vector := emb.AsFloatVector()
	//       // Use vector for similarity search, etc.
	//   }
	//
	// The context can be used to cancel the request or set timeouts:
	//   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//   defer cancel()
	//
	// Returns an error if:
	// - The request is invalid (validation error)
	// - Authentication fails
	// - The API request fails (network, server error, etc.)
	// - The context is cancelled or times out
	// - The input is too long for the model
	CreateEmbedding(ctx context.Context, req *types.EmbeddingRequest) (*types.EmbeddingResponse, error)
}

// EmbeddingServiceWithValidation extends EmbeddingService with request validation.
//
// This interface allows clients to validate requests before making API calls,
// which can help catch errors early and provide better error messages.
type EmbeddingServiceWithValidation interface {
	EmbeddingService

	// ValidateRequest validates an embedding request without making an API call.
	//
	// This can be used to check for common errors like:
	// - Empty or nil input
	// - Input too long for the model
	// - Invalid encoding format
	// - Invalid dimensions parameter
	//
	// Returns nil if the request is valid, or a ValidationError describing
	// the problem.
	ValidateRequest(req *types.EmbeddingRequest) error
}

// EmbeddingServiceWithMetrics extends EmbeddingService with metrics collection.
//
// This interface is useful for monitoring and observability, allowing
// implementations to track request counts, latencies, token usage, etc.
type EmbeddingServiceWithMetrics interface {
	EmbeddingService

	// GetMetrics returns metrics collected by this service.
	//
	// The returned map contains provider-specific metrics. Common metrics include:
	// - "requests_total": Total number of requests made
	// - "requests_failed": Number of failed requests
	// - "tokens_total": Total tokens processed
	// - "embeddings_total": Total embeddings generated
	// - "latency_ms": Average latency in milliseconds
	//
	// Implementations should document their available metrics.
	GetMetrics() map[string]interface{}

	// ResetMetrics resets all collected metrics to zero.
	// This is useful for testing or periodic metric collection.
	ResetMetrics()
}

// EmbeddingServiceWithBatch extends EmbeddingService with batch processing.
//
// This interface is useful for efficiently processing large numbers of inputs
// by automatically splitting them into optimal batch sizes for the provider.
type EmbeddingServiceWithBatch interface {
	EmbeddingService

	// CreateEmbeddingBatch generates embeddings for a large batch of inputs.
	//
	// This method automatically splits the inputs into optimal batch sizes
	// for the provider and makes multiple API calls in parallel if necessary.
	//
	// Example:
	//   inputs := []string{...} // 1000 inputs
	//   resp, err := service.CreateEmbeddingBatch(ctx, model, inputs, nil)
	//
	// The response contains all embeddings in the same order as the inputs,
	// regardless of how they were batched internally.
	//
	// Options can be provided to control batching behavior:
	// - "batch_size": Maximum number of inputs per batch
	// - "parallel": Number of parallel API calls
	//
	// Returns an error if any batch fails. Partial results are not returned.
	CreateEmbeddingBatch(ctx context.Context, model string, inputs []string, options map[string]interface{}) (*types.EmbeddingResponse, error)
}

// EmbeddingServiceWithCache extends EmbeddingService with caching.
//
// This interface allows implementations to cache embeddings to avoid
// redundant API calls for the same inputs.
type EmbeddingServiceWithCache interface {
	EmbeddingService

	// CreateEmbeddingWithCache generates embeddings, using cached results when available.
	//
	// The cache key is based on the model and input text. If a cached embedding
	// exists for the input, it is returned immediately without making an API call.
	//
	// This is particularly useful when:
	// - Processing the same inputs multiple times
	// - Building indexes that may be rebuilt
	// - Development and testing with fixed inputs
	//
	// Example:
	//   resp, cached, err := service.CreateEmbeddingWithCache(ctx, req)
	//   if cached {
	//       log.Println("Used cached embedding")
	//   }
	//
	// Returns:
	// - response: The embedding response (from cache or API)
	// - cached: true if the result came from cache, false if it required an API call
	// - error: any error that occurred
	CreateEmbeddingWithCache(ctx context.Context, req *types.EmbeddingRequest) (*types.EmbeddingResponse, bool, error)

	// ClearCache clears all cached embeddings.
	ClearCache()

	// GetCacheStats returns statistics about cache usage.
	//
	// The returned map may include:
	// - "hits": Number of cache hits
	// - "misses": Number of cache misses
	// - "size": Number of cached embeddings
	// - "memory_bytes": Approximate memory used by cache
	GetCacheStats() map[string]interface{}
}

// EmbeddingHandler is a function that processes an embedding request and returns a response.
//
// This type is used by middleware to wrap and compose request processing logic.
// See the Middleware interface for more details.
type EmbeddingHandler func(ctx context.Context, req *types.EmbeddingRequest) (*types.EmbeddingResponse, error)
