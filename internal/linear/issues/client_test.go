package issues

import (
	"testing"

	"github.com/joa23/linear-cli/internal/linear/core"
)

func TestBuildUpdateInput_DelegateID(t *testing.T) {
	tests := []struct {
		name           string
		input          core.UpdateIssueInput
		expectDelegate bool
		expectAssignee bool
	}{
		{
			name: "human user - uses assigneeId",
			input: core.UpdateIssueInput{
				AssigneeID: strPtr("user-uuid-123"),
			},
			expectAssignee: true,
			expectDelegate: false,
		},
		{
			name: "OAuth application - uses delegateId",
			input: core.UpdateIssueInput{
				DelegateID: strPtr("app-uuid-456"),
			},
			expectAssignee: false,
			expectDelegate: true,
		},
		{
			name: "unassign - empty assigneeId",
			input: core.UpdateIssueInput{
				AssigneeID: strPtr(""),
			},
			expectAssignee: true, // Still sets assigneeId to nil
			expectDelegate: false,
		},
		{
			name: "remove delegation - empty delegateId",
			input: core.UpdateIssueInput{
				DelegateID: strPtr(""),
			},
			expectAssignee: false,
			expectDelegate: true, // Still sets delegateId to nil
		},
		{
			name: "both set - both fields present",
			input: core.UpdateIssueInput{
				AssigneeID: strPtr("user-uuid"),
				DelegateID: strPtr("app-uuid"),
			},
			expectAssignee: true,
			expectDelegate: true,
		},
		{
			name:           "neither set - no assignment fields",
			input:          core.UpdateIssueInput{},
			expectAssignee: false,
			expectDelegate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildUpdateInput(tt.input)

			_, hasAssignee := result["assigneeId"]
			_, hasDelegate := result["delegateId"]

			if hasAssignee != tt.expectAssignee {
				t.Errorf("assigneeId presence = %v, want %v", hasAssignee, tt.expectAssignee)
			}
			if hasDelegate != tt.expectDelegate {
				t.Errorf("delegateId presence = %v, want %v", hasDelegate, tt.expectDelegate)
			}
		})
	}
}

func TestHasFieldsToUpdate_DelegateID(t *testing.T) {
	tests := []struct {
		name     string
		input    core.UpdateIssueInput
		expected bool
	}{
		{
			name:     "empty input",
			input:    core.UpdateIssueInput{},
			expected: false,
		},
		{
			name: "only delegateId",
			input: core.UpdateIssueInput{
				DelegateID: strPtr("app-uuid"),
			},
			expected: true,
		},
		{
			name: "only assigneeId",
			input: core.UpdateIssueInput{
				AssigneeID: strPtr("user-uuid"),
			},
			expected: true,
		},
		{
			name: "only title",
			input: core.UpdateIssueInput{
				Title: strPtr("New Title"),
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasFieldsToUpdate(tt.input)
			if result != tt.expected {
				t.Errorf("hasFieldsToUpdate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
