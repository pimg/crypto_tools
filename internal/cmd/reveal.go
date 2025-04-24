/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// revealCmd represents the reveal command
var revealCmd = &cobra.Command{
	Use:   "reveal",
	Short: "Nonce reveal attack",
	Long:  `Nonce reveal attack as described in the paper: 'ECDSA Cracking Methods'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reveal called")
	},
}

func init() {
	nonceCmd.AddCommand(revealCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// revealCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// revealCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
