package tools

import (
	"context"
	"iter"
	"testing"

	"github.com/invopop/jsonschema"
)

// echoIn / echoOut / echoProgress drive a minimal Tool implementation
// used to (a) assert interface satisfaction at compile time and
// (b) exercise Call's iter.Seq2 streaming surface.
type echoIn struct {
	Message string `json:"message" jsonschema:"required"`
}
type echoOut struct {
	Echoed string `json:"echoed"`
}
type echoProgress struct {
	Step int `json:"step"`
}

type echoTool struct{}

func (echoTool) Name() string                    { return "echo" }
func (echoTool) Description() string             { return "Echoes its input back." }
func (echoTool) InputSchema() *jsonschema.Schema { return SchemaFor[echoIn]() }
func (echoTool) ValidateInput(in echoIn) error   { return nil }
func (echoTool) CheckPermissions(ctx ToolContext, in echoIn) PermissionDecision {
	return PermissionDecision{Result: PermissionAllow, Reason: "echo is read-only"}
}
func (echoTool) CanUseTool(ctx ToolContext) bool             { return true }
func (echoTool) RenderResultForAssistant(out echoOut) string { return out.Echoed }
func (echoTool) IsReadOnly() bool                            { return true }
func (echoTool) IsEnabled(ctx ToolContext) bool              { return true }

func (echoTool) Call(ctx ToolContext, in echoIn) iter.Seq2[ToolCallProgress[echoProgress], *ToolResult[echoOut]] {
	return func(yield func(ToolCallProgress[echoProgress], *ToolResult[echoOut]) bool) {
		// emit one progress, then the final result
		if !yield(ToolCallProgress[echoProgress]{Stage: "starting", Payload: echoProgress{Step: 1}}, nil) {
			return
		}
		yield(
			ToolCallProgress[echoProgress]{},
			&ToolResult[echoOut]{
				Success:              true,
				Output:               echoOut{Echoed: in.Message},
				RenderedForAssistant: in.Message,
			},
		)
	}
}

// Compile-time assertion that echoTool satisfies the generic Tool interface.
var _ Tool[echoIn, echoOut, echoProgress] = echoTool{}

func TestEchoTool_Name(t *testing.T) {
	if got := (echoTool{}).Name(); got != "echo" {
		t.Fatalf("Name = %q, want %q", got, "echo")
	}
}

func TestEchoTool_Permissions(t *testing.T) {
	d := (echoTool{}).CheckPermissions(ToolContext{}, echoIn{})
	if d.Result != PermissionAllow {
		t.Fatalf("CheckPermissions Result = %v, want PermissionAllow", d.Result)
	}
}

func TestEchoTool_CallStream(t *testing.T) {
	ctx := ToolContext{Ctx: context.Background()}
	var (
		progressCount int
		finalResult   *ToolResult[echoOut]
	)
	for prog, res := range (echoTool{}).Call(ctx, echoIn{Message: "hi"}) {
		if res != nil {
			finalResult = res
			continue
		}
		_ = prog
		progressCount++
	}
	if progressCount != 1 {
		t.Fatalf("progressCount = %d, want 1", progressCount)
	}
	if finalResult == nil || !finalResult.Success || finalResult.Output.Echoed != "hi" {
		t.Fatalf("finalResult = %+v, want Success && Echoed=hi", finalResult)
	}
}
