// Package interfaces defines the core abstractions and contracts for interacting
// with AI service providers.
//
// This package provides a set of interfaces that enable provider-agnostic
// implementations of AI services. By programming to these interfaces rather than
// concrete implementations, applications can easily switch between providers or
// work with multiple providers simultaneously.
//
// # Core Interfaces
//
// Provider: Represents an AI service provider (OpenAI, Anthropic, Google, etc.)
// and provides access to its services.
//
// ChatService: Handles chat completion requests, both streaming and non-streaming.
//
// EmbeddingService: Handles embedding generation requests.
//
// StreamHandler: Processes streaming responses in a callback-based manner.
//
// Middleware: Enables composable request/response processing for cross-cutting
// concerns like logging, rate limiting, retries, and caching.
//
// # Design Principles
//
// 1. Provider Agnostic: Interfaces abstract away provider-specific details
// 2. Context-Aware: All operations accept context.Context for cancellation and timeouts
// 3. Composable: Middleware enables building complex behavior from simple components
// 4. Type Safe: Leverages Go's type system to prevent runtime errors
// 5. Testable: Interfaces make it easy to create mocks and test implementations
//
// # Example Usage
//
//	// Create a provider (implementation not shown)
//	provider := openai.NewProvider(apiKey)
//
//	// Get the chat service
//	chatService := provider.ChatService()
//
//	// Create a request
//	req := &types.ChatRequest{
//	    Model: "gpt-4",
//	    Messages: []*types.Message{
//	        {Role: types.RoleUser, Content: types.NewTextContent("Hello!")},
//	    },
//	}
//
//	// Execute the request
//	resp, err := chatService.CreateCompletion(ctx, req)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Process the response
//	fmt.Println(resp.GetFirstContent())
//
// # Streaming Example
//
//	// Create a streaming request
//	req.Stream = true
//
//	// Get the stream
//	stream, err := chatService.CreateCompletionStream(ctx, req)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Process chunks
//	for chunk := range stream {
//	    if chunk.IsComplete() {
//	        break
//	    }
//	    // Process chunk...
//	}
//
// # Middleware Example
//
//	// Apply middleware (implementation not shown)
//	chatService = logging.Wrap(chatService)
//	chatService = ratelimit.Wrap(chatService)
//	chatService = retry.Wrap(chatService)
//
//	// Now all requests go through the middleware chain
//	resp, err := chatService.CreateCompletion(ctx, req)
package interfaces
