/*
Copyright © 2022 Nedim Akar <nedim.akar53411@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "list all your ssh keys.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}
