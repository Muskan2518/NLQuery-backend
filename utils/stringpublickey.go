package utils

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// PublicKeyToString converts an *rsa.PublicKey to a PEM-formatted string
func PublicKeyToString(pub *rsa.PublicKey) (string, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return "", err
	}

	var pemBuffer bytes.Buffer
	err = pem.Encode(&pemBuffer, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})
	if err != nil {
		return "", err
	}

	return pemBuffer.String(), nil
}
