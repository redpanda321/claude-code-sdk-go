// Package adapters defines the contract third-party Go agent frameworks
// (langchaingo, genkit, future claude-agent-sdk-go, etc.) implement to
// wrap the claude-code-sdk-go public surface into their native shapes.
//
// Status (Phase 999.3 v1): INTERFACE ONLY.
//
// Per CONTEXT D-22, claude-code-sdk-go v1 ships no concrete adapter
// implementation. The user-facing decision was to NOT couple the alpha
// SDK to any external framework's API churn. Concrete adapters land in
// follow-up phases as separate modules:
//
//   - claude-code-sdk-go-langchaingo (planned)
//   - claude-code-sdk-go-genkit      (planned)
//   - claude-code-sdk-go-anthropic   (when an official claude-agent-sdk-go appears)
//
// Helpers that need a concrete adapter return ErrNoConcreteAdapter.
package adapters
