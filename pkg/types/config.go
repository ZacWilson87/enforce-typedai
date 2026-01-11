package types

import "time"

// ClientConfig contains configuration for AI service clients.
//
// This struct provides common configuration options that can be used
// across different provider implementations.
type ClientConfig struct {
	// APIKey is the authentication key for the provider.
	APIKey string `json:"api_key,omitempty"`

	// BaseURL is the base URL for API requests.
	// If empty, the provider's default URL is used.
	BaseURL string `json:"base_url,omitempty"`

	// Organization is the organization ID for the provider (if applicable).
	Organization string `json:"organization,omitempty"`

	// DefaultModel is the default model to use if not specified in requests.
	DefaultModel string `json:"default_model,omitempty"`

	// Timeout is the default timeout for API requests.
	// If zero, a provider-specific default is used (typically 60 seconds).
	Timeout time.Duration `json:"timeout,omitempty"`

	// HTTPConfig contains HTTP-specific configuration.
	HTTPConfig *HTTPConfig `json:"http_config,omitempty"`

	// RetryConfig contains retry configuration.
	RetryConfig *RetryConfig `json:"retry_config,omitempty"`

	// UserAgent is the User-Agent header to send with requests.
	// If empty, a default user agent is used.
	UserAgent string `json:"user_agent,omitempty"`

	// Headers contains additional HTTP headers to send with all requests.
	Headers map[string]string `json:"headers,omitempty"`

	// Debug enables debug logging.
	Debug bool `json:"debug,omitempty"`
}

// HTTPConfig contains HTTP client configuration.
type HTTPConfig struct {
	// MaxIdleConns controls the maximum number of idle connections.
	// Default is 100.
	MaxIdleConns int `json:"max_idle_conns,omitempty"`

	// MaxIdleConnsPerHost controls the maximum idle connections per host.
	// Default is 10.
	MaxIdleConnsPerHost int `json:"max_idle_conns_per_host,omitempty"`

	// MaxConnsPerHost limits the total number of connections per host.
	// Default is 0 (unlimited).
	MaxConnsPerHost int `json:"max_conns_per_host,omitempty"`

	// IdleConnTimeout is the maximum amount of time an idle connection
	// will remain idle before closing itself.
	// Default is 90 seconds.
	IdleConnTimeout time.Duration `json:"idle_conn_timeout,omitempty"`

	// TLSHandshakeTimeout specifies the maximum amount of time waiting for
	// a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration `json:"tls_handshake_timeout,omitempty"`

	// ExpectContinueTimeout specifies the amount of time to wait for a
	// server's first response headers after fully writing the request headers
	// if the request has an "Expect: 100-continue" header.
	// Zero means no timeout.
	ExpectContinueTimeout time.Duration `json:"expect_continue_timeout,omitempty"`

	// ResponseHeaderTimeout specifies the amount of time to wait for a
	// server's response headers after fully writing the request.
	// Zero means no timeout.
	ResponseHeaderTimeout time.Duration `json:"response_header_timeout,omitempty"`

	// DisableCompression disables compression for requests.
	// Default is false (compression enabled).
	DisableCompression bool `json:"disable_compression,omitempty"`

	// DisableKeepAlives disables HTTP keep-alives.
	// Default is false (keep-alives enabled).
	DisableKeepAlives bool `json:"disable_keep_alives,omitempty"`

	// ProxyURL is the URL of the proxy to use for requests.
	// If empty, no proxy is used.
	ProxyURL string `json:"proxy_url,omitempty"`
}

