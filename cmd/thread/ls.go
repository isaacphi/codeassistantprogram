package thread

import (
	"fmt"
	"os"

	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List threads",
	Long:  `List threads`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		threads, err := models.ListThreads()
		if os.IsNotExist(err) {
			return fmt.Errorf("no threads found: %w", err)
		} else if err != nil {
			return fmt.Errorf("error listing threads: %w", err)
		}
		for _, thread := range threads {
			fmt.Printf("%v\n", thread)
		}
		return nil
	},
}

func init() {
	ThreadCmd.AddCommand(lsCmd)
}
