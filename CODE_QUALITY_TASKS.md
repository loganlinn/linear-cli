# Code Quality Improvement Tasks

**Status:** 📋 Awaiting Review
**Created:** 2026-01-24
**Reviewers:** Gemini AI + Claude Code + golangci-lint

This document consolidates findings from three sources:
1. **Gemini AI Code Review** - Architectural & design issues
2. **Claude Code Exploration** - Consistency & pattern analysis
3. **golangci-lint** - Static analysis & best practices
4. **Test Coverage Analysis** - Identify undertested code

---

## 🚨 CRITICAL Priority (Blocks Production)

### C1. Fix Priority 0 Filter Bug (CRITICAL)
**Source:** Gemini Review
**File:** `internal/cli/issues.go:132`
**Issue:** Cannot search for issues with "No Priority" (priority=0)

**Problem:**
```go
if priority > 0 {
    filters.Priority = &priority
}
```

When user runs `linear issues list --priority 0`, the filter is skipped because `priority > 0` is false.

**Impact:** Users cannot filter unprioritized issues.

**Fix:**
Use `cmd.Flags().Changed("priority")` to distinguish default zero from user-provided zero:
```go
if cmd.Flags().Changed("priority") {
    filters.Priority = &priority
}
```

---

### C2. Refactor for Testability - Dependency Injection
**Source:** Gemini Review (CRITICAL)
**Files:** All command files
**Issue:** Commands are completely untestable without real Linear API

**Problem:**
```go
RunE: func(cmd *cobra.Command, args []string) error {
    client, err := getLinearClient()  // ← Requires real token file
    svc, err := getIssueService()      // ← Hits real API
    // Cannot mock, cannot unit test
}
```

**Impact:**
- 12.5% CLI test coverage (should be >80%)
- 1.3% service layer coverage (should be >90%)
- Cannot test error paths
- Requires Live API for every test

**Fix (Phased Approach):**

**Phase 1:** Create service interfaces
```go
// internal/cli/dependencies.go
type Dependencies struct {
    IssueService    IssueServiceInterface
    ProjectService  ProjectServiceInterface
    CycleService    CycleServiceInterface
    // ...
}

type IssueServiceInterface interface {
    Search(filters *service.SearchFilters) ([]Issue, error)
    Get(id string) (*Issue, error)
    Create(input *service.CreateIssueInput) (*Issue, error)
    // ...
}
```

**Phase 2:** Inject dependencies into commands
```go
func newIssuesListCmd(deps *Dependencies) *cobra.Command {
    return &cobra.Command{
        RunE: func(cmd *cobra.Command, args []string) error {
            issues, err := deps.IssueService.Search(filters)
            // ...
        },
    }
}
```

**Phase 3:** Create test mocks
```go
// internal/cli/issues_test.go
type mockIssueService struct{}
func (m *mockIssueService) Search(filters *service.SearchFilters) ([]Issue, error) {
    return []Issue{{ID: "test"}}, nil
}
```

**Estimated Effort:** 2-3 days
**Coverage Impact:** CLI: 12.5% → 75%+

---

## 🔥 HIGH Priority (Quality & Maintainability)

### H1. Extract Repeated Attachment Upload Logic
**Source:** Gemini Review
**Files:** `internal/cli/issues.go` (4 locations)
**Issue:** 60 lines of identical code copy-pasted 4 times

**Locations:**
- Lines 225-240 (Create)
- Lines 405-420 (Update)
- Lines 565-580 (Comment)
- Lines 665-680 (Reply)

**Fix:**
```go
// internal/cli/helpers.go
func uploadAndAppendAttachments(client *linear.Client, body string, filePaths []string) (string, error) {
    if len(filePaths) == 0 {
        return body, nil
    }

    for _, filePath := range filePaths {
        assetURL, err := client.Attachments.UploadFileFromPath(filePath)
        if err != nil {
            return "", fmt.Errorf("failed to upload %s: %w", filePath, err)
        }
        if body != "" {
            body += "\n\n"
        }
        body += fmt.Sprintf("![%s](%s)", filepath.Base(filePath), assetURL)
    }
    return body, nil
}
```