// RetryConfig contains retry configuration.
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts.
	// If zero, no retries are performed.
	MaxRetries int `json:"max_retries,omitempty"`

	// InitialBackoff is the initial backoff duration.
	// Default is 1 second.
	InitialBackoff time.Duration `json:"initial_backoff,omitempty"`

	// MaxBackoff is the maximum backoff duration.
	// Default is 60 seconds.
	MaxBackoff time.Duration `json:"max_backoff,omitempty"`

	// BackoffMultiplier is the multiplier for exponential backoff.
	// Default is 2.0.
	BackoffMultiplier float64 `json:"backoff_multiplier,omitempty"`

	// RetryableStatusCodes is a list of HTTP status codes that should trigger retries.
	// If nil, a default set is used (429, 500, 502, 503, 504).
	RetryableStatusCodes []int `json:"retryable_status_codes,omitempty"`

	// RetryableErrors is a list of error types that should trigger retries.
	// If nil, a default set is used (rate limit, timeout, server errors).
	RetryableErrors []ErrorType `json:"retryable_errors,omitempty"`
}

// StreamConfig contains streaming configuration.
type StreamConfig struct {
	// BufferSize is the size of the channel buffer for streaming responses.
	// Default is 100.
	BufferSize int `json:"buffer_size,omitempty"`

	// ChunkTimeout is the maximum time to wait for a chunk.
	// If a chunk is not received within this duration, the stream is terminated.
	// Zero means no timeout.
	ChunkTimeout time.Duration `json:"chunk_timeout,omitempty"`

	// EnableReconnect enables automatic reconnection on stream errors.
	// Default is false.
	EnableReconnect bool `json:"enable_reconnect,omitempty"`

	// MaxReconnectAttempts is the maximum number of reconnection attempts.
	// Only used if EnableReconnect is true.
	// Default is 3.
	MaxReconnectAttempts int `json:"max_reconnect_attempts,omitempty"`

	// ReconnectBackoff is the backoff duration between reconnection attempts.
	// Default is 1 second.
	ReconnectBackoff time.Duration `json:"reconnect_backoff,omitempty"`
}

// CacheConfig contains caching configuration.
type CacheConfig struct {
	// Enabled enables response caching.
	// Default is false.
	Enabled bool `json:"enabled,omitempty"`

	// TTL is the time-to-live for cached responses.
	// Zero means cached responses never expire.
	TTL time.Duration `json:"ttl,omitempty"`

	// MaxSize is the maximum number of cached responses.
	// Zero means no size limit.
	MaxSize int `json:"max_size,omitempty"`

	// MaxMemoryBytes is the maximum memory usage for the cache in bytes.
	// Zero means no memory limit.
	MaxMemoryBytes int64 `json:"max_memory_bytes,omitempty"`

	// CacheEmbeddings enables caching for embedding requests.
	// Default is false.
	CacheEmbeddings bool `json:"cache_embeddings,omitempty"`

	// CacheCompletions enables caching for completion requests.
	// Default is false.
	CacheCompletions bool `json:"cache_completions,omitempty"`
}

// RateLimitConfig contains rate limiting configuration.
type RateLimitConfig struct {
	// Enabled enables rate limiting.
	// Default is false.
	Enabled bool `json:"enabled,omitempty"`

	// RequestsPerSecond is the maximum number of requests per second.
	// Zero means no rate limiting.
	RequestsPerSecond float64 `json:"requests_per_second,omitempty"`

	// RequestsPerMinute is the maximum number of requests per minute.
	// Zero means no rate limiting.
	RequestsPerMinute int `json:"requests_per_minute,omitempty"`

	// TokensPerMinute is the maximum number of tokens per minute.
	// Zero means no token-based rate limiting.
	TokensPerMinute int `json:"tokens_per_minute,omitempty"`

	// Burst is the maximum burst size for the rate limiter.
	Burst int `json:"burst,omitempty"`

	// WaitTimeout is the maximum time to wait for rate limit availability.
	// Zero means requests fail immediately when rate limited.
	WaitTimeout time.Duration `json:"wait_timeout,omitempty"`
}

