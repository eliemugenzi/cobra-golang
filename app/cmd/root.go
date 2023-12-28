package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "Hugo",
	Short: "Hugo is an awesome tool out here",
	Long: "Ewana you're missing out on this Hugo thing",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
		cmd.Help()
	},
}

func Execute() {
	fmt.Println("Cobra execution starts...")
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
	os.Exit(1)
  }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}