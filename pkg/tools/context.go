package tools

import (
	"context"
	"log/slog"
)

// ToolContext is the per-call execution context passed to a Tool.Call.
// Mirrors CCB ToolContext from src/Tool.ts.
type ToolContext struct {
	// SessionID identifies the agent session this call belongs to.
	SessionID string
	// CWD is the working directory the tool should treat as the project root.
	CWD string
	// Ctx propagates cancellation. Tools MUST honor Ctx.Done().
	Ctx context.Context
	// Logger is a structured logger; nil-safe (use slog.Default() if nil).
	Logger *slog.Logger
}
