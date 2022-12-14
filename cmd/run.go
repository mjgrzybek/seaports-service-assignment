/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"
	httpServer "seaports-service-assignment/internal/application/adapters/api"
	inMemoryStore "seaports-service-assignment/internal/application/adapters/inMemoryStore"
	jsonImporter "seaports-service-assignment/internal/application/adapters/jsonImporter"
	"seaports-service-assignment/internal/application/config"
	"seaports-service-assignment/internal/application/safeExit"
	"seaports-service-assignment/internal/domain/services"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		err := setupLoger()
		if err != nil {
			log.WithError(err).Fatal("Failed to create logger")
		}

		gracefulShutdown := setupProcessShutdown()
		ctx, stopTheWorld := context.WithCancel(context.Background())
		defer stopTheWorld()

		// dependency injection:
		store := inMemoryStore.NewInMemoryStore()
		api := httpServer.NewHttpSeaportsServer(store)
		importer := jsonImporter.NewJsonImporter(store)

		svc := services.Seaports{Store: store, Api: api, Importer: importer}
		go svc.Run(ctx)

		<-gracefulShutdown
		stopTheWorld()

		select {
		case <-safeExit.Done:
			log.Info("All goroutines stopped successfully")
		case <-time.After(viper.GetDuration(config.GracefulShutdownTimeout)):
			log.Warn("Graceful shutdown timeout exceeded")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setupLoger() error {
	log.SetOutput(os.Stdout)
	lvl := viper.GetString(config.LogLevel)
	parsedLvl, err := log.ParseLevel(lvl)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse log parsedLvl: " + lvl)
	}
	log.SetLevel(parsedLvl)
	return err
}
