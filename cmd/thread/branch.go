package thread

import (
	"fmt"

	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch <new_thread_name>",
	Short: "Branch from existing thread to create new thread",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) == 1 {
			name = args[0]
		}

		currentThread, err := models.GetCurrentThread()
		if err != nil {
			return fmt.Errorf("error fetching current thread: %w", err)
		}

		newThread, err := models.NewThread(name)
		newThread.MessageIDs = currentThread.MessageIDs
		if err != nil {
			return fmt.Errorf("error creating new thread:%w", err)
		}
		newThread.Save()

		err = newThread.SetCurrent()
		if err != nil {
			return fmt.Errorf("error setting current thread: %w", err)
		}

		fmt.Println("Created and switched to thread:", newThread)

		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(branchCmd)
}
