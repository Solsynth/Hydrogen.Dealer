package main

import (
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"os"
	"os/signal"
	"syscall"

	pkg "git.solsynth.dev/hydrogen/dealer/pkg/internal"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/grpc"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/server"
	"github.com/robfig/cron/v3"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func main() {
	// Configure settings
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")

	// Load settings
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Err(err).Msg("An error occurred when loading settings.")
	}

	// Set up external services
	if err := services.SetupFirebase(); err != nil {
		log.Warn().Err(err).Msg("An error occurred when setup firebase, firebase notification push is unavailable...")
	}
	if err := services.SetupAPNS(); err != nil {
		log.Warn().Err(err).Msg("An error occurred when setup APNs, apple notification push is unavailable...")
	}

	// Set up tasks queue consumers
	for idx := 0; idx < max(1, viper.GetInt("performance.notification_deliver.worker_count")); idx++ {
		go services.ConsumeDeliveryTasks()
	}

	// Server
	go server.NewServer().Listen()

	// Grpc Server
	go grpc.NewServer().Listen()

	// Configure timed tasks
	quartz := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(&log.Logger)))
	quartz.Start()

	// Messages
	log.Info().Msgf("Dealer v%s is started...", pkg.AppVersion)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msgf("Passport v%s is quitting...", pkg.AppVersion)

	quartz.Stop()
}
