package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func newTasksCmd() *cobra.Command {
	tasksCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Export Linear issues to Claude Code task format",
		Long: `Export Linear issues and their dependencies to Claude Code task JSON files.

This command converts Linear issues (with their complete dependency trees) into
Claude Code task format, preserving hierarchies and blocking relationships.`,
	}

	tasksCmd.AddCommand(
		newTasksExportCmd(),
	)

	return tasksCmd
}

func newTasksExportCmd() *cobra.Command {
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "export <identifier> <output-folder>",
		Short: "Export Linear issue and dependencies to Claude Code tasks",
		Long: `Export a Linear issue and its complete dependency tree to Claude Code task format.

IMPORTANT BEHAVIORS:
- Always exports recursively with full dependency chain
- Detects and errors on circular dependencies
- Writes JSON files named {identifier}.json (e.g., CEN-123.json)
- Creates output folder if it doesn't exist
- Bottom-up hierarchy: children block parent

OUTPUT FOLDER:
The output folder can be any directory path, including Claude Code session folders:
- Custom folder: ./my-tasks/
- Claude session: ~/.claude/tasks/<session-uuid>/

Each task is written as a separate JSON file matching Claude Code's schema.`,
		Example: `  # Export to custom folder
  linear tasks export CEN-123 ./my-tasks

  # Export to Claude Code session
  linear tasks export CEN-123 ~/.claude/tasks/a5721284-f64e-4705-8983-b7d6c4e032aa/

  # Preview without writing files
  linear tasks export CEN-123 ./my-tasks --dry-run`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			identifier := args[0]
			outputFolder := args[1]

			deps, err := getDeps(cmd)
			if err != nil {
				return err
			}

			// Ensure TaskExport service exists
			if deps.TaskExport == nil {
				return fmt.Errorf("TaskExport service not initialized")
			}

			// Clean up folder path
			outputFolder = filepath.Clean(outputFolder)

			// Print export info
			fmt.Printf("Fetching %s and complete dependency tree...\n\n", identifier)

			// Export
			result, err := deps.TaskExport.Export(identifier, outputFolder, dryRun)
			if err != nil {
				return err
			}

			// Print results
			if dryRun {
				fmt.Printf("DRY RUN - No files written\n\n")
			}

			fmt.Printf("Analyzing dependencies:\n")

			// Show tree structure (simplified - just count)
			fmt.Printf("  ✓ %s and %d related issues\n\n", identifier, result.TotalTasks-1)

			if len(result.CircularDepsFound) > 0 {
				fmt.Printf("⚠ Circular dependencies detected:\n")
				for _, cycle := range result.CircularDepsFound {
					fmt.Printf("  %s\n", cycle)
				}
				return nil
			}

			fmt.Printf("Checking for circular dependencies... None found\n\n")

			if !dryRun {
				fmt.Printf("Exported %d tasks to: %s\n\n", result.TotalTasks, outputFolder)

				fmt.Printf("Tasks written:\n")
				for _, task := range result.Tasks {
					blockedByDesc := ""
					if len(task.BlockedBy) > 0 {
						blockedByDesc = fmt.Sprintf(" (blocked by: %s)", formatBlockedByList(task.BlockedBy))
					}
					fmt.Printf("  - %s.json%s\n", task.ID, blockedByDesc)
				}
				fmt.Printf("\n")

				fmt.Printf("Dependencies: %d tasks, %d blocking relationships\n", result.TotalTasks, result.DependencyCount)
			} else {
				fmt.Printf("Would export %d tasks to: %s\n\n", result.TotalTasks, outputFolder)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview without writing files")

	return cmd
}

// formatBlockedByList formats a list of blockers for display
func formatBlockedByList(blockers []string) string {
	if len(blockers) == 0 {
		return ""
	}
	if len(blockers) == 1 {
		return blockers[0]
	}
	if len(blockers) <= 3 {
		return strings.Join(blockers, ", ")
	}
	// Show first 3 and count
	return fmt.Sprintf("%s, +%d more", strings.Join(blockers[:3], ", "), len(blockers)-3)
}