// LoggingConfig contains logging configuration.
type LoggingConfig struct {
	// Enabled enables logging.
	// Default is false.
	Enabled bool `json:"enabled,omitempty"`

	// Level is the log level (debug, info, warn, error).
	// Default is "info".
	Level string `json:"level,omitempty"`

	// LogRequests enables logging of outgoing requests.
	// Default is false.
	LogRequests bool `json:"log_requests,omitempty"`

	// LogResponses enables logging of incoming responses.
	// Default is false.
	LogResponses bool `json:"log_responses,omitempty"`

	// LogErrors enables logging of errors.
	// Default is true.
	LogErrors bool `json:"log_errors,omitempty"`

	// LogTokenUsage enables logging of token usage.
	// Default is false.
	LogTokenUsage bool `json:"log_token_usage,omitempty"`

	// RedactAPIKey redacts API keys in logs.
	// Default is true.
	RedactAPIKey bool `json:"redact_api_key,omitempty"`

	// Format is the log format (text, json).
	// Default is "text".
	Format string `json:"format,omitempty"`
}

// ValidationConfig contains validation configuration.
type ValidationConfig struct {
	// Enabled enables request validation.
	// Default is true.
	Enabled bool `json:"enabled,omitempty"`

	// StrictMode enables strict validation that rejects any non-standard fields.
	// Default is false.
	StrictMode bool `json:"strict_mode,omitempty"`

	// ValidateOnSend validates requests before sending to the provider.
	// Default is true.
	ValidateOnSend bool `json:"validate_on_send,omitempty"`

	// ValidateOnReceive validates responses received from the provider.
	// Default is false.
	ValidateOnReceive bool `json:"validate_on_receive,omitempty"`
}

// Default configuration values

// DefaultTimeout is the default timeout for API requests.
const DefaultTimeout = 60 * time.Second

// DefaultMaxRetries is the default maximum number of retry attempts.
const DefaultMaxRetries = 3

// DefaultInitialBackoff is the default initial backoff duration.
const DefaultInitialBackoff = 1 * time.Second

// DefaultMaxBackoff is the default maximum backoff duration.
const DefaultMaxBackoff = 60 * time.Second

// DefaultBackoffMultiplier is the default backoff multiplier.
const DefaultBackoffMultiplier = 2.0

// DefaultStreamBufferSize is the default buffer size for streaming.
const DefaultStreamBufferSize = 100

// DefaultUserAgent is the default User-Agent header.
const DefaultUserAgent = "go-ai-types/0.1.0"

// NewDefaultClientConfig returns a ClientConfig with default values.
func NewDefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		Timeout:   DefaultTimeout,
		UserAgent: DefaultUserAgent,
		HTTPConfig: &HTTPConfig{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
		RetryConfig: &RetryConfig{
			MaxRetries:        DefaultMaxRetries,
			InitialBackoff:    DefaultInitialBackoff,
			MaxBackoff:        DefaultMaxBackoff,
			BackoffMultiplier: DefaultBackoffMultiplier,
		},
	}
}

// WithAPIKey returns a copy of the config with the API key set.
func (c *ClientConfig) WithAPIKey(apiKey string) *ClientConfig {
	config := *c
	config.APIKey = apiKey
	return &config
}

// WithBaseURL returns a copy of the config with the base URL set.
func (c *ClientConfig) WithBaseURL(baseURL string) *ClientConfig {
	config := *c
	config.BaseURL = baseURL
	return &config
}

// WithTimeout returns a copy of the config with the timeout set.
func (c *ClientConfig) WithTimeout(timeout time.Duration) *ClientConfig {
	config := *c
	config.Timeout = timeout
	return &config
}

// WithDefaultModel returns a copy of the config with the default model set.
func (c *ClientConfig) WithDefaultModel(model string) *ClientConfig {
	config := *c
	config.DefaultModel = model
	return &config
}

// WithDebug returns a copy of the config with debug mode set.
func (c *ClientConfig) WithDebug(debug bool) *ClientConfig {
	config := *c
	config.Debug = debug
	return &config
}
