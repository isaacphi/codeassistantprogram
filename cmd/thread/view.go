package thread

import (
	"fmt"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View active thread",
	Long:  `View active thread`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		tr, err := models.GetCurrentThread()
		if err != nil {
			fmt.Errorf("Error fetching current thread\n%v", err)
		}
		err = tr.View()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(viewCmd)
}
