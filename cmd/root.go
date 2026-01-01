package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func runFunc(cmd *cobra.Command, args []string) {

	fmt.Println("hello from cli")

}

var rootCmd = &cobra.Command{
	Use:   "pdffy",
	Short: "Use this cli for your pdfs",
	Run:   runFunc,
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("an error occured" + err.Error())
		os.Exit(1)
	}
}
