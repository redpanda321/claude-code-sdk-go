package tools_test

import (
	"fmt"
	"iter"

	"github.com/invopop/jsonschema"

	"github.com/redpanda321/claude-code-sdk-go/pkg/tools"
)

type exampleEchoIn struct {
	Message string `json:"message" jsonschema:"required"`
}

type exampleEchoOut struct {
	Echoed string `json:"echoed"`
}

type exampleEchoProgress struct{}

type exampleEchoTool struct{}

func (exampleEchoTool) Name() string                    { return "echo" }
func (exampleEchoTool) Description() string             { return "Echo the input message." }
func (exampleEchoTool) InputSchema() *jsonschema.Schema { return tools.SchemaFor[exampleEchoIn]() }
func (exampleEchoTool) ValidateInput(exampleEchoIn) error {
	return nil
}
func (exampleEchoTool) CheckPermissions(tools.ToolContext, exampleEchoIn) tools.PermissionDecision {
	return tools.PermissionDecision{Result: tools.PermissionAllow}
}
func (exampleEchoTool) CanUseTool(tools.ToolContext) bool { return true }
func (exampleEchoTool) IsReadOnly() bool                  { return true }
func (exampleEchoTool) IsEnabled(tools.ToolContext) bool  { return true }
func (exampleEchoTool) RenderResultForAssistant(o exampleEchoOut) string {
	return o.Echoed
}
func (exampleEchoTool) Call(_ tools.ToolContext, in exampleEchoIn) iter.Seq2[tools.ToolCallProgress[exampleEchoProgress], *tools.ToolResult[exampleEchoOut]] {
	return func(yield func(tools.ToolCallProgress[exampleEchoProgress], *tools.ToolResult[exampleEchoOut]) bool) {
		yield(
			tools.ToolCallProgress[exampleEchoProgress]{},
			&tools.ToolResult[exampleEchoOut]{Output: exampleEchoOut{Echoed: in.Message}},
		)
	}
}

var _ tools.Tool[exampleEchoIn, exampleEchoOut, exampleEchoProgress] = exampleEchoTool{}

// Example_tool shows the minimum surface required to implement Tool.
func Example_tool() {
	// See Example_readmeQuickstart for a compile-checked equivalent of the README sample.
	fmt.Println(tools.PermissionAllow)
	// Output: 0
}

// Example_readmeQuickstart mirrors the README sample so doc drift fails tests.
func Example_readmeQuickstart() {
	_ = exampleEchoTool{}
}
