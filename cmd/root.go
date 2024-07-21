package cmd

import (
	// "context"
	// "fmt"
	"os"

	// "github.com/henomis/lingoose/llm/antropic"
	// "github.com/henomis/lingoose/thread"
	"github.com/spf13/cobra"

	"github.com/isaacphi/codeassistantprogram/cmd/thread"
)

var rootCmd = &cobra.Command{
	Use:   "cap",
	Short: "The cap code assistant program",
	Long: `Use cap to interact with LLMs using branching threads
WIP`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("args", args)

	// 	antropicllm := antropic.New().WithModel("claude-3-opus-20240229").WithStream(
	// 		func(response string) {
	// 			if response != antropic.EOS {
	// 				fmt.Print(response)
	// 			} else {
	// 				fmt.Println()
	// 			}
	// 		},
	// 	)

	// 	t := thread.New().AddMessage(
	// 		thread.NewUserMessage().AddContent(
	// 			thread.NewTextContent("How are you?"),
	// 		),
	// 	)

	// 	err := antropicllm.Generate(context.Background(), t)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(t)
	// },
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
	rootCmd.AddCommand(thread.ThreadCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codeassistantprogram.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
