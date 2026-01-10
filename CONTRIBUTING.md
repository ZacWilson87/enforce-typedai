# Contributing to Go AI Types

Thank you for your interest in contributing to Go AI Types! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:
- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Go version and operating system
- Any relevant code snippets or error messages

### Suggesting Features

Feature suggestions are welcome! Please open an issue with:
- A clear description of the feature
- Use cases and benefits
- Example API usage (if applicable)
- Any alternative solutions you've considered

### Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following the coding standards below
3. **Add tests** for any new functionality
4. **Update documentation** as needed
5. **Run the test suite** to ensure all tests pass
6. **Run the linters** to ensure code quality
7. **Submit a pull request** with a clear description

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/go-ai-types.git
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

## Coding Standards

### General Guidelines

1. **Follow Go conventions**: Use `gofmt`, `goimports`, and follow [Effective Go](https://golang.org/doc/effective_go.html)
2. **Write idiomatic Go**: Prefer simple, clear code over clever solutions
3. **Keep it simple**: Avoid unnecessary abstractions and complexity
4. **Be explicit**: Prefer explicit code over implicit behavior

### Naming Conventions

- Use `CamelCase` for exported types and functions
- Use `camelCase` for unexported types and functions
- Use descriptive names that clearly indicate purpose
- Avoid abbreviations unless they're widely understood (e.g., `HTTP`, `JSON`)

### Documentation

- All exported types, functions, and constants must have godoc comments
- Comments should start with the name of the item being documented
- Use complete sentences with proper punctuation
- Include examples for complex functionality

Example:
```go
// Message represents a single message in a conversation.
// It contains the role of the speaker, the content, and optional metadata.
//
// Example:
//   msg := &Message{
//       Role:    RoleUser,
//       Content: NewTextContent("Hello"),
//   }
type Message struct {
    // ...
}
```

### Testing

- Write table-driven tests where appropriate
- Aim for >90% code coverage
- Test both success and error cases
- Use meaningful test names that describe what's being tested

Example:
```go
func TestMessageValidation(t *testing.T) {
    tests := []struct {
        name    string
        message *Message
        wantErr bool
    }{
        {
            name: "valid message",
            message: &Message{
                Role:    RoleUser,
                Content: NewTextContent("test"),
            },
            wantErr: false,
        },
        {
            name: "missing role",
            message: &Message{
                Content: NewTextContent("test"),
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateMessage(tt.message)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Error Handling

- Return errors rather than panicking
- Wrap errors with context using `fmt.Errorf` with `%w`
- Use custom error types for domain-specific errors
- Provide helpful error messages

Example:
```go
if err := validate(input); err != nil {
    return fmt.Errorf("failed to validate input: %w", err)
}
```

### Package Organization

- Keep packages focused and cohesive
- Use `internal/` for private implementation details
- Use `pkg/` for public APIs
- Co-locate tests with implementation files

## Commit Messages

Write clear, descriptive commit messages:

```
Short summary (50 chars or less)

More detailed explanation if needed. Wrap at 72 characters.
Explain what changed and why, not how.

- Bullet points are okay
- Use present tense ("Add feature" not "Added feature")
- Reference issues: "Fixes #123"
```

## Pull Request Process

1. Update documentation for any API changes
2. Add tests for new functionality
3. Ensure all tests pass: `make test`
4. Ensure linters pass: `make lint`
5. Update CHANGELOG.md with your changes
6. Request review from maintainers

### PR Description Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
Describe the tests you added or updated

## Checklist
- [ ] Tests pass locally
- [ ] Linters pass locally
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
```

## Code Review

All submissions require review. We use GitHub pull requests for this purpose.

Reviewers will check for:
- Code quality and style
- Test coverage
- Documentation completeness
- API design consistency
- Performance implications

## Release Process

Maintainers handle releases following this process:

1. Update CHANGELOG.md with release notes
2. Update version in documentation
3. Create and push a git tag
4. GitHub Actions will automatically publish the release

## Questions?

If you have questions about contributing, feel free to open an issue or reach out to the maintainers.

Thank you for contributing to Go AI Types!
