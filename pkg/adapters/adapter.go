package adapters

import "errors"

// FrameworkInfo identifies a third-party Go agent framework.
type FrameworkInfo struct {
	// Name is the canonical framework name (e.g., "langchaingo", "genkit",
	// "claude-agent-sdk-go").
	Name string
	// ImportPath is the Go module path (e.g., "github.com/tmc/langchaingo").
	ImportPath string
	// Version is a free-form version string the adapter supports.
	Version string
}

// Adapter wraps the claude-code-sdk-go public surface (tools, agent loop,
// providers) into a third-party framework's native shape and back.
//
// Implementations live in their own modules / packages — claude-code-sdk-go
// ships the interface only (per Phase 999.3 CONTEXT D-22). A concrete
// langchaingo adapter (or any other) lands in a follow-up phase to keep
// this alpha decoupled from external framework API churn.
//
// Implementations MUST be safe for concurrent use across goroutines.
type Adapter interface {
	// Framework returns metadata about the target framework.
	Framework() FrameworkInfo

	// Wrap takes a value from the claude-code-sdk-go public surface
	// (typically a tools.Tool implementation, agent.Agent, or provider)
	// and returns the framework-native equivalent. Returns an error
	// describing the unsupported source type when src cannot be wrapped.
	Wrap(src any) (any, error)

	// Unwrap is the inverse of Wrap: takes a framework-native value and
	// returns the claude-code-sdk-go-native equivalent. Returns an error
	// describing the unsupported framework type when src cannot be unwrapped.
	Unwrap(src any) (any, error)
}

// ErrNoConcreteAdapter is returned by helpers that look up concrete adapter
// implementations. claude-code-sdk-go v1 ships zero concrete adapters
// (CONTEXT D-22) — consumers must register their own via a sibling module.
var ErrNoConcreteAdapter = errors.New(
	"claude-code-sdk-go: no concrete Adapter shipped in v1 — " +
		"register a third-party adapter (e.g. claude-code-sdk-go-langchaingo) or implement pkg/adapters.Adapter directly",
)
