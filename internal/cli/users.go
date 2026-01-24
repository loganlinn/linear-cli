package cli

import (
	"fmt"

	"github.com/joa23/linear-cli/internal/service"
	"github.com/spf13/cobra"
)

func newUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"u", "user"},
		Short:   "Manage Linear users",
		Long:    "List users and view user details.",
	}

	cmd.AddCommand(
		newUsersListCmd(),
		newUsersGetCmd(),
		newUsersMeCmd(),
	)

	return cmd
}

func newUsersListCmd() *cobra.Command {
	var teamID string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		Long:  "List users in the workspace or a specific team.",
		RunE: func(cmd *cobra.Command, args []string) error {
			deps, err := getDeps(cmd)
			if err != nil {
				return err
			}

			filters := &service.UserFilters{
				TeamID: teamID,
				Limit:  50,
			}

			output, err := deps.Users.Search(filters)
			if err != nil {
				return fmt.Errorf("failed to list users: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVar(&teamID, "team", "", TeamFlagDescription)

	return cmd
}

func newUsersGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <user-id>",
		Short: "Get user details",
		Long:  "Display detailed information about a specific user.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			userID := args[0]

			deps, err := getDeps(cmd)
			if err != nil {
				return err
			}

			output, err := deps.Users.Get(userID)
			if err != nil {
				return fmt.Errorf("failed to get user: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}
}

func newUsersMeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Show current user",
		Long:  "Display information about the currently authenticated user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			deps, err := getDeps(cmd)
			if err != nil {
				return err
			}

			output, err := deps.Users.GetViewer()
			if err != nil {
				return fmt.Errorf("failed to get current user: %w", err)
			}

			fmt.Println(output)
			return nil
		},
	}
}
