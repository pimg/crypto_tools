package p256

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/asn1"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"os"
)

func ModInverse(a, n *big.Int) *big.Int {
	inv := new(big.Int).ModInverse(a, n)
	if inv == nil {
		slog.Error("no modular inverse exists")
		os.Exit(1)
	}
	return inv
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

type ecdsaSignature struct {
	R, S *big.Int
}

func RecoverP256KeyFromNonce(nonce *big.Int, message, signature []byte, pubkey *ecdsa.PublicKey) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256() // TODO fixed consider moving to struct and convert function to method
	order := curve.Params().N
	k := nonce

	// extract s, r from signature
	var sig ecdsaSignature
	_, err := asn1.Unmarshal(signature, &sig)
	if err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("Extracted from signature: S: %d R: %d", sig.R, sig.R))

	// re-hash the message
	hash := sha256.Sum256(message)
	m := new(big.Int).SetBytes(hash[:])

	// Recover key from nonce: (r⁻¹ * ((k * s) - m)) % order
	rInv := ModInverse(sig.R, order)
	ks := new(big.Int).Mul(k, sig.S)
	ks.Sub(ks, m)
	tryPriv := new(big.Int).Mul(rInv, ks)
	tryPriv.Mod(tryPriv, order)

	slog.Debug("found key: " + tryPriv.String())

	// create ecdsa.PrivateKey including the ecdsa.Public key
	testPx, testPy := curve.ScalarBaseMult(tryPriv.Bytes())
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     testPx,
			Y:     testPy,
		},
		D: tryPriv,
	}

	// Verify recovered key generates the same public key
	if pubkey != nil {
		if testPx.Cmp(pubkey.X) == 0 && testPy.Cmp(pubkey.Y) == 0 {
			slog.Debug(fmt.Sprintf("The private key has been found: %d", tryPriv))
		} else {
			slog.Debug("Failed to recover the private key.")
			return nil, errors.New("failed to recover the private key")
		}
	}
	return privateKey, nil
}
