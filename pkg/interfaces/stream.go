package interfaces

import (
	"context"

	"github.com/zacw/go-ai-types/pkg/types"
)

// StreamHandler processes streaming responses from AI providers.
//
// This interface provides a callback-based approach to handling streaming
// responses, which can be more convenient than channel-based streaming in
// some use cases.
//
// Implementations receive callbacks for each chunk, completion, and errors
// that occur during streaming.
//
// Example usage:
//
//	type MyHandler struct {
//	    accumulator *types.StreamAccumulator
//	}
//
//	func (h *MyHandler) OnChunk(chunk types.StreamChunk) error {
//	    h.accumulator.Add(chunk)
//	    // Process chunk
//	    return nil
//	}
//
//	func (h *MyHandler) OnComplete() error {
//	    resp := h.accumulator.ToChatResponse()
//	    // Process complete response
//	    return nil
//	}
//
//	func (h *MyHandler) OnError(err error) {
//	    log.Printf("Stream error: %v", err)
//	}
//
//	handler := &MyHandler{accumulator: types.NewStreamAccumulator()}
//	err := service.CreateCompletionStreamWithCallback(ctx, req, handler)
type StreamHandler interface {
	// OnChunk is called for each chunk received from the stream.
	//
	// Implementations should process the chunk and return nil to continue
	// streaming, or return an error to stop the stream.
	//
	// If an error is returned, streaming stops immediately and OnError is called.
	//
	// Example:
	//   func (h *MyHandler) OnChunk(chunk types.StreamChunk) error {
	//       choices := chunk.GetChoices()
	//       if len(choices) > 0 && choices[0].Delta != nil {
	//           fmt.Print(choices[0].Delta.Content)
	//       }
	//       return nil
	//   }
	OnChunk(chunk types.StreamChunk) error

	// OnComplete is called when the stream completes successfully.
	//
	// This is called after all chunks have been processed and the stream
	// has been closed normally (not due to an error or cancellation).
	//
	// If an error is returned, OnError is called with that error.
	//
	// Example:
	//   func (h *MyHandler) OnComplete() error {
	//       fmt.Println("\nStream completed")
	//       return nil
	//   }
	OnComplete() error

	// OnError is called when an error occurs during streaming.
	//
	// This includes:
	// - Network errors
	// - API errors
	// - Errors returned by OnChunk or OnComplete
	// - Context cancellation
	//
	// This method should not return an error. Any cleanup or error handling
	// should be done within this method.
	//
	// Example:
	//   func (h *MyHandler) OnError(err error) {
	//       log.Printf("Stream error: %v", err)
	//   }
	OnError(err error)
}

// StreamHandlerFunc is a function type that implements StreamHandler.
//
// This type allows using separate functions for chunk, complete, and error
// handling without defining a new struct type.
//
// Example:
//
//	handler := interfaces.NewStreamHandlerFunc(
//	    func(chunk types.StreamChunk) error {
//	        // Handle chunk
//	        return nil
//	    },
//	    func() error {
//	        // Handle completion
//	        return nil
//	    },
//	    func(err error) {
//	        // Handle error
//	    },
//	)
type StreamHandlerFunc struct {
	ChunkFunc    func(types.StreamChunk) error
	CompleteFunc func() error
	ErrorFunc    func(error)
}

// OnChunk implements StreamHandler.
func (f *StreamHandlerFunc) OnChunk(chunk types.StreamChunk) error {
	if f.ChunkFunc != nil {
		return f.ChunkFunc(chunk)
	}
	return nil
}

// OnComplete implements StreamHandler.
func (f *StreamHandlerFunc) OnComplete() error {
	if f.CompleteFunc != nil {
		return f.CompleteFunc()
	}
	return nil
}

// OnError implements StreamHandler.
func (f *StreamHandlerFunc) OnError(err error) {
	if f.ErrorFunc != nil {
		f.ErrorFunc(err)
	}
}

// NewStreamHandlerFunc creates a new StreamHandlerFunc.
//
// Any of the parameters can be nil if that callback is not needed.
func NewStreamHandlerFunc(
	chunkFunc func(types.StreamChunk) error,
	completeFunc func() error,
	errorFunc func(error),
) *StreamHandlerFunc {
	return &StreamHandlerFunc{
		ChunkFunc:    chunkFunc,
		CompleteFunc: completeFunc,
		ErrorFunc:    errorFunc,
	}
}

