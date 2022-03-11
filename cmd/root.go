package main

import (
	"fmt"

	"github.com/mohammadne/nobahar-1401/cmd/server"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	short = "short description"
	long  = `long description`
)

func main() {
	cmd := &cobra.Command{Short: short, Long: long}
	cmd.AddCommand(server.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
