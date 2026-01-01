package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "get information about a pdf file",
	Aliases: []string{"information"},
	Args:    cobra.ExactArgs(0),
	Run:     runFuncInfo,
}

func runFuncInfo(cmd *cobra.Command, args []string) {
	fileName, _ := cmd.Flags().GetString("file")

	fileObj, err := os.Stat(fileName)
	check(err)

	var inputFile fileInfo

	inputFile.rawFileName = fileName
	inputFile.size = fileObj.Size()

	fmt.Printf("File name is %s and size is %d\n", inputFile.rawFileName, inputFile.size)

}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().String("file", "", "file you want the information for")
}
