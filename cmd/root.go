package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/philborlin/zsuite/internal/workflow"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zsuite",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior: run the workflow
		err := workflow.Workflow()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
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

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the enrollment workflow",
	Long: `Run the enrollment workflow to copy activities between record books.
This will guide you through selecting an enrollment, choosing source and target record books,
and copying activities between them.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := workflow.Workflow()
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	},
}

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix [activityID]",
	Short: "Fix issues with record books",
	Long: `Fix various issues with record books, such as data inconsistencies,
missing activities, or other problems that may occur.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		activityID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid activity ID '%s'. Please provide a valid integer.\n", args[0])
			os.Exit(1)
		}

		err = workflow.Fix(activityID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Fix command completed successfully for activity ID: %d\n", activityID)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zsuite.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add subcommands
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(fixCmd)
}
