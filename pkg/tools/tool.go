package tools

import (
	"iter"

	"github.com/invopop/jsonschema"
)

// Tool is the generic contract every SDK tool implements.
//
// Type parameters:
//   - In:       structured input (must be JSON-marshalable; SchemaFor[In] derives the wire schema)
//   - Out:      structured output type (returned inside ToolResult[Out])
//   - Progress: per-call progress payload type (emitted via Call's iterator)
//
// Mirrors CCB src/Tool.ts Tool<Input, Output, Progress>.
type Tool[In any, Out any, Progress any] interface {
	// Name returns the canonical tool identifier (snake_case).
	Name() string

	// Description returns a one-paragraph summary the model sees.
	Description() string

	// InputSchema returns the JSON Schema for In.
	// Default implementation: return SchemaFor[In]().
	InputSchema() *jsonschema.Schema

	// ValidateInput validates a decoded In before Call. Returning a non-nil
	// error short-circuits the call with a validation failure.
	ValidateInput(in In) error

	// CheckPermissions evaluates whether this call is permitted in the
	// current context. Returns an Allow / Deny / Prompt decision.
	CheckPermissions(ctx ToolContext, in In) PermissionDecision

	// CanUseTool is a coarse gate the agent loop calls before any
	// permission check; e.g. honors --read-only mode.
	CanUseTool(ctx ToolContext) bool

	// Call executes the tool. Returns an iterator that yields zero or more
	// progress events followed by exactly one final ToolResult[Out].
	// The iterator MUST honor ctx.Ctx cancellation.
	//
	// Per CONTEXT D-12, iter.Seq2 is the public streaming surface.
	Call(ctx ToolContext, in In) iter.Seq2[ToolCallProgress[Progress], *ToolResult[Out]]

	// RenderResultForAssistant formats Out as the assistant-visible string
	// spliced into the transcript. Pure function (no side effects).
	RenderResultForAssistant(out Out) string

	// IsReadOnly returns true if this tool never mutates filesystem,
	// network, or external state. Used by --read-only and permission gates.
	IsReadOnly() bool

	// IsEnabled returns true if this tool is enabled in the current
	// configuration; agent loop skips disabled tools when listing.
	IsEnabled(ctx ToolContext) bool
}
