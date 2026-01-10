package types

import "fmt"

// ErrorType represents the category of an AI API error.
type ErrorType string

const (
	// ErrorTypeInvalidRequest indicates the request was malformed or invalid.
	ErrorTypeInvalidRequest ErrorType = "invalid_request_error"

	// ErrorTypeAuthentication indicates authentication failed.
	ErrorTypeAuthentication ErrorType = "authentication_error"

	// ErrorTypePermission indicates insufficient permissions.
	ErrorTypePermission ErrorType = "permission_error"

	// ErrorTypeNotFound indicates the requested resource was not found.
	ErrorTypeNotFound ErrorType = "not_found_error"

	// ErrorTypeRateLimit indicates rate limit was exceeded.
	ErrorTypeRateLimit ErrorType = "rate_limit_error"

	// ErrorTypeQuotaExceeded indicates quota/credit limit was exceeded.
	ErrorTypeQuotaExceeded ErrorType = "quota_exceeded_error"

	// ErrorTypeServer indicates an internal server error occurred.
	ErrorTypeServer ErrorType = "server_error"

	// ErrorTypeTimeout indicates the request timed out.
	ErrorTypeTimeout ErrorType = "timeout_error"

	// ErrorTypeContentFilter indicates content was blocked by filters.
	ErrorTypeContentFilter ErrorType = "content_filter_error"

	// ErrorTypeValidation indicates validation failed.
	ErrorTypeValidation ErrorType = "validation_error"

	// ErrorTypeUnknown indicates an unknown error occurred.
	ErrorTypeUnknown ErrorType = "unknown_error"
)

// String returns the string representation of the ErrorType.
func (e ErrorType) String() string {
	return string(e)
}

// AIError represents an error that occurred during AI API operations.
// It extends the standard error interface with additional context.
type AIError interface {
	error

	// Type returns the category of error.
	Type() ErrorType

	// Code returns the provider-specific error code.
	Code() string

	// StatusCode returns the HTTP status code (if applicable).
	StatusCode() int

	// Provider returns the provider where the error occurred.
	Provider() Provider

	// Retryable indicates if the operation can be retried.
	Retryable() bool
}

// ProviderError represents an error from an AI provider.
type ProviderError struct {
	// ErrorType categorizes the error.
	ErrorType ErrorType `json:"type"`

	// Message is the human-readable error message.
	Message string `json:"message"`

	// ErrorCode is the provider-specific error code.
	ErrorCode string `json:"code,omitempty"`

	// Param is the parameter that caused the error (if applicable).
	Param string `json:"param,omitempty"`

	// HTTPStatus is the HTTP status code.
	HTTPStatus int `json:"status_code,omitempty"`

	// ProviderName is the name of the provider.
	ProviderName Provider `json:"provider,omitempty"`

	// IsRetryable indicates if the error is transient and can be retried.
	IsRetryable bool `json:"retryable,omitempty"`

	// InnerError is the underlying error (if any).
	InnerError error `json:"-"`
}

// Error implements the error interface.
func (e *ProviderError) Error() string {
	if e.ErrorCode != "" {
		return fmt.Sprintf("%s: %s (code: %s)", e.ErrorType, e.Message, e.ErrorCode)
	}
	return fmt.Sprintf("%s: %s", e.ErrorType, e.Message)
}

// Type returns the error type.
func (e *ProviderError) Type() ErrorType {
	return e.ErrorType
}

// Code returns the provider-specific error code.
func (e *ProviderError) Code() string {
	return e.ErrorCode
}

// StatusCode returns the HTTP status code.
func (e *ProviderError) StatusCode() int {
	return e.HTTPStatus
}

// Provider returns the provider name.
func (e *ProviderError) Provider() Provider {
	return e.ProviderName
}

// Retryable indicates if the error is retryable.
func (e *ProviderError) Retryable() bool {
	return e.IsRetryable
}

// Unwrap returns the underlying error for error wrapping support.
func (e *ProviderError) Unwrap() error {
	return e.InnerError
}

// NewProviderError creates a new ProviderError.
func NewProviderError(errType ErrorType, message string) *ProviderError {
	return &ProviderError{
		ErrorType: errType,
		Message:   message,
	}
}

// NewProviderErrorWithCode creates a new ProviderError with a code.
func NewProviderErrorWithCode(errType ErrorType, message, code string) *ProviderError {
	return &ProviderError{
		ErrorType: errType,
		Message:   message,
		ErrorCode: code,
	}
}

// ValidationError represents a validation error.
type ValidationError struct {
	// Field is the name of the field that failed validation.
	Field string `json:"field"`

	// Message is the validation error message.
	Message string `json:"message"`

	// Value is the invalid value (optional, for debugging).
	Value interface{} `json:"value,omitempty"`
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// NewValidationError creates a new ValidationError.
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// Helper functions for error type checking

// IsRateLimitError returns true if the error is a rate limit error.
func IsRateLimitError(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Type() == ErrorTypeRateLimit
	}
	return false
}

// IsAuthError returns true if the error is an authentication error.
func IsAuthError(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Type() == ErrorTypeAuthentication
	}
	return false
}

// IsInvalidRequestError returns true if the error is an invalid request error.
func IsInvalidRequestError(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Type() == ErrorTypeInvalidRequest
	}
	return false
}

// IsServerError returns true if the error is a server error.
func IsServerError(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Type() == ErrorTypeServer
	}
	return false
}

// IsTimeoutError returns true if the error is a timeout error.
func IsTimeoutError(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Type() == ErrorTypeTimeout
	}
	return false
}

// IsRetryable returns true if the error is retryable.
func IsRetryable(err error) bool {
	if aiErr, ok := err.(AIError); ok {
		return aiErr.Retryable()
	}
	return false
}
