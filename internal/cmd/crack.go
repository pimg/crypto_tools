/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// crackCmd represents the crack command
var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "Various commands for cracking ECDSA.",
	Long:  `Various commands for cracking ECDSA. Based on the 'ECDSA Cracking Methods' paper: https://arxiv.org/pdf/2504.07265`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crack called")
	},
}

func init() {
	ecdsaCmd.AddCommand(crackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
