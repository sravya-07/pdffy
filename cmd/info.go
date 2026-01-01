package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "get information about a pdf file",
	Aliases: []string{"information"},
	Run:     runFuncInfo,
}

func runFuncInfo(cmd *cobra.Command, args []string) {
	fmt.Println("inside info sub command")
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
