/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ecdsaCmd represents the ecc command
var ecdsaCmd = &cobra.Command{
	Use:   "ecdsa",
	Short: "Various operations related to ECDSA",
	Long:  `Various opererations releated to ECDSA, for exloring various concepts and weaknesses releated to ECDSA.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ecdsa called")
	},
}

func init() {
	rootCmd.AddCommand(ecdsaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ecdsaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ecdsaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
