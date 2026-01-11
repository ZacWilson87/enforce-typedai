package interfaces

import (
	"time"

	"github.com/zacw/go-ai-types/pkg/types"
)

// Middleware provides a way to wrap and compose request processing logic.
//
// Middleware enables cross-cutting concerns like logging, rate limiting,
// retries, caching, and metrics to be implemented in a modular, composable way.
//
// Middleware follows the decorator pattern, wrapping a Handler with additional
// behavior and returning a new Handler.
//
// Example middleware implementation:
//
//	type LoggingMiddleware struct {
//	    logger *log.Logger
//	}
//
//	func (m *LoggingMiddleware) Wrap(next Handler) Handler {
//	    return func(ctx context.Context, req *types.ChatRequest) (*types.ChatResponse, error) {
//	        m.logger.Printf("Request: model=%s messages=%d", req.Model, len(req.Messages))
//	        resp, err := next(ctx, req)
//	        if err != nil {
//	            m.logger.Printf("Error: %v", err)
//	        } else {
//	            m.logger.Printf("Response: tokens=%d", resp.Usage.TotalTokens)
//	        }
//	        return resp, err
//	    }
//	}
//
// Example usage:
//
//	service := provider.ChatService()
//
//	// Wrap with middleware
//	loggingMiddleware := &LoggingMiddleware{logger: log.Default()}
//	retryMiddleware := &RetryMiddleware{maxRetries: 3}
//
//	// Create wrapped handler
//	handler := loggingMiddleware.Wrap(retryMiddleware.Wrap(service.CreateCompletion))
//
//	// Use the wrapped handler
//	resp, err := handler(ctx, req)
type Middleware interface {
	// Wrap wraps a handler with additional behavior.
	//
	// Implementations should call next(ctx, req) to invoke the next handler
	// in the chain, and can perform actions before and/or after this call.
	//
	// Middleware should preserve the context and request, only modifying them
	// if that is the explicit purpose of the middleware (e.g., timeout middleware
	// may wrap the context with a timeout).
	Wrap(next Handler) Handler
}

// StreamingMiddleware provides a way to wrap streaming request processing.
//
// This is the streaming equivalent of Middleware, used to wrap StreamingHandler
// functions with additional behavior.
type StreamingMiddleware interface {
	// WrapStream wraps a streaming handler with additional behavior.
	//
	// Implementations can intercept the stream channel returned by next(ctx, req)
	// and wrap it with additional processing, such as logging chunks, applying
	// rate limits, or collecting metrics.
	WrapStream(next StreamingHandler) StreamingHandler
}

// LoggingConfig configures logging middleware behavior.
type LoggingConfig struct {
	// LogRequests enables logging of outgoing requests.
	LogRequests bool

	// LogResponses enables logging of incoming responses.
	LogResponses bool

	// LogErrors enables logging of errors.
	LogErrors bool

	// LogTokenUsage enables logging of token usage.
	LogTokenUsage bool

	// IncludeTimestamps includes timestamps in log messages.
	IncludeTimestamps bool

	// IncludeRequestID includes request IDs in log messages.
	IncludeRequestID bool
}

// RateLimitConfig configures rate limiting middleware behavior.
type RateLimitConfig struct {
	// RequestsPerSecond is the maximum number of requests per second.
	// If zero, no rate limiting is applied.
	RequestsPerSecond float64

	// RequestsPerMinute is the maximum number of requests per minute.
	// If zero, no rate limiting is applied.
	RequestsPerMinute int

	// TokensPerMinute is the maximum number of tokens per minute.
	// If zero, no token-based rate limiting is applied.
	TokensPerMinute int

	// Burst is the maximum burst size for the rate limiter.
	// This allows short bursts of requests above the sustained rate.
	Burst int

	// WaitTimeout is the maximum time to wait for rate limit availability.
	// If zero, requests fail immediately when rate limited.
	WaitTimeout time.Duration
}

// RetryConfig configures retry middleware behavior.
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts.
	// If zero, no retries are performed.
	MaxRetries int

	// InitialBackoff is the initial backoff duration.
	// Defaults to 1 second if zero.
	InitialBackoff time.Duration

	// MaxBackoff is the maximum backoff duration.
	// Defaults to 60 seconds if zero.
	MaxBackoff time.Duration

	// BackoffMultiplier is the multiplier for exponential backoff.
	// Defaults to 2.0 if zero.
	BackoffMultiplier float64

	// RetryableErrors is a list of error types that should trigger retries.
	// If nil, a default set of retryable errors is used (rate limit, timeout, server errors).
	RetryableErrors []types.ErrorType

	// ShouldRetry is a custom function to determine if an error should be retried.
	// If provided, this takes precedence over RetryableErrors.
	ShouldRetry func(error) bool

	// OnRetry is called before each retry attempt.
	// It receives the attempt number (starting at 1) and the error that triggered the retry.
	OnRetry func(attempt int, err error)
}

