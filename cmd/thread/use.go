package thread

import (
	"fmt"

	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use thread-name-or-id",
	Short: "Switch active thread",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := models.LoadThread(args[0])
		if err != nil {
			return fmt.Errorf("error loading thread: %w", err)
		}
		err = t.SetCurrent()
		if err != nil {
			return fmt.Errorf("error switching to thread: %w", err)
		}
		fmt.Println("Switched to thread:", t)
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(useCmd)
}
