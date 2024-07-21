package thread

import (
	"fmt"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
	"os"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List threads",
	Long:  `List threads`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		threads, err := models.ListThreads(config.DataDirectory)
		if os.IsNotExist(err) {
			return fmt.Errorf("No threads found\n")
		} else if err != nil {
			return fmt.Errorf("Error listing threads\n%v", err)
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
