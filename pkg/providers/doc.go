// Package providers is the carve-in home for LLM provider integrations.
//
// Ports CCB src/services/api/claude.ts and sibling provider files. Covers
// Anthropic, Bedrock, Vertex, OpenAI-compatible, and GLM. Uses stdlib net/http
// + a custom SSE reader per Phase 999.3 CONTEXT D-18 and D-19 (no litellm
// dependency).
//
// Status: doc-only barrel in v0.1.0-alpha.1. Populated in v0.2/v0.3-alpha.
package providers
