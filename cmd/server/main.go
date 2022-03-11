package server

import (
	"os"
	oss "os/signal"
	"syscall"

	"log"

	"github.com/mohammadne/nobahar-1401/internal/config"
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
		panic(map[string]interface{}{"msg": errString, "err": err})
	}

	signalChannel := make(chan os.Signal, 1)
	oss.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	jwt := jwt.New(cfg.JWT)
	go http.New(cfg.HTTP, jwt).Serve()

	exitReason := "exiting due to recieving an unix signal, signal: %s\n"
	log.Printf(exitReason, (<-signalChannel).String())
}
