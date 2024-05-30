/*
Copyright © 2024 Vlad Zaharia hey@vlad.gg

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

func init() {
	cobra.OnInitialize(initLog, initConfig)
	setupPFlags()
}

func MakeCmd(cli command.Cli) *cobra.Command {
	return rootCmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage Docker compose services at scale",
	Long:  `docker-apps lets you manage Docker compose services at scale.`,
}

// setupFlags sets up the persistent flags (across commands)
func setupPFlags() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docker-apps.yaml)")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "whether to show debug log")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initLog initializes the logging system
func initLog() {
	styles := log.DefaultStyles()

	// Override timestamp formatting
	log.SetTimeFormat(time.Kitchen)
	styles.Timestamp = lipgloss.NewStyle().
		Faint(true)

	baseStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("0"))

	// Override the default error level style.
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FTL").
		Background(lipgloss.Color("134")).
		Inherit(baseStyle)

	// Override the default error level style.
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERR").
		Background(lipgloss.Color("204")).
		Inherit(baseStyle)

	// Override the default debug level style.
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WRN").
		Background(lipgloss.Color("192")).
		Inherit(baseStyle)

	// Override the default debug level style.
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INF").
		Background(lipgloss.Color("86")).
		Inherit(baseStyle)

	// Override the default debug level style.
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DBG").
		Background(lipgloss.Color("63")).
		Inherit(baseStyle)

	log.SetStyles(styles)

	// Set log level
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.Debug("Set log level to debug!")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Debug("Initializing config...")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/docker")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".docker-apps")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}
}
