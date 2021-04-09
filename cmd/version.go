package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// These are set through build flags. Details available here:
// https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
var (
	Version   = "unknown-version"
	Commit    = "unknown-commit"
	BuildTime = "unknown-buildtime"
)
var versionCmd = &cobra.Command{
	Use:    "version",
	Short:  "Print version data",
	Long:   "Print version, git commit, and build time.",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Version: %s", Version)
		log.Infof("Commit: %s", Commit)
		log.Infof("Build time: %s", BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
