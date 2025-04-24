package p256

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverP256KeyFromNonce(t *testing.T) {
	curve := elliptic.P256()
	order := curve.Params().N

	message := []byte("Hello World")
	hash := sha256.Sum256(message)
	m := new(big.Int).SetBytes(hash[:])

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	// Pick a random nonce k
	k, err := rand.Int(rand.Reader, FloatToBigInt(math.Pow(2, 127)))
	if err != nil {
		t.Fatal(err)
	}

	// Manually sign to include insecure Nonce
	rx, _ := curve.ScalarBaseMult(k.Bytes())
	r := new(big.Int).Mod(rx, order)

	if r.Sign() == 0 {
		t.Fatal("r is zero")
	}

	kInv := ModInverse(k, order)
	s := new(big.Int).Mul(privKey.D, r)
	s.Add(s, m)
	s.Mul(s, kInv)
	s.Mod(s, order)

	t.Log("Sig 1 r,s:", r, s)
	sigBytes, err := asn1.Marshal(ecdsaSignature{R: r, S: s})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Sig bytes: %x\n", sigBytes)

	recreatedPrivKey, err := RecoverP256KeyFromNonce(k, message, sigBytes, nil)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, privKey.D.Bytes(), recreatedPrivKey.D.Bytes())
	assert.Equal(t, privKey.X.Bytes(), recreatedPrivKey.X.Bytes())
	assert.Equal(t, privKey.Y.Bytes(), recreatedPrivKey.Y.Bytes())

	fmt.Printf("Recovered Private key: %s\n", recreatedPrivKey.D.String())
	fmt.Printf("Public key X: %d\nPublic key Y: %d\n", recreatedPrivKey.X, recreatedPrivKey.Y)
	fmt.Printf("From Nonce: %s\n", k.String())
	fmt.Printf("Signature: %s\n", hex.EncodeToString(sigBytes))

	k2 := new(big.Int)
	k2, ok := k2.SetString(k.String(), 10)
	if !ok {
		t.Fatal("failed to set k")
	}

	s2, err := hex.DecodeString(hex.EncodeToString(sigBytes))
	if err != nil {
		t.Fatal(err)
	}
	anotherKey, err := RecoverP256KeyFromNonce(k2, message, s2, nil)
	assert.NoError(t, err)
	assert.Equal(t, recreatedPrivKey.D.Bytes(), anotherKey.D.Bytes())
}

func TestRecoverP256KeyFromNonceWithPubKey(t *testing.T) {
	curve := elliptic.P256()
	order := curve.Params().N

	message := []byte("Hello World")
	hash := sha256.Sum256(message)
	m := new(big.Int).SetBytes(hash[:])

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	pubKey := privKey.PublicKey

	// Pick a random nonce k
	k, err := rand.Int(rand.Reader, FloatToBigInt(math.Pow(2, 127)))
	if err != nil {
		t.Fatal(err)
	}

	// Manually sign to include insecure Nonce
	rx, _ := curve.ScalarBaseMult(k.Bytes())
	r := new(big.Int).Mod(rx, order)

	if r.Sign() == 0 {
		t.Fatal("r is zero")
	}

	kInv := ModInverse(k, order)
	s := new(big.Int).Mul(privKey.D, r)
	s.Add(s, m)
	s.Mul(s, kInv)
	s.Mod(s, order)

	t.Log("Sig 1 r,s:", r, s)
	sigBytes, err := asn1.Marshal(ecdsaSignature{R: r, S: s})
	if err != nil {
		t.Fatal(err)
	}

	recreatedPrivKey, err := RecoverP256KeyFromNonce(k, message, sigBytes, &pubKey)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, privKey.D.Bytes(), recreatedPrivKey.D.Bytes())
	assert.Equal(t, privKey.X.Bytes(), recreatedPrivKey.X.Bytes())
	assert.Equal(t, privKey.Y.Bytes(), recreatedPrivKey.Y.Bytes())
}
