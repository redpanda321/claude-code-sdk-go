package tools_test

import (
	"fmt"

	"github.com/redpanda321/claude-code-sdk-go/pkg/tools"
)

// Example_tool shows the minimum surface required to implement Tool.
func Example_tool() {
	// See pkg/tools/tool_test.go for a fully runnable example.
	// The shape is: implement all 10 methods of Tool[In, Out, Progress],
	// returning iter.Seq2 from Call.
	fmt.Println(tools.PermissionAllow)
	// Output: 0
}
