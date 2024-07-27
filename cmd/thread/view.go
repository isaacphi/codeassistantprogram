package thread

import (
	"fmt"

	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View active thread",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		tr, err := models.GetCurrentThread()
		if err != nil {
			return fmt.Errorf("error fetching current thread: %w", err)
		}
		err = tr.View()
		if err != nil {
			return fmt.Errorf("error viewing thread: %w", err)
		}
		return nil

	},
}

func init() {
	ThreadCmd.AddCommand(viewCmd)
}
