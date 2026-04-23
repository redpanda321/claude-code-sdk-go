# Roadmap — claude-code-sdk-go

Carve-in plan for porting [hanggent/external/claude-code](https://github.com/redpanda321/claude-code) (CCB)
into idiomatic Go packages. Sibling of `claude-code-sdk-ts` and `claude-code-sdk-python`.

## v0.1.0-alpha.1 (current)
- [x] Module skeleton + 13 `pkg/*` barrels
- [x] `pkg/tools`: generic Tool[In,Out,Progress] interface + supporting types
- [x] `pkg/adapters`: generic Adapter interface (interface-only per CONTEXT D-22)
- [x] CI (matrix 1.23 + 1.24) + release workflow
- [x] Docs (README / CONTRIBUTING / SECURITY / CODE_OF_CONDUCT / CHANGELOG)

## v0.2.0-alpha (next)
- [ ] `pkg/agent`: agentic loop returning `iter.Seq2[Event, error]`
- [ ] `pkg/providers/anthropic`: Anthropic provider with custom SSE reader
- [ ] First end-to-end smoke test (echo tool + Anthropic provider)

## v0.3.0-alpha
- [ ] `pkg/providers/{bedrock,vertex,openai,glm}`
- [ ] `pkg/mcp`: MCP client built on `modelcontextprotocol/go-sdk`

## v0.4.0-alpha
- [ ] `pkg/{hooks,plugins,skills,commands}`: extensibility surfaces

## v0.5.0-alpha
- [ ] `pkg/{memdir,subagents,permissions,compaction}`: remaining barrels

## v0.6.0-beta
- [ ] First concrete adapter (separate module — `claude-code-sdk-go-langchaingo`)
- [ ] API stability review

## v1.0.0
- [ ] Public API frozen
- [ ] Sibling-CLI (`ccb-go`) repo announced

## Out of scope (per CONTEXT deferred section)
- Go CLI / TUI binary in this repo (separate `ccb-go` repo, future)
- LiteLLM / langchaingo / genkit dependencies in core
- Windows / macOS CI runners
- Sync wrapper (iterators are blocking-friendly)