# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0-alpha.1] - 2026-04-23

### Added
- Initial alpha release of `claude-code-sdk-go`.
- Module path: `github.com/redpanda321/claude-code-sdk-go`.
- Standard Go layout: `cmd/`, `internal/`, `pkg/` (CONTEXT D-08).
- 13 public package barrels under `pkg/*` with carve-in plan docstrings.
- `pkg/tools`: generic `Tool[In, Out, Progress any]` interface, supporting types
  (`ToolContext`, `ToolResult`, `ToolCallProgress`, `PermissionResult`,
  `PermissionDecision`), and `SchemaFor[T any]()` helper using
  `github.com/invopop/jsonschema`.
- `pkg/adapters`: generic `Adapter` interface, `FrameworkInfo`, `ErrNoConcreteAdapter`
  sentinel. **No concrete adapter ships in v1** (CONTEXT D-22).
- Streaming surface: `iter.Seq2` (Go 1.23+ range-over-func iterators, CONTEXT D-12).
- CI: GitHub Actions matrix Go 1.23 + 1.24 on ubuntu-latest, `golangci-lint`
  with errcheck/gofmt/goimports/gosimple/govet/ineffassign/staticcheck/unused.
- Release: tag-triggered GitHub Release workflow.
- Dependabot: weekly `gomod` + `github-actions` updates.
- Docs: README, CONTRIBUTING, SECURITY, CODE_OF_CONDUCT, ROADMAP.

### Status
- Public API is **unstable** during alpha. Stability targeted for `v1.0.0`.
- 11 of 13 `pkg/*` packages are doc-only barrels — see [ROADMAP.md](./ROADMAP.md).

[Unreleased]: https://github.com/redpanda321/claude-code-sdk-go/compare/v0.1.0-alpha.1...HEAD
[0.1.0-alpha.1]: https://github.com/redpanda321/claude-code-sdk-go/releases/tag/v0.1.0-alpha.1
