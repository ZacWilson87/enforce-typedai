# Go AI Types

A comprehensive, type-safe Go library providing standardized data structures and interfaces for interacting with AI/ML services and models.

## Overview

Go AI Types is a foundational package that enables developers to work with AI services in a type-safe, provider-agnostic manner. The library provides common types, interfaces, and utilities that work seamlessly across multiple AI providers including OpenAI, Anthropic, Google, and others.

## Features

- **Type Safety**: Strong typing to prevent runtime errors when working with AI APIs
- **Provider Agnostic**: Unified abstractions that work across multiple AI providers
- **Zero Dependencies**: Core types have no external runtime dependencies
- **Comprehensive Coverage**: Support for chat completions, embeddings, function calling, streaming, and more
- **Idiomatic Go**: Follows Go best practices and idioms
- **Well Tested**: >90% test coverage with comprehensive unit and integration tests
- **Fully Documented**: Complete godoc documentation and examples

## Status

🚧 **Currently in Development** - v0.x.x

This library is under active development. APIs may change before v1.0.0 release.

## Supported Features

- ✅ Chat completions (non-streaming and streaming)
- ✅ Embeddings generation
- ✅ Function/tool calling
- ✅ Multi-modal content (text, images, audio)
- ✅ Token counting and usage tracking
- ✅ Comprehensive error handling
- ✅ Request validation and builders
- ✅ Provider-specific converters
- 🔄 **Self-healing/repair logic** for invalid outputs (Phase 4)

## Installation

```bash
go get github.com/zacw/go-ai-types
```

## Quick Start

```go
package main

import (
    "github.com/zacw/go-ai-types/pkg/types"
    "github.com/zacw/go-ai-types/pkg/builders"
)

func main() {
    // Build a message using the fluent builder API
    msg := builders.NewMessageBuilder().
        Role(types.RoleUser).
        TextContent("Hello, AI!").
        Build()

    // Create a chat request
    req := &types.ChatRequest{
        Model:    "gpt-4",
        Messages: []*types.Message{msg},
        Temperature: 0.7,
        MaxTokens: 100,
    }

    // Use with your preferred AI provider client...
}
```

## Architecture

The library is organized into several key packages:

### `pkg/types`
Core type definitions including:
- Message structures
- Request/Response types
- Function calling types
- Streaming types
- Error types

### `pkg/interfaces`
Provider and service interfaces that define contracts for implementations.

### `pkg/validators`
Validation functions to ensure data integrity before API calls.

### `pkg/builders`
Fluent builder APIs for constructing complex types.

### `pkg/converters`
Converters for different AI provider formats.

### `pkg/repair` (Phase 4)
Deterministic self-healing logic for handling invalid AI outputs:
- Bounded repair attempts (configurable, max 3)
- Schema validation and structural repair
- Explicit error handling and logging
- Opt-in configuration with full transparency

See [SELF_HEALING.md](./SELF_HEALING.md) for detailed repair specifications.

## Examples

See the [examples](./examples) directory for complete, runnable examples:

- [Basic Chat](./examples/basic_chat) - Simple chat completion
- [Streaming](./examples/streaming) - Streaming chat responses
- [Function Calling](./examples/function_calling) - Using function/tool calling
- [Multi-Provider](./examples/multi_provider) - Working with multiple providers

## Documentation

- [Architecture](./docs/architecture.md) - Architecture decisions and rationale
- [Design Patterns](./docs/design_patterns.md) - Design patterns used
- [Getting Started](./docs/getting_started.md) - Comprehensive getting started guide
- [API Documentation](https://pkg.go.dev/github.com/zacw/go-ai-types) - Auto-generated API docs

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, but recommended)

### Setup

```bash
# Clone the repository
git clone https://github.com/zacw/go-ai-types.git
cd go-ai-types

# Install development tools
make install-tools

# Run tests
make test

# Run linters
make lint

# Format code
make fmt
```

### Running Tests

```bash
# Unit tests only
make test

# Integration tests (requires API keys)
make test-integration

# All tests
make test-all

# Generate coverage report
make coverage
```

## Design Principles

1. **Type Safety First**: Leverage Go's type system to catch errors at compile time
2. **Provider Agnostic**: Core types work across all providers
3. **Zero Magic**: Explicit, predictable behavior with no hidden complexity
4. **Idiomatic Go**: Follow established Go conventions and best practices
5. **Minimal Dependencies**: Keep the dependency footprint small
6. **Comprehensive Testing**: High test coverage with robust test suites
7. **Deterministic Repair**: Self-healing logic is bounded, explicit, and transparent (no randomness or adaptive behavior)

## Versioning

This project follows [Semantic Versioning](https://semver.org/):
- v0.x.x - Initial development (current)
- v1.x.x - Stable API (future)

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](./LICENSE) for details.

## Roadmap

See [project_plan.md](./project_plan.md) for the detailed implementation roadmap.

### Current Phase: Phase 1 - Foundation Setup ✅

### Next Phases:
- Phase 2: Core Type Definitions
- Phase 3: Interfaces and Abstractions
- Phase 4: Validators and Builders
- Phase 5: Type Converters
- Phase 6: Examples and Documentation
- Phase 7: Testing and Quality
- Phase 8: Release Preparation

## Acknowledgments

This library is designed to be provider-agnostic and does not directly compete with or replace provider-specific SDKs. Instead, it provides common types and interfaces that can be used alongside provider SDKs.

## Support

For questions, issues, or feature requests, please open an issue on GitHub.
