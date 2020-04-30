package cmd

import (
	"log"
	"time"

	"github.com/justym/spotify-timer/pkg/auth"
	"github.com/justym/spotify-timer/pkg/player"
	"github.com/justym/spotify-timer/pkg/util"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "spotify-timer [minutes]",
		Short: "spotify-timer is sleep timer for spotify",
		Run:   pause,
	}

	minutes time.Duration
)

func init() {
	tenMinutes := time.Duration(10) * time.Minute
	rootCmd.PersistentFlags().DurationVarP(&minutes, "minutes", "m", tenMinutes, "You can use flags to define duration(minutes)")
}

func pause(_ *cobra.Command, args []string) {
	if len(args) == 1 {
		minutes, err := util.AtoTime(args[0])
		if err != nil || minutes == -1 {
			log.Fatal(err, minutes)
		}
	}

	authrizedClient := auth.NewClinet()
	if err := player.Pause(authrizedClient.Client, minutes); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
