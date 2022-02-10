package cmd

import (
	"context"
	"log"
	"os"

	"redis-viewer/internal/config"
	"redis-viewer/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redis-viewer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		cfg := config.GetConfig()

		var rdb *redis.Client
		switch cfg.Mode {
		case "sentinel":
			rdb = redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    cfg.MasterName,
				SentinelAddrs: cfg.SentinelAddrs,
				Password:      cfg.SentinelPassword,
			})
		default:
			rdb = redis.NewClient(&redis.Options{
				Addr:     cfg.Addr,
				Password: cfg.Password,
				DB:       cfg.DB,
			})
		}

		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("connect to redis failed: ", err)
		}

		p := tea.NewProgram(tui.New(rdb), tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatal("start failed: ", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.redis-viewer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
