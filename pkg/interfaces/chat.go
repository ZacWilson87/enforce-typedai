package interfaces

import (
	"context"

	"github.com/zacw/go-ai-types/pkg/types"
)

// ChatService provides methods for creating chat completions.
//
// This interface abstracts the chat completion API across different providers,
// allowing applications to work with multiple AI providers using a unified API.
//
// Implementations must handle provider-specific request/response conversion,
// authentication, and error handling internally.
//
// Example usage:
//
//	chatService := provider.ChatService()
//	req := &types.ChatRequest{
//	    Model: "gpt-4",
//	    Messages: []*types.Message{
//	        {Role: types.RoleUser, Content: types.NewTextContent("Hello!")},
//	    },
//	}
//	resp, err := chatService.CreateCompletion(ctx, req)
type ChatService interface {
	// CreateCompletion generates a chat completion for the given request.
	//
	// This method makes a synchronous API call and returns the complete response.
	// For streaming responses, use CreateCompletionStream instead.
	//
	// The context can be used to cancel the request or set timeouts:
	//   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//   defer cancel()
	//   resp, err := service.CreateCompletion(ctx, req)
	//
	// Returns an error if:
	// - The request is invalid (validation error)
	// - Authentication fails
	// - The API request fails (network, server error, etc.)
	// - The context is cancelled or times out
	CreateCompletion(ctx context.Context, req *types.ChatRequest) (*types.ChatResponse, error)

	// CreateCompletionStream generates a streaming chat completion for the given request.
	//
	// This method returns a channel that yields stream chunks as they arrive.
	// The channel is closed when the stream completes or encounters an error.
	//
	// Implementations must set req.Stream = true before making the API call.
	//
	// Example:
	//   stream, err := service.CreateCompletionStream(ctx, req)
	//   if err != nil {
	//       return err
	//   }
	//
	//   for chunk := range stream {
	//       if chunk.IsComplete() {
	//           break
	//       }
	//       // Process chunk
	//       fmt.Print(chunk.GetChoices()[0].Delta.Content)
	//   }
	//
	// The context can be used to cancel the stream:
	//   ctx, cancel := context.WithCancel(context.Background())
	//   defer cancel()
	//   // Call cancel() to stop the stream
	//
	// Returns an error if:
	// - The request is invalid
	// - Authentication fails
	// - The stream cannot be established
	//
	// Note: Errors that occur during streaming are sent as error chunks
	// rather than being returned from this method.
	CreateCompletionStream(ctx context.Context, req *types.ChatRequest) (<-chan types.StreamChunk, error)
}

// ChatServiceWithCallback extends ChatService with callback-based streaming.
//
// This interface is useful for providers or use cases where callback-based
// streaming is more natural than channel-based streaming.
type ChatServiceWithCallback interface {
	ChatService

	// CreateCompletionStreamWithCallback generates a streaming chat completion
	// and calls the provided handler for each chunk.
	//
	// The handler is called for each chunk as it arrives. If the handler returns
	// an error, streaming stops and that error is returned.
	//
	// Example:
	//   handler := &MyStreamHandler{}
	//   err := service.CreateCompletionStreamWithCallback(ctx, req, handler)
	//
	// Returns an error if:
	// - The request is invalid
	// - Authentication fails
	// - The stream cannot be established
	// - The handler returns an error
	CreateCompletionStreamWithCallback(ctx context.Context, req *types.ChatRequest, handler StreamHandler) error
}

// ChatServiceWithValidation extends ChatService with request validation.
//
// This interface allows clients to validate requests before making API calls,
// which can help catch errors early and provide better error messages.
type ChatServiceWithValidation interface {
	ChatService

	// ValidateRequest validates a chat request without making an API call.
	//
	// This can be used to check for common errors like:
	// - Empty messages array
	// - Invalid role values
	// - Missing required fields
	// - Invalid parameter ranges
	//
	// Returns nil if the request is valid, or a ValidationError describing
	// the problem.
	ValidateRequest(req *types.ChatRequest) error
}

// ChatServiceWithMetrics extends ChatService with metrics collection.
//
// This interface is useful for monitoring and observability, allowing
// implementations to track request counts, latencies, token usage, etc.
type ChatServiceWithMetrics interface {
	ChatService

	// GetMetrics returns metrics collected by this service.
	//
	// The returned map contains provider-specific metrics. Common metrics include:
	// - "requests_total": Total number of requests made
	// - "requests_failed": Number of failed requests
	// - "tokens_total": Total tokens used
	// - "latency_ms": Average latency in milliseconds
	//
	// Implementations should document their available metrics.
	GetMetrics() map[string]interface{}

	// ResetMetrics resets all collected metrics to zero.
	// This is useful for testing or periodic metric collection.
	ResetMetrics()
}

// Handler is a function that processes a chat request and returns a response.
//
// This type is used by middleware to wrap and compose request processing logic.
// See the Middleware interface for more details.
type Handler func(ctx context.Context, req *types.ChatRequest) (*types.ChatResponse, error)

// StreamingHandler is a function that processes a streaming chat request.
//
// This type is used by middleware to wrap and compose streaming request processing.
type StreamingHandler func(ctx context.Context, req *types.ChatRequest) (<-chan types.StreamChunk, error)
