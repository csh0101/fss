package main

import (
	"fss/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCMD = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	rootCMD.AddCommand(cmd.NewRunCMD())
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
