package cmd

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"time"

	"git.bode.fun/orders/config"
	"git.bode.fun/orders/db"
	"git.bode.fun/orders/server"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func NewServeCommand(logger *log.Logger, templateFS embed.FS) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {

			cnf, err := config.NewFromEnv()
			if err != nil {
				return err
			}

			host, err := cmd.Flags().GetString("host")
			if err != nil {
				return err
			}
			cnf.Server.Host = host

			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			cnf.Server.Port = port

			db, err := db.Connect(cnf.DB.DSN, cnf.IsDevelopment)
			if err != nil {
				return err
			}

			addr := fmt.Sprintf("%s:%d", cnf.Server.Host, cnf.Server.Port)
			handler := server.New(db, cnf.IsDevelopment, templateFS)

			// TODO: find a good value for the timeouts.
			// Fixes gosec issue G114
			srv := &http.Server{
				Addr:         addr,
				Handler:      handler,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			logger.Infof("Starting server on %s", addr)
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}

			return nil
		},
	}

	defaultCnf := config.Default()

	command.Flags().String("host", defaultCnf.Server.Host, "Host to listen on")
	command.Flags().IntP("port", "p", defaultCnf.Server.Port, "Port to listen on")

	return command
}
