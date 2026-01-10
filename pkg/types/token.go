package types

// TokenEstimate represents an estimated token count.
type TokenEstimate struct {
	// PromptTokens is the estimated number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`

	// CompletionTokens is the estimated number of tokens in the completion.
	CompletionTokens int `json:"completion_tokens"`

	// TotalTokens is the estimated total number of tokens.
	TotalTokens int `json:"total_tokens"`

	// Method describes how the estimate was calculated.
	Method string `json:"method,omitempty"`
}

// TokenCounter provides token counting functionality.
type TokenCounter interface {
	// CountTokens counts the number of tokens in the given text.
	CountTokens(text string) int

	// CountMessagesTokens counts the tokens in a list of messages.
	CountMessagesTokens(messages []*Message) int

	// EstimateRequestTokens estimates the tokens for a chat request.
	EstimateRequestTokens(req *ChatRequest) *TokenEstimate
}

// TokenLimit represents token limits for a model.
type TokenLimit struct {
	// Model is the model ID.
	Model string `json:"model"`

	// MaxTokens is the maximum total tokens (context window).
	MaxTokens int `json:"max_tokens"`

	// MaxInputTokens is the maximum input/prompt tokens.
	MaxInputTokens int `json:"max_input_tokens,omitempty"`

	// MaxOutputTokens is the maximum output/completion tokens.
	MaxOutputTokens int `json:"max_output_tokens,omitempty"`

	// MaxContextTokens is the maximum context window size.
	// This is typically the same as MaxTokens but can differ.
	MaxContextTokens int `json:"max_context_tokens,omitempty"`
}

// CanAccommodate checks if the model can accommodate the given token counts.
func (l *TokenLimit) CanAccommodate(promptTokens, completionTokens int) bool {
	totalNeeded := promptTokens + completionTokens

	// Check against max total tokens
	if l.MaxTokens > 0 && totalNeeded > l.MaxTokens {
		return false
	}

	// Check against max input tokens
	if l.MaxInputTokens > 0 && promptTokens > l.MaxInputTokens {
		return false
	}

	// Check against max output tokens
	if l.MaxOutputTokens > 0 && completionTokens > l.MaxOutputTokens {
		return false
	}

	return true
}

// AvailableOutputTokens returns the number of tokens available for output
// given the current prompt token count.
func (l *TokenLimit) AvailableOutputTokens(promptTokens int) int {
	// Start with the maximum output tokens
	available := l.MaxOutputTokens
	if available == 0 {
		available = l.MaxTokens
	}

	// Subtract prompt tokens from total capacity
	if l.MaxTokens > 0 {
		remaining := l.MaxTokens - promptTokens
		if remaining < available {
			available = remaining
		}
	}

	if available < 0 {
		return 0
	}
	return available
}

// TokenPricing represents the pricing for token usage.
type TokenPricing struct {
	// Model is the model ID.
	Model string `json:"model"`

	// PromptTokenPrice is the price per prompt token (in dollars or credits).
	PromptTokenPrice float64 `json:"prompt_token_price"`

	// CompletionTokenPrice is the price per completion token.
	CompletionTokenPrice float64 `json:"completion_token_price"`

	// CachedTokenPrice is the price per cached token (if applicable).
	CachedTokenPrice float64 `json:"cached_token_price,omitempty"`

	// Currency is the currency for pricing (e.g., "USD", "credits").
	Currency string `json:"currency,omitempty"`

	// Per1000Tokens indicates if prices are per 1000 tokens.
	Per1000Tokens bool `json:"per_1000_tokens,omitempty"`
}

// CalculateCost calculates the cost for the given usage.
func (p *TokenPricing) CalculateCost(usage *Usage) float64 {
	if usage == nil {
		return 0
	}

	var cost float64

	// Calculate prompt token cost
	promptTokens := float64(usage.PromptTokens)
	if p.Per1000Tokens {
		cost += (promptTokens / 1000.0) * p.PromptTokenPrice
	} else {
		cost += promptTokens * p.PromptTokenPrice
	}

	// Calculate completion token cost
	completionTokens := float64(usage.CompletionTokens)
	if p.Per1000Tokens {
		cost += (completionTokens / 1000.0) * p.CompletionTokenPrice
	} else {
		cost += completionTokens * p.CompletionTokenPrice
	}

	// Calculate cached token cost (if applicable)
	if usage.CachedTokens > 0 && p.CachedTokenPrice > 0 {
		cachedTokens := float64(usage.CachedTokens)
		if p.Per1000Tokens {
			cost += (cachedTokens / 1000.0) * p.CachedTokenPrice
		} else {
			cost += cachedTokens * p.CachedTokenPrice
		}
	}

	return cost
}

// EstimateCost estimates the cost for the given token counts.
func (p *TokenPricing) EstimateCost(promptTokens, completionTokens, cachedTokens int) float64 {
	return p.CalculateCost(&Usage{
		PromptTokens:     promptTokens,
		CompletionTokens: completionTokens,
		CachedTokens:     cachedTokens,
		TotalTokens:      promptTokens + completionTokens,
	})
}

// TokenBudget helps manage token budgets for requests.
type TokenBudget struct {
	// TotalBudget is the total token budget.
	TotalBudget int `json:"total_budget"`

	// ReservedForOutput is the number of tokens reserved for output.
	ReservedForOutput int `json:"reserved_for_output"`

	// Used tracks the number of tokens used so far.
	Used int `json:"used"`
}

// NewTokenBudget creates a new TokenBudget.
func NewTokenBudget(total, reservedForOutput int) *TokenBudget {
	return &TokenBudget{
		TotalBudget:       total,
		ReservedForOutput: reservedForOutput,
	}
}

// Remaining returns the number of tokens remaining in the budget.
func (b *TokenBudget) Remaining() int {
	remaining := b.TotalBudget - b.Used - b.ReservedForOutput
	if remaining < 0 {
		return 0
	}
	return remaining
}

// CanUse checks if the given number of tokens can be used.
func (b *TokenBudget) CanUse(tokens int) bool {
	return b.Remaining() >= tokens
}

// Use marks the given number of tokens as used.
func (b *TokenBudget) Use(tokens int) bool {
	if !b.CanUse(tokens) {
		return false
	}
	b.Used += tokens
	return true
}

// Reset resets the used token count.
func (b *TokenBudget) Reset() {
	b.Used = 0
}

// AvailableForPrompt returns the number of tokens available for the prompt.
func (b *TokenBudget) AvailableForPrompt() int {
	return b.Remaining()
}

// AvailableForOutput returns the number of tokens reserved for output.
func (b *TokenBudget) AvailableForOutput() int {
	return b.ReservedForOutput
}
