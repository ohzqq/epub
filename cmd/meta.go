package cmd

import (
	"fmt"

	"github.com/ohzqq/epub"
	"github.com/spf13/cobra"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fn := args[0]
		book := epub.OpenDir(fn)
		fmt.Printf("meta %v\n", book.Opf)
	},
}

func init() {
	rootCmd.AddCommand(metaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
