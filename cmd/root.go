package cmd

import (
	"context"
	"fmt"
	"github.com/henomis/lingoose/llm/antropic"
	"github.com/henomis/lingoose/thread"
	cmdThread "github.com/isaacphi/codeassistantprogram/cmd/thread"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cap",
	Short: "The cap code assistant program",
	Long: `Use cap to interact with LLMs using branching threads
WIP`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get thread
		tr, err := models.GetCurrentThread()
		if err != nil {
			fmt.Errorf("Error fetching current thread\n%v", err)
		}
		fmt.Println("Thread:", tr)

		// Create LLM thread
		antropicllm := antropic.New().WithModel("claude-3-opus-20240229").WithStream(
			func(response string) {
				if response != antropic.EOS {
					fmt.Print(response)
					fmt.Print(".")
				} else {
					fmt.Println()
				}
			},
		)
		t := thread.New().AddMessage(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent("How are you?"),
			),
		)

		// Make LLM request
		err = antropicllm.Generate(context.Background(), t)
		if err != nil {
			return err
		}

		// Save answer
		lastMessage := t.LastMessage().Contents[0].AsString()
		fmt.Print(lastMessage)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmdThread.ThreadCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codeassistantprogram.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
