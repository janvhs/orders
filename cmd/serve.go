package cmd

import (
	"errors"
	"net/http"

	"git.bode.fun/orders/server"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // FIXME: Move to modernc.org/sqlite
	"github.com/spf13/cobra"
)

func NewServeCommand(logger *log.Logger) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {

			addr, err := cmd.Flags().GetString("addr")
			if err != nil {
				return err
			}

			db, err := sqlx.Connect("sqlite3", ":memory:")
			if err != nil {
				return err
			}

			srv := server.New(logger, db)

			logger.Infof("Starting server on %s", addr)
			if err := http.ListenAndServe(addr, srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}

			return nil
		},
	}

	command.Flags().StringP("addr", "a", "127.0.0.1:8080", "Address to listen on")

	return command
}