// CacheConfig configures caching middleware behavior.
type CacheConfig struct {
	// TTL is the time-to-live for cached responses.
	// If zero, cached responses never expire.
	TTL time.Duration

	// MaxSize is the maximum number of cached responses.
	// If zero, no size limit is applied.
	MaxSize int

	// MaxMemoryBytes is the maximum memory usage for the cache in bytes.
	// If zero, no memory limit is applied.
	MaxMemoryBytes int64

	// KeyFunc generates cache keys from requests.
	// If nil, a default key function is used based on model and messages.
	KeyFunc func(*types.ChatRequest) string

	// ShouldCache determines whether a response should be cached.
	// If nil, all successful responses are cached.
	ShouldCache func(*types.ChatRequest, *types.ChatResponse) bool
}

// MetricsCollector defines an interface for collecting metrics.
//
// Middleware implementations can use this interface to report metrics
// to various backends (Prometheus, StatsD, CloudWatch, etc.).
type MetricsCollector interface {
	// RecordRequest records a request being made.
	RecordRequest(provider types.Provider, model string)

	// RecordResponse records a successful response.
	RecordResponse(provider types.Provider, model string, duration time.Duration, tokens int)

	// RecordError records an error.
	RecordError(provider types.Provider, model string, errorType types.ErrorType)

	// RecordTokenUsage records token usage.
	RecordTokenUsage(provider types.Provider, model string, promptTokens, completionTokens int)

	// RecordCacheHit records a cache hit.
	RecordCacheHit(provider types.Provider, model string)

	// RecordCacheMiss records a cache miss.
	RecordCacheMiss(provider types.Provider, model string)

	// RecordRetry records a retry attempt.
	RecordRetry(provider types.Provider, model string, attempt int)
}

// TimeoutConfig configures timeout middleware behavior.
type TimeoutConfig struct {
	// RequestTimeout is the timeout for individual requests.
	// If zero, no timeout is applied at the middleware level
	// (the underlying service may still have its own timeout).
	RequestTimeout time.Duration

	// StreamChunkTimeout is the timeout for receiving individual stream chunks.
	// If a chunk is not received within this duration, the stream is terminated.
	// If zero, no per-chunk timeout is applied.
	StreamChunkTimeout time.Duration
}

// CircuitBreakerConfig configures circuit breaker middleware behavior.
//
// A circuit breaker prevents cascading failures by temporarily blocking
// requests when a service is experiencing high error rates.
type CircuitBreakerConfig struct {
	// MaxFailures is the number of consecutive failures before opening the circuit.
	MaxFailures int

	// Timeout is how long to wait in the open state before attempting recovery.
	Timeout time.Duration

	// HalfOpenMaxRequests is the maximum number of requests allowed in half-open state.
	// These requests are used to test if the service has recovered.
	HalfOpenMaxRequests int

	// ShouldTrip is a custom function to determine if the circuit should open.
	// If provided, this takes precedence over MaxFailures.
	ShouldTrip func(counts CircuitBreakerCounts) bool

	// OnStateChange is called when the circuit breaker changes state.
	OnStateChange func(from, to CircuitBreakerState)
}

// CircuitBreakerState represents the state of a circuit breaker.
type CircuitBreakerState int

const (
	// CircuitBreakerClosed means the circuit is closed and requests flow normally.
	CircuitBreakerClosed CircuitBreakerState = iota

	// CircuitBreakerOpen means the circuit is open and requests are blocked.
	CircuitBreakerOpen

	// CircuitBreakerHalfOpen means the circuit is testing if the service has recovered.
	CircuitBreakerHalfOpen
)

// String returns the string representation of the circuit breaker state.
func (s CircuitBreakerState) String() string {
	switch s {
	case CircuitBreakerClosed:
		return "closed"
	case CircuitBreakerOpen:
		return "open"
	case CircuitBreakerHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// CircuitBreakerCounts tracks circuit breaker statistics.
type CircuitBreakerCounts struct {
	// Requests is the total number of requests.
	Requests uint32

	// TotalSuccesses is the total number of successful requests.
	TotalSuccesses uint32

	// TotalFailures is the total number of failed requests.
	TotalFailures uint32

	// ConsecutiveSuccesses is the number of consecutive successful requests.
	ConsecutiveSuccesses uint32

	// ConsecutiveFailures is the number of consecutive failed requests.
	ConsecutiveFailures uint32
}