**Usage:**
```go
desc, err := uploadAndAppendAttachments(client, desc, attachFiles)
if err != nil {
    return err
}
```

**Estimated Effort:** 2 hours

---

### H2. Standardize Limit Validation
**Source:** Gemini Review + Claude Exploration
**Files:** `issues.go`, `cycles.go`, `projects.go`, `search.go`
**Issue:** Inconsistent defaults (10 vs 25) and validation

**Current State:**
- `issues.go:109` - Default: 10, validates > 250 ✓
- `cycles.go:79` - Default: 25, NO validation ✗
- `projects.go:55` - Default: 25, NO validation ✗
- `search.go:177` - Default: 10, validates > 250 ✓

**Fix:**
```go
// internal/cli/helpers.go
const (
    DefaultLimit = 25
    MaxLimit     = 250
)

func validateAndNormalizeLimit(limit int) (int, error) {
    if limit <= 0 {
        return DefaultLimit, nil
    }
    if limit > MaxLimit {
        return 0, fmt.Errorf("--limit cannot exceed %d (Linear API maximum), got %d", MaxLimit, limit)
    }
    return limit, nil
}
```

**Usage:**
```go
limit, err := validateAndNormalizeLimit(limit)
if err != nil {
    return err
}
```

**Estimated Effort:** 1 hour

---

### H3. Standardize Team Flag Descriptions
**Source:** Claude Exploration
**Files:** All command files (9 locations)
**Issue:** 10 different descriptions for the same `--team` flag

**Current Variations:**
| Command | Description |
|---------|-------------|
| `issues list` | "Team ID or key (uses .linear.yaml default)" |
| `issues create` | "Team name or key (uses .linear.yaml default if not specified)" |
| `issues update` | "Team context for cycle resolution (optional, uses .linear.yaml or issue identifier)" |
| `cycles list` | "Team ID or key (uses .linear.yaml default)" |
| `projects list` | "Team key or ID (uses .linear.yaml if not specified)" |
| `search` | "Filter by team (uses .linear.yaml default)" |
| `deps` | "Show dependencies for all issues in team" |

**Fix:**
```go
// internal/cli/flags.go
const TeamFlagDescription = "Team ID or key (uses .linear.yaml default)"

// Usage in all commands:
cmd.Flags().StringVarP(&teamID, "team", "t", "", TeamFlagDescription)
```

**Estimated Effort:** 30 minutes

---

### H4. Standardize Error Messages
**Source:** Claude Exploration
**Files:** `issues.go`, `cycles.go`, `projects.go`
**Issue:** Inconsistent error messages for same condition

**Current Variations:**
- `"--team is required (or run 'linear init' to set a default)"`
- `"team is required. Use --team or run 'linear init' to set a default"`

**Fix:**
```go
// internal/cli/errors.go
const (
    ErrTeamRequired = "--team is required (or run 'linear init' to set a default)"
)

// Usage:
if teamID == "" {
    return fmt.Errorf(ErrTeamRequired)
}
```

**Estimated Effort:** 30 minutes

---

### H5. Refactor Search Command Arguments
**Source:** Gemini Review
**File:** `internal/cli/search.go:126`
**Issue:** `searchIssues()` takes 16 arguments (code smell)

**Current:**
```go
err := searchIssues(cmd, textQuery, team, priority, state, assignee,
                   labels, blockedBy, blocks, hasBlockers, hasDeps,
                   hasCircular, maxDepth, cycle, limit, formatStr)
```

**Fix:**
```go
type SearchOptions struct {
    TextQuery    string
    Team         string
    Priority     int
    State        string
    Assignee     string
    Labels       string
    BlockedBy    string
    Blocks       string
    HasBlockers  bool
    HasDeps      bool
    HasCircular  bool
    MaxDepth     int
    Cycle        string
    Limit        int
    Format       string
}

func searchIssues(cmd *cobra.Command, opts SearchOptions) error {
    // ...
}
```

