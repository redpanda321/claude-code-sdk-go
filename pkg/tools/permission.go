package tools

// PermissionResult is the outcome of a permission check.
type PermissionResult int

const (
	// PermissionAllow lets the tool call proceed.
	PermissionAllow PermissionResult = iota
	// PermissionDeny blocks the tool call.
	PermissionDeny
	// PermissionPrompt requires user confirmation before proceeding.
	PermissionPrompt
)

// PermissionDecision is a permission check outcome with a human-readable reason.
type PermissionDecision struct {
	Result PermissionResult
	Reason string
}
