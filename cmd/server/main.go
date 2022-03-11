package server

import (
	"os"
	oss "os/signal"
	"syscall"

	"log"

	"github.com/mohammadne/nobahar-1401/internal/config"
	"github.com/mohammadne/nobahar-1401/internal/db"
	"github.com/mohammadne/nobahar-1401/internal/http"
	"github.com/mohammadne/nobahar-1401/internal/jwt"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}
	return cmd
}

func main(cmd *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		errString := "failed to load configs"
		log.Panic(errString, err)
	}

	signalChannel := make(chan os.Signal, 1)
	oss.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	db, err := db.New(cfg.DB)
	if err != nil {
		log.Panic(err)
	}

	if err := db.Migrate(); err != nil {
		log.Panic(err)
	}

	jwt := jwt.New(cfg.JWT)
	go http.New(cfg.HTTP, db, jwt).Serve()

	exitReason := "exiting due to recieving an unix signal, signal: %s\n"
	log.Printf(exitReason, (<-signalChannel).String())
}
