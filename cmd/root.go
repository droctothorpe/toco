package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	color "github.com/logrusorgru/aurora"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "toco",
	Short: "A brief description of your application",
	Long: color.Sprintf(color.Bold(color.Cyan(` ______   ______     ______     ______    
/\__  _\ /\  __ \   /\  ___\   /\  __ \   
\/_/\ \/ \ \ \/\ \  \ \ \____  \ \ \/\ \  
   \ \_\  \ \_____\  \ \_____\  \ \_____\ 
    \/_/   \/_____/   \/_____/   \/_____/ 
                                          
`))) + `A CLI to automate TOC (table of contents) generation for GitHub wikis.

It generates a TOC based on your files and injects them into the
homepage (Home.md) and sidebar (_Sidebar.md) files.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toco.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "verbose logging")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".toco" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".toco")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
