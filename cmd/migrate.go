package cmd

import (
	"git.bode.fun/orders/config"
	odb "git.bode.fun/orders/db"
	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			cnf, err := config.NewFromEnv()
			if err != nil {
				return err
			}

			shouldSeed, err := cmd.Flags().GetBool("seed")
			if err != nil {
				return err
			}

			db, err := odb.Connect(cnf.DB.DSN, cnf.IsDevelopment)
			if err != nil {
				return err
			}

			err = odb.Migrate(db.DB)
			if err != nil {
				return err
			}

			if shouldSeed {
				err = odb.Seed(db)
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
