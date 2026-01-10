# Go AI Type Library - Project Structure

This document outlines the complete directory structure for the Go AI Type Library project.

## Directory Layout

```
go_ai_type_library/
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── README.md                        # Project overview and documentation
├── LICENSE                          # MIT License (recommended)
├── CONTRIBUTING.md                  # Contribution guidelines
├── CHANGELOG.md                     # Version history and changes
├── Makefile                         # Build and test automation
├── .gitignore                       # Git ignore patterns
├── .editorconfig                    # Editor configuration
├── .golangci.yml                    # Linter configuration
│
├── pkg/                             # Public API packages
│   ├── types/                       # Core type definitions
│   │   ├── doc.go                   # Package documentation
│   │   ├── common.go                # Common types (Role, ContentType, etc.)
│   │   ├── message.go               # Message types
│   │   ├── chat.go                  # Chat completion types
│   │   ├── embedding.go             # Embedding types
│   │   ├── function.go              # Function calling types
│   │   ├── image.go                 # Image-related types
│   │   ├── audio.go                 # Audio-related types
│   │   ├── stream.go                # Streaming types
│   │   ├── token.go                 # Token counting types
│   │   ├── error.go                 # Error types
│   │   ├── metadata.go              # Metadata and usage types
│   │   ├── common_test.go           # Tests for common.go
│   │   ├── message_test.go          # Tests for message.go
│   │   ├── chat_test.go             # Tests for chat.go
│   │   ├── embedding_test.go        # Tests for embedding.go
│   │   ├── function_test.go         # Tests for function.go
│   │   ├── image_test.go            # Tests for image.go
│   │   ├── audio_test.go            # Tests for audio.go
│   │   ├── stream_test.go           # Tests for stream.go
│   │   ├── token_test.go            # Tests for token.go
│   │   ├── error_test.go            # Tests for error.go
│   │   └── metadata_test.go         # Tests for metadata.go
│   │
│   ├── interfaces/                  # Interface definitions
│   │   ├── doc.go                   # Package documentation
│   │   ├── provider.go              # Provider interface
│   │   ├── chat.go                  # Chat service interface
│   │   ├── embedding.go             # Embedding service interface
│   │   ├── stream.go                # Stream handler interface
│   │   ├── middleware.go            # Middleware interfaces
│   │   ├── provider_test.go         # Tests for provider.go
│   │   ├── chat_test.go             # Tests for chat.go
│   │   ├── embedding_test.go        # Tests for embedding.go
│   │   ├── stream_test.go           # Tests for stream.go
│   │   └── middleware_test.go       # Tests for middleware.go
│   │
│   ├── validators/                  # Validation utilities
│   │   ├── doc.go                   # Package documentation
│   │   ├── message.go               # Message validation
│   │   ├── function.go              # Function schema validation
│   │   ├── common.go                # Common validators
│   │   ├── message_test.go          # Tests for message.go
│   │   ├── function_test.go         # Tests for function.go
│   │   └── common_test.go           # Tests for common.go
│   │
│   ├── builders/                    # Builder patterns for complex types
│   │   ├── doc.go                   # Package documentation
│   │   ├── message.go               # Message builders
│   │   ├── chat.go                  # Chat request builders
│   │   ├── function.go              # Function definition builders
│   │   ├── message_test.go          # Tests for message.go
│   │   ├── chat_test.go             # Tests for chat.go
│   │   └── function_test.go         # Tests for function.go
│   │
│   └── converters/                  # Type converters for different providers
│       ├── doc.go                   # Package documentation
│       ├── openai.go                # OpenAI format converters
│       ├── anthropic.go             # Anthropic format converters
│       ├── common.go                # Common conversion utilities
│       ├── openai_test.go           # Tests for openai.go
│       ├── anthropic_test.go        # Tests for anthropic.go
│       └── common_test.go           # Tests for common.go
│
├── internal/                        # Private implementation packages
│   ├── validation/                  # Internal validation logic
│   │   ├── schema.go                # JSON schema validation
│   │   └── schema_test.go           # Tests for schema.go
│   │
│   └── utils/                       # Internal utilities
│       ├── helpers.go               # Helper functions
│       └── helpers_test.go          # Tests for helpers.go
│
├── examples/                        # Example usage
│   ├── basic_chat/                  # Basic chat example
│   │   ├── main.go                  # Example code
│   │   └── README.md                # Example documentation
│   │
│   ├── streaming/                   # Streaming example
│   │   ├── main.go                  # Example code
│   │   └── README.md                # Example documentation
│   │
│   ├── function_calling/            # Function calling example
│   │   ├── main.go                  # Example code
│   │   └── README.md                # Example documentation
│   │
│   └── multi_provider/              # Multi-provider example
│       ├── main.go                  # Example code
│       └── README.md                # Example documentation
│
├── docs/                            # Documentation
│   ├── architecture.md              # Architecture decisions
│   ├── design_patterns.md           # Design patterns used
│   ├── getting_started.md           # Getting started guide
│   ├── migration_guides/            # Migration guides
│   │   └── v0-to-v1.md             # Example migration guide
│   │
│   └── api/                         # API documentation
│       ├── types.md                 # Types documentation
│       ├── interfaces.md            # Interfaces documentation
│       ├── validators.md            # Validators documentation
│       ├── builders.md              # Builders documentation
│       └── converters.md            # Converters documentation
│
├── tests/                           # Integration and benchmarks
│   ├── integration/                 # Integration tests
│   │   ├── provider_test.go         # Provider integration tests
│   │   └── README.md                # Integration testing guide
│   │
│   └── benchmarks/                  # Performance benchmarks
│       ├── types_bench_test.go      # Type operation benchmarks
│       ├── validation_bench_test.go # Validation benchmarks
│       └── README.md                # Benchmark documentation
│
└── .github/                         # GitHub-specific files
    ├── workflows/                   # GitHub Actions workflows
    │   ├── ci.yml                   # CI pipeline
    │   ├── release.yml              # Release workflow
    │   └── lint.yml                 # Linting workflow
    │
    ├── ISSUE_TEMPLATE/              # Issue templates
    │   ├── bug_report.md            # Bug report template
    │   └── feature_request.md       # Feature request template
    │
    ├── PULL_REQUEST_TEMPLATE.md     # PR template
    └── CODEOWNERS                   # Code owners file
```

