package cmd

import (
	"context"
	"log"
	"os"

	"github.com/saltfishpr/redis-viewer/internal/config"
	"github.com/saltfishpr/redis-viewer/internal/constant"
	"github.com/saltfishpr/redis-viewer/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "redis-viewer",
	Short: "view redis data in terminal.",
	Long:  `Redis Viewer is a tool to view redis data in terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		cfg := config.GetConfig()

		rdb := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:        cfg.Addrs,
			DB:           cfg.DB,
			Username:     cfg.Username,
			Password:     cfg.Password,
			MaxRetries:   constant.MaxRetries,
			MaxRedirects: constant.MaxRedirects,
			MasterName:   cfg.MasterName,
		})
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal("connect to redis failed: ", err)
		}

		p := tea.NewProgram(tui.New(rdb), tea.WithAltScreen(), tea.WithMouseCellMotion())
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
