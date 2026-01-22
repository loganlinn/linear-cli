# CLAUDE.md

Linear CLI - A token-efficient CLI for Linear, written in Go.

## Commands

```bash
make build    # Build binary (bin/linear)
make test     # Run unit tests
make clean    # Clean build artifacts
```

## Project Structure

```
cmd/linear/          # CLI entry point
internal/cli/        # CLI commands (Cobra)
internal/format/     # ASCII formatters for token-efficient output
internal/linear/     # Linear GraphQL client
internal/service/    # Service layer
internal/skills/     # Embedded Claude Code skills
internal/oauth/      # OAuth2 flow
internal/token/      # Secure token storage
```

## For AI Agents (Claude Code)

### Essential Setup

**ALWAYS run this first** before any other Linear commands:
```bash
linear init  # Select default team - required for cycle operations
```

This creates `.linear.yaml` with your default team. Without this, many commands will fail with "team is required".

### Common Patterns

#### Listing Issues
```bash
# List all issues (NOT just assigned - this is the default behavior)
linear issues list --format full

# Filter by state
linear issues list --team CEN --state "In Progress" --format full

# Filter by priority (1=urgent, 2=high, 3=normal, 4=low)
linear issues list --priority 1 --format full

# Get issues in specific cycle
linear issues list --cycle 65 --format full

# Filter by assignee
linear issues list --assignee me --format full

# Combine filters
linear issues list --state Backlog --labels customer --priority 1 --format full
```

#### Cycle Operations
```bash
# CRITICAL: Analyze velocity BEFORE planning cycles
linear cycles analyze --team CEN --count 10

# Get cycle details (by number - requires 'linear init')
linear cycles get 65

# List all cycles
linear cycles list --team CEN --format full

# List only active cycles
linear cycles list --active
```

#### Creating Issues
```bash
# Minimal - requires 'linear init' first
linear issues create "Fix authentication bug"

# Complete with all parameters
linear issues create "Implement feature" \
  --team CEN \
  --state "In Progress" \
  --priority 1 \
  --assignee me \
  --estimate 5 \
  --cycle 65 \
  --labels "backend,security"
```

#### Working with Team Context
- Cycle numbers (65, 66) require team context from `linear init`
- Issue identifiers (CEN-123) work without team context
- UUIDs work without team context

### Error Handling

| Error | Solution |
|-------|----------|
| `--team is required` | Run `linear init` or add `--team CEN` flag |
| `Entity not found: Cycle` | Verify team context: run `linear init` |
| `Unknown type "IssueOrderBy"` | Update CLI to latest version |

### Skills Usage

After running `linear init`, use these skills:
- `/cycle-plan` - Analyze and plan cycles
- `/triage` - Prioritize backlog issues
- `/deps` - Visualize dependency graphs

## CLI Commands

### Dependency Graph
```bash
linear deps ENG-100          # Show deps for issue
linear deps --team ENG       # Show all deps for team
```

### Skills Management
```bash
linear skills list           # List available skills
linear skills install --all  # Install all skills
linear skills install prd    # Install specific skill
```

Available skills: `/prd`, `/triage`, `/cycle-plan`, `/retro`, `/deps`

## Key Design Decisions

- **ASCII output** - Token-efficient, no JSON overhead
- **Human-readable IDs** - "TEST-123" not UUIDs
- **Service layer** - Validation and formatting abstraction

## Testing

```bash
go test -v ./internal/linear -run TestCreateIssue
go test -cover ./...
```

## Session Completion

1. `make test` must pass
2. Commit with clear messages
3. Push to remote - work is NOT complete until pushed
