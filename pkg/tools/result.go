package tools

// ToolResult is the typed return value of a Tool.Call.
// Out is the tool-specific structured output.
type ToolResult[Out any] struct {
	// Success indicates whether the call completed without error.
	Success bool
	// Output is the structured result (only meaningful when Success is true).
	Output Out
	// Err carries the failure cause when Success is false.
	Err error
	// RenderedForAssistant is the user-visible string the agent loop will
	// splice into the assistant transcript (RenderResultForAssistant output).
	RenderedForAssistant string
}

// ToolCallProgress is an in-flight progress event a Tool may emit.
// Progress is the tool-specific progress payload type.
type ToolCallProgress[Progress any] struct {
	Stage   string
	Payload Progress
}
