package main

import (
	"os"
	"path"
	"runtime/debug"
	"syscall"

	"git.bode.fun/orders/cmd"
	"github.com/charmbracelet/log"
)

var (
	defaultVersion = "unknown (built from source)"
	defaultAppName = "app"

	Version   = defaultVersion
	AppName   = defaultAppName
	Vendor    = "fun.bode"
	CommitSHA = ""
)

func setMetaDefaults() {
	if info, ok := debug.ReadBuildInfo(); ok {
		if Version == defaultVersion && info.Main.Sum != "" {
			Version = info.Main.Version
		}

		if AppName == defaultAppName && info.Main.Path != "" {
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
		cmd.NewServeCommand(logger),
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
