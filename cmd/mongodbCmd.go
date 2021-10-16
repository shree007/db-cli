package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mongodbCmd = &cobra.Command{
	Use:   "MongoDB",
	Short: "MongoDB database command will be performed",
	Long:  `MongoDB operations will be performed`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inside the mongoDB function")
		fmt.Println(args)
	},
}

func init() {
	mongodbCmd.AddCommand(mongodbCmd)
}
