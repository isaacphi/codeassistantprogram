package thread

import (
	"fmt"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete thread-name-or-id",
	Short: "Delete a thread",
	Long:  `Delete a thread`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		thread, err := models.LoadThread(args[0], config.DataDirectory)
		if err != nil {
			return fmt.Errorf("Couldn't find thread %v\n%v", args[0], err)
		}
		thread.Delete(config.DataDirectory)
		fmt.Println("Deleted thread:", thread)
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(deleteCmd)
}