## Package Descriptions

### `pkg/types/`
The core package containing all type definitions. This is the most important package and will be used by all consumers of the library. It defines:
- Basic types (Role, ContentType, etc.)
- Message structures
- Request/Response types for chat, embeddings, etc.
- Function calling types
- Streaming types
- Error types

### `pkg/interfaces/`
Defines interfaces for providers and services. This package establishes the contracts that implementations must follow and enables dependency injection and testing.

### `pkg/validators/`
Provides validation functions for various types. These validators ensure data integrity before making API calls.

### `pkg/builders/`
Implements builder patterns for complex types, offering a fluent API for constructing messages, requests, and function definitions.

### `pkg/converters/`
Contains converters for different AI provider formats. This allows the library to work with multiple providers while maintaining a consistent internal representation.

### `internal/`
Private packages that are not exposed to library consumers. Used for internal implementation details.

### `examples/`
Runnable examples demonstrating how to use the library for various use cases.

### `docs/`
Comprehensive documentation including architecture decisions, design patterns, and API references.

### `tests/`
Integration tests and benchmarks that don't belong to specific packages.

## File Naming Conventions

- **Implementation files**: `<feature>.go` (e.g., `message.go`, `chat.go`)
- **Test files**: `<feature>_test.go` (e.g., `message_test.go`, `chat_test.go`)
- **Benchmark files**: `<feature>_bench_test.go` (e.g., `types_bench_test.go`)
- **Package documentation**: `doc.go` (contains package-level documentation)

## Key Configuration Files

### `go.mod`
Defines the module path and dependencies. Example:
```go
module github.com/username/go-ai-types

go 1.21

require (
    // dependencies will be added as needed
)
```

### `.gitignore`
Standard Go gitignore patterns:
```
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test coverage
*.out
coverage.html

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
```

### `Makefile`
Common tasks for development:
```makefile
.PHONY: test lint fmt build clean install-tools

test:
	go test -v -race -cover ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -s -w .
	goimports -w .

build:
	go build ./...

clean:
	go clean ./...

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
```

### `.golangci.yml`
Linter configuration:
```yaml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - ineffassign
    - typecheck
    - misspell

linters-settings:
  gofmt:
    simplify: true
  govet:
    check-shadowing: true

run:
  timeout: 5m
  tests: true
```

### `.editorconfig`
Editor configuration for consistent formatting:
```
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.go]
indent_style = tab
indent_size = 4

[*.{yml,yaml,json}]
indent_style = space
indent_size = 2

[Makefile]
indent_style = tab
```

## Testing Organization

- **Unit tests**: Co-located with implementation files (`*_test.go`)
- **Integration tests**: In `tests/integration/` directory
- **Benchmarks**: In `tests/benchmarks/` directory
- **Test coverage**: Aim for >90% coverage in `pkg/` packages

## Documentation Organization

- **Code comments**: Godoc-style comments for all exported types and functions
- **Package docs**: `doc.go` files in each package
- **External docs**: Comprehensive guides in `docs/` directory
- **Examples**: Runnable examples in `examples/` directory with accompanying READMEs

## Build Tags

Integration tests should use build tags:
```go
//go:build integration
// +build integration

package integration
```

This allows running unit tests separately from integration tests:
```bash
go test ./...                    # Unit tests only
go test -tags=integration ./...  # Including integration tests
```

## Next Steps

Once this structure is approved, we can proceed with Phase 1 of the implementation plan:
1. Initialize the Go module
2. Create the basic directory structure
3. Set up configuration files
4. Begin implementing Phase 2 (Core Type Definitions)