**Estimated Effort:** 2 hours

---

## ⚠️ MEDIUM Priority (Fixes & Improvements)

### M1. Fix searchAll Error Handling
**Source:** Gemini Review + Claude Exploration
**File:** `internal/cli/search.go:309-343`
**Issue:** Prints errors to stdout, returns nil even if all searches fail

**Current:**
```go
func searchAll(...) error {
    if err := searchIssues(...); err != nil {
        fmt.Printf("Error searching issues: %v\n", err)  // ← stdout
    }
    // ... more searches
    return nil  // ← Always success!
}
```

**Problems:**
1. Errors go to stdout (breaks piping)
2. Command exits 0 even if all searches failed
3. User doesn't know searches failed

**Fix:**
```go
func searchAll(cmd *cobra.Command, ...) error {
    var errs []error

    if err := searchIssues(...); err != nil {
        errs = append(errs, fmt.Errorf("issues: %w", err))
        fmt.Fprintf(cmd.ErrOrStderr(), "Warning: failed to search issues: %v\n", err)
    }

    // ... more searches

    if len(errs) == 4 {  // All failed
        return fmt.Errorf("all searches failed")
    }
    if len(errs) > 0 {
        fmt.Fprintf(cmd.ErrOrStderr(), "\nWarning: %d search(es) failed\n", len(errs))
    }
    return nil
}
```

**Estimated Effort:** 1 hour

---

### M2. Fix Unused Format Parameter
**Source:** Claude Exploration
**Files:** `internal/cli/search.go`
**Issue:** `searchProjects()` and `searchUsers()` accept `formatStr` but ignore it

**Current:**
- `searchIssues()` - Uses `formatStr` ✓
- `searchCycles()` - Uses `formatStr` ✓
- `searchProjects()` - **Ignores `formatStr`** ✗
- `searchUsers()` - **Ignores `formatStr`** ✗

**Fix (Option 1 - Implement):**
Implement format handling in ProjectService and UserService

**Fix (Option 2 - Remove):**
Remove parameter, document that these searches use default format

**Recommended:** Option 2 (simpler)
**Estimated Effort:** 30 minutes

---

### M3. Use strconv.Atoi for Cycle Number Check
**Source:** Gemini Review
**File:** `internal/cli/cycles.go:137-143`
**Issue:** Manual rune loop to check if string is a number

**Current:**
```go
isNumber := true
for _, c := range cycleID {
    if c < '0' || c > '9' {
        isNumber = false
        break
    }
}
```

**Fix:**
```go
_, err := strconv.Atoi(cycleID)
isNumber := err == nil
```

**Estimated Effort:** 5 minutes

---

### M4. Standardize Variable Naming
**Source:** Claude Exploration
**Files:** Multiple
**Issue:** Same concept uses `team`, `teamID`, `teamKey` inconsistently

**Current:**
- `issues list` → `teamID`
- `issues create` → `team`
- `cycles` → `teamID`
- `deps` → `teamKey`

**Fix:**
Standardize all to `teamID` (most common, clearest meaning)

**Estimated Effort:** 1 hour

---

### M5. Fix golangci-lint Issues
**Source:** golangci-lint (18 issues found)

**errcheck (5 issues):**
- `auth.go:295` - Unchecked `cmd.Run()`
- `init.go:165` - Unchecked `f.Close()`
- `init.go:170,172` - Unchecked `f.WriteString()`
- `skills.go:175` - Unchecked `os.RemoveAll()`

**staticcheck (12 issues):**
- `auth.go:312, init.go:47, issues.go:929` - Using deprecated `TokenExists()`
- `cycles.go:162` - Error string ends with punctuation/newline
- `deps.go:335,341` - Use `fmt.Fprintf` instead of `WriteString(fmt.Sprintf())`
- `issues.go:142, search.go:219,246` - Could use tagged switch on formatStr
- `issues.go:835,881,920` - Use `fmt.Printf` instead of `fmt.Println(fmt.Sprintf())`

