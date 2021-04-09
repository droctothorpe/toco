package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command.
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Consolidate git add, commit, and push",
	Long: `This subcommand reduces contribution friction by staging, commiting, and pushing all of your changes with a single, simple command.
	
It executes the following commands in sequence: 
	> git pull
	> git add .
	> git commit -m "Update"
	> git push`,
	RunE: func(cmd *cobra.Command, args []string) error {
		base := "git"
		argsPull := []string{"pull", "origin", "master"}
		argsAdd := []string{"add", "."}
		argsCommit := []string{"commit", "-m", "Update"}
		argsPush := []string{"push", "origin", "master"}

		var argsSlice [][]string
		argsSlice = append(argsSlice, argsPull)
		argsSlice = append(argsSlice, argsAdd)
		argsSlice = append(argsSlice, argsCommit)
		argsSlice = append(argsSlice, argsPush)

		for _, arguments := range argsSlice {
			command := exec.Command(base, arguments...)
			configureCommand(command)

			if err := command.Run(); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func configureCommand(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
}
