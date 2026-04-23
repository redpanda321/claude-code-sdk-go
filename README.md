# claude-code-sdk-go

Go SDK port of [claude-code (CCB)](https://github.com/redpanda321/claude-code) — Anthropic's
Claude Code CLI repackaged as a library-first agent toolkit. Sibling of
[`claude-code-sdk-ts`](https://github.com/redpanda321/claude-code-sdk-ts) and
[`claude-code-sdk-python`](https://github.com/redpanda321/claude-code-sdk-python).

## Status

**Alpha.** Public API is unstable until `v1.0.0`. In `v0.1.0-alpha.1` only `pkg/tools` and
`pkg/adapters` are populated; the remaining 11 `pkg/*` packages ship as doc-only barrels
with carve-in plans. See [ROADMAP.md](./ROADMAP.md).

## Install

```sh
go get github.com/redpanda321/claude-code-sdk-go@latest
```

Requires Go 1.23+ (needs `iter.Seq2` range-over-func).

## Package Layout

| Package | Status |
|---------|--------|
| `pkg/tools` | populated (v0.1) — generic `Tool[In, Out, Progress any]` + supporting types |
| `pkg/adapters` | populated (v0.1) — interface only; no concrete impl in v1 |
| `pkg/agent` | doc-only barrel |
| `pkg/mcp` | doc-only barrel |
| `pkg/providers` | doc-only barrel |
| `pkg/hooks` | doc-only barrel |
| `pkg/plugins` | doc-only barrel |
| `pkg/skills` | doc-only barrel |
| `pkg/commands` | doc-only barrel |
| `pkg/memdir` | doc-only barrel |
| `pkg/subagents` | doc-only barrel |
| `pkg/permissions` | doc-only barrel |
| `pkg/compaction` | doc-only barrel |

## Sibling SDKs

- [claude-code-sdk-ts](https://github.com/redpanda321/claude-code-sdk-ts) — TypeScript
- [claude-code-sdk-python](https://github.com/redpanda321/claude-code-sdk-python) — Python

## License

MIT — see [LICENSE](./LICENSE).

See [ROADMAP.md](./ROADMAP.md) and [CHANGELOG.md](./CHANGELOG.md).
