# Security Policy

## Supported versions

During the `0.x` alpha series, **only the latest released minor** receives security fixes.
After `v1.0.0`, this policy will be revisited.

| Version | Supported |
|---------|-----------|
| `0.1.x` (current alpha) | :white_check_mark: |
| `< 0.1.0` | :x: |

## Reporting a vulnerability

**Do not report security issues via public GitHub Issues, Discussions, or pull requests.**

Please report privately via GitHub Security Advisories:

<https://github.com/redpanda321/claude-code-sdk-go/security/advisories/new>

Include, if possible:

- A description of the issue and its impact.
- Steps to reproduce, or a minimal proof-of-concept.
- Affected versions / commit SHAs.
- Your suggested fix or mitigation, if you have one.

## Response SLA

During alpha, we aim for **best-effort triage within 7 days** of receiving a report.
We will acknowledge receipt, investigate, and coordinate disclosure with you before
publishing any advisory.

## Scope

This policy covers code shipped under the `github.com/redpanda321/claude-code-sdk-go`
Go module. Issues in direct dependencies (e.g. `github.com/invopop/jsonschema`) should
be reported to those upstreams; we will track and pin fixes via Dependabot.
