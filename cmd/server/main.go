package server

import "github.com/spf13/cobra"

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}
	return cmd
}

func main(cmd *cobra.Command, _ []string) {}