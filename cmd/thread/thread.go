package thread

import (
	"github.com/spf13/cobra"
)

var ThreadCmd = &cobra.Command{
	Use:   "thread",
	Short: "Thread subcommands",
	Long:  "Manage threads",
}

func init() {
}