**unused (1 issue):**
- `onboard.go:154` - Unused function `getHomeDir()`

**Estimated Effort:** 2 hours

---

## 📝 LOW Priority (Polish & Documentation)

### L1. Add Package Documentation
**Source:** Gemini Review
**File:** `internal/cli/doc.go` (missing)
**Issue:** No package-level documentation

**Fix:**
```go
// Package cli implements the command-line interface for the Linear CLI.
//
// The CLI is built using Cobra and provides commands for interacting with
// the Linear API including issues, projects, cycles, and search.
//
// Commands follow a consistent pattern:
//   - Team resolution: --team flag → .linear.yaml → error
//   - Limit validation: Default 25, max 250
//   - Error handling: Wrap with fmt.Errorf and %w
//
// See individual command files for details:
//   - issues.go: Issue management
//   - projects.go: Project management
//   - cycles.go: Cycle/sprint management
//   - search.go: Unified search across entities
package cli
```

**Estimated Effort:** 30 minutes

---

### L2. Extract Cycle Number Detection Helper
**Source:** Claude Exploration
**File:** `internal/cli/cycles.go:137-169`
**Issue:** Complex logic could be reusable helper

**Fix:**
```go
// internal/cli/helpers.go
func isCycleNumber(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}
```

**Estimated Effort:** 15 minutes

---

## 📊 Test Coverage Improvements

**Current Coverage:**
```
cmd/linear:          0.0%  (untestable due to main)
internal/cli:       12.5%  ← CRITICAL
internal/config:    71.0%  ✓
internal/format:    82.3%  ✓
internal/linear:    64.5%  ✓
internal/oauth:      3.9%  ← HIGH
internal/service:    1.3%  ← CRITICAL
internal/skills:     0.0%  ← HIGH
internal/token:     41.7%  ← MEDIUM
```

**Target Coverage:**
- internal/cli: 75%+ (requires C2: Dependency Injection)
- internal/service: 90%+
- internal/oauth: 60%+
- internal/skills: 50%+
- internal/token: 80%+

**Dependencies:**
- Must complete **C2 (Dependency Injection)** first
- Then write unit tests for each command
- Mock Linear API responses
- Test error paths

**Estimated Effort:** 5-7 days (after C2 complete)

---

## 🎯 Summary

| Priority | Count | Estimated Effort |
|----------|-------|------------------|
| CRITICAL | 2 | 3-4 days |
| HIGH | 5 | 1.5 days |
| MEDIUM | 5 | 5 hours |
| LOW | 2 | 1 hour |
| **Coverage** | - | 5-7 days |
| **TOTAL** | 14+ | **2-3 weeks** |

**Recommended Order:**
1. ✅ C1 - Priority 0 bug (30 min) - Quick win
2. ✅ H1-H5 - All HIGH items (1.5 days) - Foundation
3. ✅ M1-M5 - All MEDIUM items (5 hours) - Clean up
4. ✅ C2 - Dependency injection (2-3 days) - Enables testing
5. ✅ Coverage - Write tests (5-7 days) - Production ready
6. ✅ L1-L2 - Polish (1 hour) - Final touch

**Phase 1 (This Week):** C1, H1-H5, M1-M5 = Production-quality code
**Phase 2 (Next Week):** C2, Coverage = Production-ready with tests

---

## 📋 Next Steps

1. **Stefan reviews this task list**
2. **Prioritize tasks** - Which to do first?
3. **Create GitHub issues** - Track progress
4. **Start with quick wins** - C1, H1-H4 (< 1 day)
5. **Tackle dependency injection** - C2 (2-3 days, enables testing)
6. **Write comprehensive tests** - Coverage improvements

---

**Generated by:** Gemini AI + Claude Code + golangci-lint
**Date:** 2026-01-24
**Status:** 📋 Awaiting Stefan's Review
