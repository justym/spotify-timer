package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
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
	var stopTime = minutes

	if len(args) > 0 {
		t, err := util.AtoTime(args[0])
		if err != nil || t == -1 {
			log.Fatal(err, t)
		}
		stopTime = t
	}
	fmt.Println(color.YellowString("[TIME]: %s", stopTime.String()))

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
