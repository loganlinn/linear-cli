package service

import (
	"github.com/joa23/linear-cli/internal/linear/core"
)

// IssueClientOperations defines the minimal interface needed by IssueService.
// This follows the "consumer defines interface" pattern for dependency injection,
// enabling mock implementations for unit testing.
type IssueClientOperations interface {
	// Smart resolver-aware methods
	CreateIssue(title, description, teamKeyOrName string) (*core.Issue, error)
	GetIssue(identifierOrID string) (*core.Issue, error)
	UpdateIssue(identifierOrID string, input core.UpdateIssueInput) (*core.Issue, error)
	UpdateIssueState(identifierOrID, stateID string) error
	AssignIssue(identifierOrID, assigneeNameOrEmail string) error
	ListAssignedIssues(limit int) ([]core.Issue, error)
	SearchIssues(filters *core.IssueSearchFilters) (*core.IssueSearchResult, error)

	// Resolver operations
	ResolveTeamIdentifier(keyOrName string) (string, error)
	ResolveUserIdentifier(nameOrEmail string) (string, error)
	ResolveCycleIdentifier(numberOrNameOrID, teamID string) (string, error)

	// Metadata operations
	UpdateIssueMetadataKey(issueID, key string, value interface{}) error

	// Sub-client delegations (will be refactored in Phase 2)
	CreateComment(issueID, body string) (*core.Comment, error)
	CreateCommentReply(issueID, parentID, body string) (*core.Comment, error)
	AddReaction(targetID, emoji string) error
	GetWorkflowStateByName(teamID, stateName string) (*core.WorkflowState, error)
	ListAllIssues(filter *core.IssueFilter) (*core.ListAllIssuesResult, error)
	GetViewer() (*core.User, error)
}

// CycleClientOperations defines the minimal interface needed by CycleService
type CycleClientOperations interface {
	// Smart resolver-aware methods
	ListCycles(filter *core.CycleFilter) (*core.CycleSearchResult, error)
	GetActiveCycle(teamKeyOrName string) (*core.Cycle, error)
	CreateCycle(input *core.CreateCycleInput) (*core.Cycle, error)

	// Pass-through methods (will be removed in Phase 2)
	GetCycle(cycleID string) (*core.Cycle, error)
	GetCycleIssues(cycleID string, limit int) ([]core.Issue, error)

	// Resolver operations
	ResolveTeamIdentifier(keyOrName string) (string, error)
	ResolveCycleIdentifier(numberOrNameOrID, teamID string) (string, error)
	ResolveUserIdentifier(nameOrEmail string) (string, error)
}

// ProjectClientOperations defines the minimal interface needed by ProjectService
type ProjectClientOperations interface {
	// Smart resolver-aware methods
	CreateProject(name, description, teamKeyOrName string) (*core.Project, error)

	// Resolver operations
	ResolveTeamIdentifier(keyOrName string) (string, error)
	ResolveUserIdentifier(nameOrEmail string) (string, error)

	// Pass-through methods (will be removed in Phase 2)
	GetProject(projectID string) (*core.Project, error)
	ListAllProjects(limit int) ([]core.Project, error)
	ListByTeam(teamID string, limit int) ([]core.Project, error)
	ListUserProjects(userID string, limit int) ([]core.Project, error)
	UpdateProject(projectID string, input interface{}) (*core.Project, error)
	UpdateProjectState(projectID, state string) error
	UpdateProjectDescription(projectID, newDescription string) error
	UpdateProjectMetadataKey(projectID, key string, value interface{}) error
	GetViewer() (*core.User, error)
}

// UserClientOperations defines the minimal interface needed by UserService
type UserClientOperations interface {
	// Smart resolver-aware methods
	GetUser(idOrEmail string) (*core.User, error)

	// Pass-through methods (will be removed in Phase 2)
	GetViewer() (*core.User, error)
	ListUsersWithPagination(filter *core.UserFilter) (*core.ListUsersResult, error)

	// Resolver operations
	ResolveUserIdentifier(nameOrEmail string) (string, error)
	ResolveTeamIdentifier(keyOrName string) (string, error)
}

// SearchClientOperations defines the minimal interface needed by SearchService
type SearchClientOperations interface {
	// Smart resolver-aware methods
	SearchIssues(filters *core.IssueSearchFilters) (*core.IssueSearchResult, error)
	ListCycles(filter *core.CycleFilter) (*core.CycleSearchResult, error)

	// Resolver operations
	ResolveTeamIdentifier(keyOrName string) (string, error)
	ResolveUserIdentifier(nameOrEmail string) (string, error)
	ResolveCycleIdentifier(numberOrNameOrID, teamID string) (string, error)

	// Pass-through methods (will be removed in Phase 2)
	GetIssueWithRelations(identifier string) (*core.IssueWithRelations, error)
	ListAllProjects(limit int) ([]core.Project, error)
	ListUsers(filter *core.UserFilter) ([]core.User, error)
}
