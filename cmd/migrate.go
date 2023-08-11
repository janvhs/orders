package cmd

import "github.com/spf13/cobra"

func NewMigrateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: add migrations via bun migrate, if using bun, or goose
			return nil
		},
	}

	return command
}