// StreamProcessor processes stream chunks and builds a complete response.
//
// This interface is useful for implementations that need to accumulate
// streaming chunks into a complete response while also performing
// side effects like logging or metrics collection.
type StreamProcessor interface {
	StreamHandler

	// GetResponse returns the accumulated response.
	//
	// This should be called after OnComplete has been called.
	// If called before the stream completes, it may return a partial response.
	GetResponse() *types.ChatResponse

	// Reset resets the processor to its initial state.
	//
	// This allows reusing the same processor for multiple streams.
	Reset()
}

// StreamAdapter converts a channel-based stream to a callback-based stream.
//
// This utility function bridges the gap between channel-based and callback-based
// streaming APIs.
//
// Example:
//
//	stream, err := service.CreateCompletionStream(ctx, req)
//	if err != nil {
//	    return err
//	}
//
//	handler := &MyHandler{}
//	err = StreamAdapter(ctx, stream, handler)
func StreamAdapter(ctx context.Context, stream <-chan types.StreamChunk, handler StreamHandler) error {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			handler.OnError(err)
			return err

		case chunk, ok := <-stream:
			if !ok {
				// Channel closed, stream complete
				return handler.OnComplete()
			}

			if err := handler.OnChunk(chunk); err != nil {
				handler.OnError(err)
				return err
			}
		}
	}
}

// StreamInterceptor intercepts and potentially modifies stream chunks.
//
// This interface is useful for middleware that needs to transform or
// filter stream chunks before they reach the handler.
type StreamInterceptor interface {
	// Intercept is called for each chunk before it's passed to the handler.
	//
	// The interceptor can:
	// - Return the chunk unchanged to pass it through
	// - Return a modified chunk
	// - Return nil to filter out the chunk
	// - Return an error to stop streaming
	//
	// If nil is returned (with no error), the chunk is skipped and the
	// handler's OnChunk is not called for that chunk.
	Intercept(chunk types.StreamChunk) (types.StreamChunk, error)
}

// StreamObserver observes stream chunks without modifying them.
//
// This interface is useful for logging, metrics collection, or other
// side effects that should not affect the stream.
type StreamObserver interface {
	// Observe is called for each chunk.
	//
	// Unlike StreamInterceptor, observers cannot modify or filter chunks.
	// Any errors returned are logged but do not stop the stream.
	Observe(chunk types.StreamChunk)

	// ObserveComplete is called when the stream completes.
	ObserveComplete()

	// ObserveError is called when an error occurs.
	ObserveError(err error)
}

// ChainedStreamHandler chains multiple StreamHandlers together.
//
// All handlers receive all callbacks. If any handler returns an error
// from OnChunk or OnComplete, the chain stops and OnError is called
// on all handlers.
type ChainedStreamHandler struct {
	handlers []StreamHandler
}

// NewChainedStreamHandler creates a new ChainedStreamHandler.
func NewChainedStreamHandler(handlers ...StreamHandler) *ChainedStreamHandler {
	return &ChainedStreamHandler{handlers: handlers}
}

// OnChunk calls OnChunk on all handlers in order.
func (c *ChainedStreamHandler) OnChunk(chunk types.StreamChunk) error {
	for _, handler := range c.handlers {
		if err := handler.OnChunk(chunk); err != nil {
			return err
		}
	}
	return nil
}

// OnComplete calls OnComplete on all handlers in order.
func (c *ChainedStreamHandler) OnComplete() error {
	for _, handler := range c.handlers {
		if err := handler.OnComplete(); err != nil {
			return err
		}
	}
	return nil
}

// OnError calls OnError on all handlers.
func (c *ChainedStreamHandler) OnError(err error) {
	for _, handler := range c.handlers {
		handler.OnError(err)
	}
}

// StreamFilter filters stream chunks based on a predicate.
//
// This is a convenience type for common filtering operations.
type StreamFilter struct {
	handler   StreamHandler
	predicate func(types.StreamChunk) bool
}

// NewStreamFilter creates a new StreamFilter.
//
// The predicate function should return true for chunks that should be
// passed to the handler, and false for chunks that should be filtered out.
func NewStreamFilter(handler StreamHandler, predicate func(types.StreamChunk) bool) *StreamFilter {
	return &StreamFilter{
		handler:   handler,
		predicate: predicate,
	}
}

// OnChunk filters the chunk and calls the handler if the predicate returns true.
func (f *StreamFilter) OnChunk(chunk types.StreamChunk) error {
	if f.predicate(chunk) {
		return f.handler.OnChunk(chunk)
	}
	return nil
}

// OnComplete calls the handler's OnComplete.
func (f *StreamFilter) OnComplete() error {
	return f.handler.OnComplete()
}

// OnError calls the handler's OnError.
func (f *StreamFilter) OnError(err error) {
	f.handler.OnError(err)
}
