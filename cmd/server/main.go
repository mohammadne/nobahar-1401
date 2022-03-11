package server

import (
	"os"
	oss "os/signal"
	"syscall"

	"github.com/mohammadne/nobahar-1401/internal/http"
	"github.com/pingcap/log"
	"github.com/spf13/cobra"
	"honnef.co/go/tools/config"
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

	go http.New().Serve()

	log.Info("exiting due to recieving an unix signal")
}
