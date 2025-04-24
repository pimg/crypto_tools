/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sharedCmd represents the shared command
var sharedCmd = &cobra.Command{
	Use:   "shared",
	Short: "Shared nonce attack",
	Long:  `Shared nonce attack as described in the paper: 'ECDSA Cracking Methods'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shared called")
	},
}

func init() {
	nonceCmd.AddCommand(sharedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sharedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sharedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
