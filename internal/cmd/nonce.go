/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nonceCmd represents the nonce command
var nonceCmd = &cobra.Command{
	Use:   "nonce",
	Short: "ECDSA nonce based attacks",
	Long:  `Various noce based attacks related to ECDSA.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nonce called")
	},
}

func init() {
	crackCmd.AddCommand(nonceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nonceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nonceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
