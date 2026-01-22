package cli

import (
	"fmt"

	"github.com/joa23/linear-cli/internal/format"
	"github.com/joa23/linear-cli/internal/service"
	"github.com/spf13/cobra"
)

func newCyclesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cycles",
		Aliases: []string{"c", "cycle"},
		Short:   "Manage Linear cycles",
		Long:    "List, view, and analyze Linear cycles (sprints).",
	}

	cmd.AddCommand(
		newCyclesListCmd(),
		newCyclesGetCmd(),
		newCyclesAnalyzeCmd(),
	)

	return cmd
}

func newCyclesListCmd() *cobra.Command {
	var teamID string
	var activeOnly bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cycles for a team",
		Long: `List cycles (sprints) for a team.

REQUIRED:
- Team (via 'linear init' or --team flag)

OPTIONAL:
- --active: Filter to only active cycles

TIP: Run 'linear init' to set default team.`,
		Example: `  # Minimal - list all cycles (requires 'linear init')
  linear cycles list

  # With explicit team
  linear cycles list --team CEN

  # Only active cycles
  linear cycles list --active

  # Specific team, active only
  linear cycles list --team CEN --active`,
		Annotations: map[string]string{
			"required": "team (via init or --team flag)",
			"optional": "--active flag",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Use default team if not specified
			if teamID == "" {
				teamID = GetDefaultTeam()
			}
			if teamID == "" {
				return fmt.Errorf("--team is required (or run 'linear init' to set a default)")
			}

			svc, err := getCycleService()
			if err != nil {
				return err
			}

			filters := &service.CycleFilters{
				TeamID: teamID,
				Limit:  25,
				Format: format.Compact,
			}
			if activeOnly {
				filters.IsActive = &activeOnly
			}

			output, err := svc.Search(filters)
			if err != nil {
				return fmt.Errorf("failed to list cycles: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVar(&teamID, "team", "", "Team ID or key (uses .linear.yaml default)")
	cmd.Flags().BoolVar(&activeOnly, "active", false, "Only show active cycles")

	return cmd
}

func newCyclesGetCmd() *cobra.Command {
	var teamID string

	cmd := &cobra.Command{
		Use:   "get <cycle-id>",
		Short: "Get cycle details by number or UUID",
		Long: `Get details for a specific cycle.

CYCLE ID FORMATS:
- Number: 65 (requires team context from 'linear init')
- UUID: cycle-abc123-def456 (works without team)
- Name: "Sprint 2024-01" (requires team context)

REQUIRED:
- Cycle ID/number (positional argument)
- Team context (from 'linear init' or --team flag) if using numbers/names

TIP: Run 'linear init' once to set default team, then use cycle numbers directly.`,
		Example: `  # Minimal - get by number (requires 'linear init')
  linear cycles get 65

  # With explicit team
  linear cycles get 65 --team CEN

  # Get by UUID (no team needed)
  linear cycles get cycle-abc123-def456-789

  # Get by name
  linear cycles get "Sprint 2024-01" --team CEN`,
		Annotations: map[string]string{
			"required": "cycle-id, team (via init or --team flag for numbers/names)",
			"optional": "--team flag (if not using init)",
		},
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cycleID := args[0]

			// Use default team if not specified
			if teamID == "" {
				teamID = GetDefaultTeam()
			}

			svc, err := getCycleService()
			if err != nil {
				return err
			}

			output, err := svc.Get(cycleID, teamID, format.Full)
			if err != nil {
				return fmt.Errorf("failed to get cycle: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVar(&teamID, "team", "", "Team ID or key (uses .linear.yaml default, required for cycle numbers)")

	return cmd
}

func newCyclesAnalyzeCmd() *cobra.Command {
	var teamID string
	var cycleCount int
	var assigneeID string

	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze team velocity and cycle completion rates",
		Long: `Analyze historical cycle performance to inform planning.

Calculates:
- Average velocity (points/cycle)
- Completion rate (% of scoped work completed)
- Scope creep (% of work added mid-cycle)
- Recommended capacity for next cycles (P20/P50/P80)

REQUIRED:
- Team (via 'linear init' or --team flag)

OPTIONAL:
- --count: Number of past cycles to analyze (default: 10)
- --assignee: Filter by specific assignee

USE THIS BEFORE PLANNING: Always run analyze before planning cycles to understand capacity.`,
		Example: `  # Minimal - analyze last 10 cycles (requires 'linear init')
  linear cycles analyze

  # Complete - with all parameters
  linear cycles analyze --team CEN --count 15 --assignee stefan@centrum-ai.com

  # Common pattern - analyze for planning
  linear cycles analyze --team CEN --count 10

  # Analyze specific team member
  linear cycles analyze --assignee me --count 5`,
		Annotations: map[string]string{
			"required": "team (via init or --team flag)",
			"optional": "--count, --assignee flags",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Use default team if not specified
			if teamID == "" {
				teamID = GetDefaultTeam()
			}
			if teamID == "" {
				return fmt.Errorf("--team is required (or run 'linear init' to set a default)")
			}

			svc, err := getCycleService()
			if err != nil {
				return err
			}

			input := &service.AnalyzeInput{
				TeamID:                teamID,
				CycleCount:            cycleCount,
				AssigneeID:            assigneeID,
				IncludeRecommendation: true,
			}

			output, err := svc.Analyze(input)
			if err != nil {
				return fmt.Errorf("failed to analyze cycles: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVar(&teamID, "team", "", "Team ID or key (uses .linear.yaml default)")
	cmd.Flags().IntVar(&cycleCount, "count", 10, "Number of cycles to analyze")
	cmd.Flags().StringVar(&assigneeID, "assignee", "", "Filter by assignee ID")

	return cmd
}

func getCycleService() (*service.CycleService, error) {
	client, err := getLinearClient()
	if err != nil {
		return nil, err
	}
	return service.New(client).Cycles, nil
}
