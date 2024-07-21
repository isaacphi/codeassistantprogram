package thread

import (
	"fmt"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use thread-name-or-id",
	Short: "Switch active thread",
	Long:  `Switch active thread`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, err := models.SetCurrentThread(args[0])
		if err != nil {
			return err
		}
		fmt.Println("Switched to thread:", thread)
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(useCmd)
}
