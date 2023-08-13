package cmd

import (
	"git.bode.fun/orders/db"
	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			shouldSeed, err := cmd.Flags().GetBool("seed")
			if err != nil {
				return err
			}

			err = db.Migrate()
			if err != nil {
				return err
			}

			if shouldSeed {
				err = db.Seed()
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	command.Flags().Bool("seed", false, "Seed the database with test data")

	return command
}
