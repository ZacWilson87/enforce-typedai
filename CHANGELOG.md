# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure
- Go module initialization
- Development tooling configuration
- Project documentation (README, CONTRIBUTING, LICENSE)
- Comprehensive implementation plan

### Phase 1: Foundation Setup ✅
- Initialized Go module (github.com/zacw/go-ai-types)
- Created directory structure for all packages
- Configured development tools (.gitignore, .editorconfig, .golangci.yml)
- Created Makefile with common development tasks
- Added MIT License
- Created comprehensive README with project overview
- Added CONTRIBUTING guidelines

### Phase 2: Core Type Definitions ✅
- Implemented `pkg/types/common.go` - Core enums (Role, ContentType, FinishReason, ModelCapability, Provider, ImageDetail)
- Implemented `pkg/types/error.go` - Comprehensive error types (AIError interface, ProviderError, ValidationError, helper functions)
- Implemented `pkg/types/metadata.go` - Metadata types (Usage, MessageMetadata, ResponseMetadata, RequestMetadata, ModelInfo)
- Implemented `pkg/types/message.go` - Message types (Message, Content interface, TextContent, ImageContent, AudioContent, MultiContent)
- Implemented `pkg/types/function.go` - Function calling types (ToolDefinition, FunctionDefinition, ToolCall, FunctionCall, JSONSchema, ResponseFormat)
- Implemented `pkg/types/chat.go` - Chat completion types (ChatRequest, ChatResponse, Choice, LogProbability, helper methods)
- Implemented `pkg/types/stream.go` - Streaming types (StreamChunk, ChatStreamChunk, StreamChoice, MessageDelta, StreamAccumulator)
- Implemented `pkg/types/embedding.go` - Embedding types (EmbeddingRequest, EmbeddingResponse, Embedding, helper methods)
- Implemented `pkg/types/token.go` - Token management types (TokenEstimate, TokenCounter, TokenLimit, TokenPricing, TokenBudget)
- Added comprehensive godoc documentation for all types
- Total: 2,118 lines of production code
- All code compiles and passes go vet

### Documentation Updates
- Updated `project_plan.md` Phase 4 to include self-healing/repair logic
  - Added repair types, configuration, and orchestration details
  - Defined deterministic repair flow with bounded retry logic
  - Integrated SELF_HEALING.md requirements into implementation plan
- Updated `README.md` to feature self-healing capability
  - Added repair logic to supported features list
  - Documented `pkg/repair` package architecture
  - Added deterministic repair design principle
- Integrated SELF_HEALING.md specifications into project documentation

## [0.1.0] - TBD

### Planned Features
- Core type definitions (Message, ChatRequest, ChatResponse)
- Provider and service interfaces
- Validation utilities
- Builder patterns for complex types
- Provider-specific converters (OpenAI, Anthropic)
- Comprehensive examples
- Full test coverage (>90%)

---

## Version History

### Phase Roadmap

- **Phase 1**: Foundation Setup ✅
- **Phase 2**: Core Type Definitions ✅ (Current)
- **Phase 3**: Interfaces and Abstractions (Next)
- **Phase 4**: Validators and Builders
- **Phase 5**: Type Converters
- **Phase 6**: Examples and Documentation
- **Phase 7**: Testing and Quality
- **Phase 8**: Release Preparation

---

[Unreleased]: https://github.com/zacw/go-ai-types/compare/main...HEAD
