package cmd

import (
	"log"

	"github.com/justym/spotify-timer/pkg/auth"
	"github.com/justym/spotify-timer/pkg/player"
	"github.com/justym/spotify-timer/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spotify-timer [minutes]",
	Short: "spotify-timer is sleep timer for spotify",
	Run:   pause,
}

func pause(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("Please input [minutes]")
	}
	stopTime, err := util.AtoTime(args[0])
	if err != nil || stopTime == -1 {
		log.Fatal(err, stopTime)
	}
	authrizedClient := auth.NewClinet()
	if err := player.Pause(authrizedClient.Client, stopTime); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
