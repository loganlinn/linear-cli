package service

import (
	"github.com/joa23/linear-cli/internal/format"
	"github.com/joa23/linear-cli/internal/linear"
)

// IssueServiceInterface defines the contract for issue operations
type IssueServiceInterface interface {
	Get(identifier string, outputFormat format.Format) (string, error)
	Search(filters *SearchFilters) (string, error)
	ListAssigned(limit int, outputFormat format.Format) (string, error)
	ListAssignedWithPagination(pagination *linear.PaginationInput) (string, error)
	Create(input *CreateIssueInput) (string, error)
	Update(identifier string, input *UpdateIssueInput) (string, error)
	GetComments(identifier string) (string, error)
	AddComment(identifier, body string) (string, error)
	ReplyToComment(issueIdentifier, parentCommentID, body string) (*linear.Comment, error)
	AddReaction(targetID, emoji string) error
	GetIssueID(identifier string) (string, error)
}

// CycleServiceInterface defines the contract for cycle operations
type CycleServiceInterface interface {
	Get(cycleIDOrNumber string, teamID string, outputFormat format.Format) (string, error)
	Search(filters *CycleFilters) (string, error)
	Create(input *CreateCycleInput) (string, error)
	Analyze(input *AnalyzeInput) (string, error)
}

// ProjectServiceInterface defines the contract for project operations
type ProjectServiceInterface interface {
	Get(projectID string) (string, error)
	ListAll(limit int) (string, error)
	ListByTeam(teamID string, limit int) (string, error)
	ListUserProjects(limit int) (string, error)
	Create(input *CreateProjectInput) (string, error)
	Update(projectID string, input *UpdateProjectInput) (string, error)
}

// SearchServiceInterface defines the contract for unified search
type SearchServiceInterface interface {
	Search(opts *SearchOptions) (string, error)
}

// TeamServiceInterface defines the contract for team operations
type TeamServiceInterface interface {
	Get(identifier string) (string, error)
	ListAll() (string, error)
	GetLabels(identifier string) (string, error)
	GetWorkflowStates(identifier string) (string, error)
}

// UserServiceInterface defines the contract for user operations
type UserServiceInterface interface {
	GetViewer() (string, error)
	Get(identifier string) (string, error)
	Search(filters *UserFilters) (string, error)
	ResolveByName(name string) (string, error)
}

// Verify implementations satisfy interfaces (compile-time check)
var (
	_ IssueServiceInterface   = (*IssueService)(nil)
	_ CycleServiceInterface   = (*CycleService)(nil)
	_ ProjectServiceInterface = (*ProjectService)(nil)
	_ SearchServiceInterface  = (*SearchService)(nil)
	_ TeamServiceInterface    = (*TeamService)(nil)
	_ UserServiceInterface    = (*UserService)(nil)
)
