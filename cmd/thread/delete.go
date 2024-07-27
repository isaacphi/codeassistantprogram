package thread

import (
	"fmt"

	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete thread-name-or-id",
	Short: "Delete a thread",
	Long:  `Delete a thread`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, err := models.LoadThread(args[0])
		if err != nil {
			return fmt.Errorf("failed to load thread %q: %w", args[0], err)
		}
		if err := thread.Delete(); err != nil {
			return fmt.Errorf("failed to delete thread %q: %w", args[0], err)
		}
		fmt.Printf("Deleted thread: %s\n", thread)
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(deleteCmd)
}
