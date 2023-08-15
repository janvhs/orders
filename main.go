// orders - an program manage your orders
// Copyright (C) 2023  Jan Fooken
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"embed"
	"os"
	"path"
	"runtime/debug"
	"syscall"

	"git.bode.fun/orders/cmd"
	"github.com/charmbracelet/log"
)

var (
	// FIXME: This kind of doesn't override, when building with -ldflags
	defaultVersion = "unknown (built from source)"
	defaultAppName = "app"

	Version   = defaultVersion
	AppName   = defaultAppName
	Vendor    = "fun.bode"
	CommitSHA = ""
)

// TODO: Get rid of this because it is annoying to pass around
//
//go:embed web/templates/*.html web/templates/layouts/*.html
var templateFS embed.FS

// TODO: Get rid of this because it is annoying to pass around
//
//go:embed web/static
var staticFS embed.FS

// FIXME: This currently doesn't work
func setMetaDefaults() {
	if info, ok := debug.ReadBuildInfo(); ok {
		if (Version == defaultVersion || Version == "") && info.Main.Sum != "" {
			Version = info.Main.Version
		}

		if (AppName == defaultAppName || AppName == "") && info.Main.Path != "" {
			AppName = path.Base(path.Clean(info.Main.Path))
		}
	}
}

func main() {
	prepareProcess()

	logger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix: AppName,
	})

	app := cmd.New(AppName, Version, CommitSHA)

	app.AddCommand(
		cmd.NewServeCommand(logger, templateFS, staticFS),
		cmd.NewMigrateCommand(),
	)

	err := app.Execute()
	if err != nil {
		logger.Fatal(err)
	}
}

func prepareProcess() {
	setMetaDefaults()
	ensureFileOwner()
}

// Files created by this process,
// are only accessible to the user,
// who started this process.
// Their group can not access them
func ensureFileOwner() {
	syscall.Umask(0177)
}
