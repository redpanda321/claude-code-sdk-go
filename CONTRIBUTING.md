# Contributing to claude-code-sdk-go

Thanks for your interest! This SDK is in **alpha**; expect churn until `v1.0.0`.

## Prerequisites

- **Go 1.23+** (module floor; CI also tests on 1.24). `iter.Seq2` range-over-func is required.
- **golangci-lint v1.60+** — install from <https://golangci-lint.run/welcome/install/>.

## Dev loop

From the repo root:

```sh
go vet ./...
golangci-lint run
go test ./... -race -count=1
```

All three must pass before you open a PR. CI runs the same commands on ubuntu-latest.

## Fork + branch + PR flow

1. Fork `github.com/redpanda321/claude-code-sdk-go` to your GitHub account.
2. Clone your fork, create a topic branch:
   ```sh
   git checkout -b feat/my-change
   ```
3. Make your change with tests. Keep commits small and atomic.
4. Push and open a PR against `main`. Link any related issue.

## Commit conventions

We use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` — new public API or behavior
- `fix:` — bug fix in existing behavior
- `docs:` — README / CHANGELOG / godoc
- `test:` — adding or refactoring tests only
- `ci:` — workflows, linters, dependabot
- `chore:` — tooling, deps, non-user-facing housekeeping

Scopes are optional (e.g. `feat(tools): add RenderResult helper`).

## PR checklist

Before requesting review, confirm:

- [ ] Tests added or updated for the changed behavior.
- [ ] `go test ./... -race -count=1` passes locally.
- [ ] `golangci-lint run` is clean (no new findings).
- [ ] `go vet ./...` is clean.
- [ ] `CHANGELOG.md` updated under `## [Unreleased]` with a one-line user-facing note.
- [ ] Public API changes include godoc comments on every exported symbol.
- [ ] No new direct dependencies added without justification in the PR description.

## Code style

- `gofmt` / `goimports` clean (CI enforces).
- Prefer interfaces over concrete types for public APIs.
- Stream via `iter.Seq2` (CONTEXT D-12) instead of channels for tool/agent loops.
- Errors: wrap with `%w`; expose sentinel errors (`ErrXxx`) for callers to `errors.Is` against.

## Questions

Open a [Discussion](https://github.com/redpanda321/claude-code-sdk-go/discussions) for
design questions, or an Issue for reproducible bugs. For security, see
[SECURITY.md](./SECURITY.md).
