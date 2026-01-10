package types

import "encoding/json"

// ToolType represents the type of tool.
type ToolType string

const (
	// ToolTypeFunction represents a function tool.
	ToolTypeFunction ToolType = "function"
)

// ToolChoice represents how the model should use tools.
type ToolChoice string

const (
	// ToolChoiceNone means the model will not call any tools.
	ToolChoiceNone ToolChoice = "none"

	// ToolChoiceAuto means the model can choose to call tools.
	ToolChoiceAuto ToolChoice = "auto"

	// ToolChoiceRequired means the model must call at least one tool.
	ToolChoiceRequired ToolChoice = "required"

	// ToolChoiceAny means the model must call a tool (alias for Required).
	ToolChoiceAny ToolChoice = "any"
)

// String returns the string representation of the ToolChoice.
func (t ToolChoice) String() string {
	return string(t)
}

// ToolDefinition defines a tool that can be called by the model.
type ToolDefinition struct {
	// Type is the type of tool. Currently only "function" is supported.
	Type ToolType `json:"type"`

	// Function describes the function to call.
	Function FunctionDefinition `json:"function"`
}

// FunctionDefinition defines a function that can be called.
type FunctionDefinition struct {
	// Name is the name of the function.
	Name string `json:"name"`

	// Description explains what the function does.
	// This helps the model decide when to call it.
	Description string `json:"description,omitempty"`

	// Parameters defines the function parameters as JSON Schema.
	Parameters interface{} `json:"parameters"`

	// Strict enables strict schema adherence (if supported by provider).
	// When true, the model's output must exactly match the schema.
	Strict bool `json:"strict,omitempty"`
}

// NewFunctionDefinition creates a new FunctionDefinition with the given parameters.
func NewFunctionDefinition(name, description string, parameters interface{}) *FunctionDefinition {
	return &FunctionDefinition{
		Name:        name,
		Description: description,
		Parameters:  parameters,
	}
}

// ToolCall represents a call to a tool/function by the model.
type ToolCall struct {
	// ID is a unique identifier for this tool call.
	// Used to match tool responses back to the call.
	ID string `json:"id"`

	// Type is the type of tool being called.
	Type ToolType `json:"type"`

	// Function contains the function call details.
	Function FunctionCall `json:"function"`

	// Index is the index of this tool call in the list (for streaming).
	Index int `json:"index,omitempty"`
}

// FunctionCall represents a function call made by the model.
type FunctionCall struct {
	// Name is the name of the function to call.
	Name string `json:"name"`

	// Arguments is a JSON string containing the function arguments.
	Arguments string `json:"arguments"`
}

// ParseArguments parses the JSON arguments into the provided struct.
func (f *FunctionCall) ParseArguments(v interface{}) error {
	return json.Unmarshal([]byte(f.Arguments), v)
}

// GetArgumentsMap returns the arguments as a map.
func (f *FunctionCall) GetArgumentsMap() (map[string]interface{}, error) {
	var args map[string]interface{}
	err := json.Unmarshal([]byte(f.Arguments), &args)
	return args, err
}

// ToolCallFunction is a helper to create a tool call function.
func ToolCallFunction(id, name, arguments string) *ToolCall {
	return &ToolCall{
		ID:   id,
		Type: ToolTypeFunction,
		Function: FunctionCall{
			Name:      name,
			Arguments: arguments,
		},
	}
}

// ResponseFormat specifies the format of the model's output.
type ResponseFormat struct {
	// Type is the format type. Common values: "text", "json_object", "json_schema".
	Type string `json:"type"`

	// JSONSchema is the JSON schema for the response (if Type is "json_schema").
	JSONSchema interface{} `json:"json_schema,omitempty"`
}

// NewTextResponseFormat creates a response format for plain text.
func NewTextResponseFormat() *ResponseFormat {
	return &ResponseFormat{Type: "text"}
}

// NewJSONResponseFormat creates a response format for JSON output.
func NewJSONResponseFormat() *ResponseFormat {
	return &ResponseFormat{Type: "json_object"}
}

// NewJSONSchemaResponseFormat creates a response format with a JSON schema.
func NewJSONSchemaResponseFormat(schema interface{}) *ResponseFormat {
	return &ResponseFormat{
		Type:       "json_schema",
		JSONSchema: schema,
	}
}

// JSONSchema represents a JSON Schema for function parameters.
// This is a flexible structure that can represent any JSON Schema.
type JSONSchema struct {
	// Type is the JSON type (e.g., "object", "string", "number").
	Type string `json:"type"`

	// Description describes the schema.
	Description string `json:"description,omitempty"`

	// Properties defines the properties for object types.
	Properties map[string]*JSONSchema `json:"properties,omitempty"`

	// Required lists required property names for object types.
	Required []string `json:"required,omitempty"`

	// Items defines the schema for array items.
	Items *JSONSchema `json:"items,omitempty"`

	// Enum lists allowed values for enum types.
	Enum []interface{} `json:"enum,omitempty"`

	// AdditionalProperties controls whether additional properties are allowed.
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
}

// NewObjectSchema creates a new object schema.
func NewObjectSchema(description string, properties map[string]*JSONSchema, required []string) *JSONSchema {
	return &JSONSchema{
		Type:        "object",
		Description: description,
		Properties:  properties,
		Required:    required,
	}
}

// NewStringSchema creates a new string schema.
func NewStringSchema(description string) *JSONSchema {
	return &JSONSchema{
		Type:        "string",
		Description: description,
	}
}

// NewNumberSchema creates a new number schema.
func NewNumberSchema(description string) *JSONSchema {
	return &JSONSchema{
		Type:        "number",
		Description: description,
	}
}

// NewBooleanSchema creates a new boolean schema.
func NewBooleanSchema(description string) *JSONSchema {
	return &JSONSchema{
		Type:        "boolean",
		Description: description,
	}
}

// NewArraySchema creates a new array schema.
func NewArraySchema(description string, items *JSONSchema) *JSONSchema {
	return &JSONSchema{
		Type:        "array",
		Description: description,
		Items:       items,
	}
}

// NewEnumSchema creates a new enum schema.
func NewEnumSchema(description string, enumValues []interface{}) *JSONSchema {
	return &JSONSchema{
		Type:        "string",
		Description: description,
		Enum:        enumValues,
	}
}
