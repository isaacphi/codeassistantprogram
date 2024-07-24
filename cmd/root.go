package cmd

import (
	"bufio"
	"fmt"
	"github.com/isaacphi/codeassistantprogram/cmd/thread"
	"github.com/isaacphi/codeassistantprogram/internal/config"
	"github.com/isaacphi/codeassistantprogram/internal/llm"
	"github.com/isaacphi/codeassistantprogram/internal/models"
	"github.com/isaacphi/codeassistantprogram/internal/ui"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "cap",
	Short: "The cap code assistant program",
	Long:  `Use cap to interact with LLMs using branching threads. WIP`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input, err := getInput()
		if err != nil {
			return err
		}

		tr, err := models.GetCurrentThread()
		if err != nil {
			fmt.Errorf("Error fetching current thread\n%v", err)
		}

		userMessage, err := models.NewMessage(input, "user")
		if err != nil {
			return err
		}
		userMessage.Save(config.DataDirectory)
		tr.AddMessage(userMessage)
		tr.Save(config.DataDirectory)

		llm := llm.New("claude-3-opus-20240229")
		llm.LoadThread(tr)

		response, err := llm.GenerateResponse()
		if err != nil {
			return err
		}

		assistantMessage, err := models.NewMessage(response, "assistant")
		if err != nil {
			return err
		}
		assistantMessage.Save(config.DataDirectory)
		tr.AddMessage(assistantMessage)
		tr.Save(config.DataDirectory)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(thread.ThreadCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codeassistantprogram.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getInput() (string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped to stdin
		reader := bufio.NewReader(os.Stdin)
		var builder strings.Builder
		_, err := io.Copy(&builder, reader)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(builder.String()), nil
	}

	// No data piped, prompt user for input
	input, err := ui.GetInput()
	if err != nil {
		return "", err
	}
	return input, nil
}
