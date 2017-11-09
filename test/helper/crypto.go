package helper

import (
	"github.com/ameteiko/golang-kit/errors"
	"gopkg.in/virgil.v4/virgilcrypto"
)

//
// GenerateKeys returns a newly generated private and public keys.
//
func GenerateKeys() (virgilcrypto.PrivateKey, virgilcrypto.PublicKey) {
	kp, err := virgilcrypto.NewKeypair()
	if nil != err {
		errors.ReportStartupErrorAndExit(err)
	}

	return kp.PrivateKey(), kp.PublicKey()
}

//
// Sign returns the data signature.
//
func Sign(data []byte, key virgilcrypto.PrivateKey) []byte {
	s, _ := virgilcrypto.Signer.Sign(data, key)

	return s
}
