/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/pimg/crypto_tools/ecdsa/p256"
	"github.com/spf13/cobra"
)

// reuseCmd represents the reuse command
var reuseCmd = &cobra.Command{
	Use:   "reuse",
	Short: "Nonce reuse attack",
	Long:  `Nonce reuse attack as described in the paper: 'ECDSA Cracking Methods'.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := new(big.Int)
		k, ok := k.SetString(opts.Nonce, 10)
		if !ok {
			fmt.Println("Invalid nonce number")
			return
		}

		s, err := hex.DecodeString(opts.Signature)
		if err != nil {
			fmt.Println("invalid signature, not a valid hex string")
			return
		}

		switch strings.ToLower(opts.Curve) {
		case "p256":
			priv, err := p256.RecoverP256KeyFromNonce(k, []byte(opts.Message), s, nil)
			if err != nil {
				fmt.Println("could not recover p256 key: " + err.Error())
				return
			}
			fmt.Printf("recovered p256 private key:\nPrivate Key: %d\nPublic Key X: %d\nPublic Key Y: %d\n", priv.D, priv.X, priv.Y)
		default:
			fmt.Println("Invalid curve")
			return
		}
	},
}

var opts struct {
	Curve     string
	Nonce     string
	Signature string
	Message   string
}

func init() {
	nonceCmd.AddCommand(reuseCmd)
	reuseCmd.Flags().StringVarP(&opts.Curve, "curve", "c", "p256", "Curve to use")
	reuseCmd.Flags().StringVarP(&opts.Nonce, "nonce", "n", "", "Leaked Nonce")
	reuseCmd.Flags().StringVarP(&opts.Signature, "signature", "s", "", "Signature encoded in hex string")
	reuseCmd.Flags().StringVarP(&opts.Message, "message", "m", "", "Message")

	if err := reuseCmd.MarkFlagRequired("curve"); err != nil {
		fmt.Println("curve flag is required")
		return
		// os.Exit(1)
	}
	if err := reuseCmd.MarkFlagRequired("nonce"); err != nil {
		fmt.Println("nonce flag is required")
		return
		// os.Exit(1)
	}
	if err := reuseCmd.MarkFlagRequired("signature"); err != nil {
		fmt.Println("signature flag is required")
		return
		// os.Exit(1)
	}
	if err := reuseCmd.MarkFlagRequired("message"); err != nil {
		fmt.Println("message flag is required")
		return
		// os.Exit(1)
	}
}
