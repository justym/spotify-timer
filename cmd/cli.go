package cmd

import (
	"log"
	"time"

	"github.com/justym/spotify-timer/pkg/auth"
	"github.com/justym/spotify-timer/pkg/player"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "spotify-timer",
		Short: "spotify-timer is sleep timer for spotify",
		Run:   run,
	}

	pauseDuration time.Duration
)

func init() {
	rootCmd.PersistentFlags().DurationVarP(
		&pauseDuration, "duration", "d", pauseDuration, "pause duration")
	rootCmd.MarkPersistentFlagRequired("duration")
}

func run(cmd *cobra.Command, args []string) {
	if err := pause(); err != nil {
		log.Fatal(err)
	}
}

func pause() error {
	authrizedClient, err := auth.NewClient()
	if err != nil {
		return err
	}

	if err := player.Pause(authrizedClient.Client, pauseDuration); err != nil {
		return err
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
