// Package types provides core type definitions for AI/ML service interactions.
//
// This package contains standardized, type-safe data structures for working with
// AI providers in a consistent, provider-agnostic manner. It includes types for:
//
//   - Messages and content (text, images, audio)
//   - Chat completions and streaming
//   - Function/tool calling
//   - Embeddings generation
//   - Token usage and metadata
//   - Error handling
//
// The types in this package are designed to work across multiple AI providers
// (OpenAI, Anthropic, Google, etc.) while maintaining strong type safety and
// idiomatic Go patterns.
//
// Example usage:
//
//	msg := &types.Message{
//		Role:    types.RoleUser,
//		Content: types.NewTextContent("Hello, AI!"),
//	}
//
//	req := &types.ChatRequest{
//		Model:       "gpt-4",
//		Messages:    []*types.Message{msg},
//		Temperature: 0.7,
//		MaxTokens:   100,
//	}
package types
