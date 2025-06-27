package cmd

import (
	"fmt"
	"github.com/analog-substance/sulfur/cmd/acid/pkg/sulfur"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var sulfurAPIClient *sulfur.APIClient
var jsonOutput = false

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "acid",
	Short: "get data into Sulfur",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		sulfurAPIClient = sulfur.New(viper.GetString("api-endpoint"), viper.GetString("api-user"), viper.GetString("api-pass"))
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sulfur.yaml)")
	//RootCmd.PersistentFlags().Bool("debug", false, "Debug mode")

	RootCmd.PersistentFlags().String("org", "", "org id")
	RootCmd.PersistentFlags().String("url", "http://127.0.0.1:8090", "URL to use as api endpoint")
	RootCmd.PersistentFlags().String("api-user", "", "Username to use for api")
	RootCmd.PersistentFlags().String("api-pass", "", "Password to use for api")
	RootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", jsonOutput, "Output as JSON")
	//RootCmd.PersistentFlags().StringP("scope", "s", scopious.DefaultScope, "Scope name")

	//rootCmd.PersistentFlags().String("domains-file", "scope-domains.txt", "where in-scope domains are located.")
	//rootCmd.PersistentFlags().String("ips-file", "scope-ips.txt", "where in-scope IP addresses are located.")
	//rootCmd.PersistentFlags().String("ignore-domains", "ignore-scope-domains.txt", "where out-of-scope IP addresses are located.")
	//rootCmd.PersistentFlags().String("ignore-ips", "ignore-scope-ips.txt", "where out-of-scope domains addresses are located.")

	viper.BindPFlag("organization", RootCmd.PersistentFlags().Lookup("org"))
	viper.BindPFlag("api-endpoint", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("api-user", RootCmd.PersistentFlags().Lookup("api-user"))
	viper.BindPFlag("api-pass", RootCmd.PersistentFlags().Lookup("api-pass"))
	//viper.BindPFlag("ignore-domains", rootCmd.PersistentFlags().Lookup("ignore-domains"))
	//viper.BindPFlag("ignore-ips", rootCmd.PersistentFlags().Lookup("ignore-ips"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".scopious" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sulfur")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
}
