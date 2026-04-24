# claude-code-sdk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/redpanda321/claude-code-sdk-go.svg)](https://pkg.go.dev/github.com/redpanda321/claude-code-sdk-go)
[![CI](https://github.com/redpanda321/claude-code-sdk-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/redpanda321/claude-code-sdk-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-00ADD8?logo=go)](./go.mod)

Go SDK port of [claude-code (CCB)](https://github.com/redpanda321/claude-code) — Anthropic's
Claude Code CLI repackaged as a library-first agent toolkit. Sibling of
[`claude-code-sdk-ts`](https://github.com/redpanda321/claude-code-sdk-ts) and
[`claude-code-sdk-python`](https://github.com/redpanda321/claude-code-sdk-python).

## Status

**Alpha.** Public API is unstable until `v1.0.0`. In the `v0.1.0` alpha line only `pkg/tools` and
`pkg/adapters` are populated; the remaining 11 `pkg/*` packages ship as doc-only barrels
with carve-in plans. See [ROADMAP.md](./ROADMAP.md).

## Install

```sh
go get github.com/redpanda321/claude-code-sdk-go@latest
```

Requires Go 1.23+ (needs `iter.Seq2` range-over-func).

## Quickstart

Define a tool by implementing `tools.Tool[In, Out, Progress]`. The streaming surface is
`iter.Seq2[ToolCallProgress[Progress], *ToolResult[Out]]` (CONTEXT D-12):

```go
package main

import (
    "iter"

    "github.com/invopop/jsonschema"
    "github.com/redpanda321/claude-code-sdk-go/pkg/tools"
)

type EchoIn struct {
    Message string `json:"message" jsonschema:"required"`
}
type EchoOut struct{ Echoed string `json:"echoed"` }
type EchoProgress struct{}

type EchoTool struct{}

func (EchoTool) Name() string                            { return "echo" }
func (EchoTool) Description() string                     { return "Echo the input message." }
func (EchoTool) InputSchema() *jsonschema.Schema         { return tools.SchemaFor[EchoIn]() }
func (EchoTool) ValidateInput(_ EchoIn) error            { return nil }
func (EchoTool) CheckPermissions(tools.ToolContext, EchoIn) tools.PermissionDecision {
    return tools.PermissionDecision{Result: tools.PermissionAllow}
}
func (EchoTool) CanUseTool(tools.ToolContext) bool       { return true }
func (EchoTool) IsReadOnly() bool                        { return true }
func (EchoTool) IsEnabled(tools.ToolContext) bool        { return true }
func (EchoTool) RenderResultForAssistant(o EchoOut) string { return o.Echoed }

func (EchoTool) Call(_ tools.ToolContext, in EchoIn) iter.Seq2[tools.ToolCallProgress[EchoProgress], *tools.ToolResult[EchoOut]] {
    return func(yield func(tools.ToolCallProgress[EchoProgress], *tools.ToolResult[EchoOut]) bool) {
        yield(tools.ToolCallProgress[EchoProgress]{}, &tools.ToolResult[EchoOut]{Output: EchoOut{Echoed: in.Message}})
    }
}
```

## Package Layout

| Package | Status |
|---------|--------|
| `pkg/tools` | POPULATED (v0.1) — generic `Tool[In, Out, Progress any]` + supporting types |
| `pkg/adapters` | POPULATED (v0.1) — interface only; no concrete impl in v1 (CONTEXT D-22) |
| `pkg/agent` | doc-only barrel — future plan (v0.2) |
| `pkg/mcp` | doc-only barrel — future plan (v0.3) |
| `pkg/providers` | doc-only barrel — future plan (v0.2 / v0.3) |
| `pkg/hooks` | doc-only barrel — future plan (v0.4) |
| `pkg/plugins` | doc-only barrel — future plan (v0.4) |
| `pkg/skills` | doc-only barrel — future plan (v0.4) |
| `pkg/commands` | doc-only barrel — future plan (v0.4) |
| `pkg/memdir` | doc-only barrel — future plan (v0.5) |
| `pkg/subagents` | doc-only barrel — future plan (v0.5) |
| `pkg/permissions` | doc-only barrel — future plan (v0.5) |
| `pkg/compaction` | doc-only barrel — future plan (v0.5) |

## Sibling SDKs

- [claude-code-sdk-ts](https://github.com/redpanda321/claude-code-sdk-ts) — TypeScript
- [claude-code-sdk-python](https://github.com/redpanda321/claude-code-sdk-python) — Python

## License

MIT — see [LICENSE](./LICENSE).

See [ROADMAP.md](./ROADMAP.md) and [CHANGELOG.md](./CHANGELOG.md).
