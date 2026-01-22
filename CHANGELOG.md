# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.1] - 2026-01-21

### Added
- `--version` flag to display version information
- Version is auto-generated from git tags during build
- Comprehensive OAuth setup documentation for agent mode
- Specific callback URL guidance (`http://localhost:3000/oauth-callback` for agents)

### Fixed
- **Critical**: GraphQL pagination query error (removed invalid `totalCount` field from IssueConnection)
- Users can now run `linear issues list` without HTTP 400 errors

### Changed
- Removed "MCP server" references from all user-facing text
- Updated authentication prompts to say "automation, bots" instead of "MCP server"
- Changed onboard command title from "Light Linear MCP" to "Linear CLI"
- Updated README with correct Homebrew installation instructions
- Updated manual download URLs to use tar.gz archives

## [1.0.0] - 2026-01-20

### Initial Release

A token-efficient CLI for Linear.

#### Features

**CLI Commands:**
- `linear auth` - OAuth2 authentication (login, logout, status)
- `linear issues` - Full issue management (list, get, create, update, comment, reply, react)
- `linear projects` - Project management (list, get, create, update)
- `linear cycles` - Cycle management and velocity analytics
- `linear teams` - Team info, labels, workflow states
- `linear users` - User listing and lookup
- `linear deps` - Dependency graph visualization
- `linear skills` - Claude Code skill installation

**Key Capabilities:**
- Human-readable identifiers (TEST-123, not UUIDs)
- Token-efficient ASCII output format
- OAuth2 authentication with secure local storage
- Cycle velocity analytics with recommendations
- Comment threading and emoji reactions
- Issue dependency tracking

#### Platform Support

- macOS (Apple Silicon and Intel)
- Linux (64-bit)
- Windows (64-bit)

[1.0.1]: https://github.com/joa23/linear-cli/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/joa23/linear-cli/releases/tag/v1.0.0
