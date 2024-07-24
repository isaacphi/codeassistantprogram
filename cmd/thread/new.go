package thread

import (
	"fmt"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new thread",
	Long:  `Create a new thread`,
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) == 1 {
			name = args[0]
		}
		thread, err := models.NewThread(name)
		if err != nil {
			return fmt.Errorf("Error creating new thread\n%v", err)
		}
		thread.Save(config.DataDirectory)

		thread, err = models.SetCurrentThread(name)
		if err != nil {
			return err
		}
		fmt.Println("Created and switched to thread:", thread)

		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(newCmd)
}
