// Package jwt provides and registers a custom Virgil JWT signing method.
//
// Usage:
//
//      import (
//			"github.com/dgrijalva/jwt-go"
//			"gopkg.in/virgil.v4/virgilcrypto"
//			virgilJWT github.com/ameteiko/golang-kit/jwt"
// 		)
//
// 		// Generate new custom token
//		keyPair, _ := virgilcrypto.NewKeypair()
//		tokenClaims := jwt.MapClaims{"foo": "bar"}
//		t := jwt.NewWithClaims(virgilJWT.GetVirgilSigningMethod(), tokenClaims)
//		tokenString, _ := t.SignedString(keyPair.PrivateKey())
//		println(tokenString)
//
//		// Validate the token
//		c := new(jwt.MapClaims)
// 		tt, _ := jwt.ParseWithClaims(
// 			tokenString,
// 			new(jwt.MapClaims),
// 			func(*jwt.Token) (interface{}, error) {
// 				return keyPair.PublicKey(), nil
// 			},
// 		)
//		println(tt.Valid == true)
//
//
package jwt

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/ameteiko/golang-kit/errors"
	"gopkg.in/virgil.v4/virgilcrypto"
)

//
// Virgil signing constants.
//
const (
	VirgilSigningAlgorithm = "VIRGIL"
)

// Virgil signing method instance
var VirgilSigningMethod = new(virgilSigner)

//
// init registers Virgil Security token validator.
//
func init() {
	jwt.RegisterSigningMethod(VirgilSigningAlgorithm, GetVirgilSigningMethod)
}

//
// virgilSigner is a Virgil Security implementation for the token signing.
//
type virgilSigner struct{}

//
// Verify performs token verification.
// Expects key to be virgilcrypto.PublicKey
//
func (s *virgilSigner) Verify(signingString, signature string, key interface{}) error {
	publicKey, ok := key.(virgilcrypto.PublicKey)
	if !ok {
		return errors.WithMessage(
			jwt.ErrInvalidKeyType,
			"kit.jwt@virgilSigner.Verify [key instance is not a virgilcrypto.PublicKey]",
		)
	}

	decodedSignature, err := jwt.DecodeSegment(signature)
	if nil != err {
		return errors.WrapError(
			err,
			ErrSignatureDecode,
		)
	}

	if ok, err = virgilcrypto.DefaultCrypto.Verify([]byte(signingString), decodedSignature, publicKey); !ok || nil != err {
		if !ok {
			return errors.WithMessage(
				ErrSignatureIsInvalid,
				"kit.jwt@virgilSigner.Verify [signature value is invalid]",
			)
		}

		return errors.WithMessage(
			err,
			"kit.jwt@virgilSigner.Verify [signature verification failed]",
		)
	}

	return nil
}

//
// Sign performs token signing.
// Expects key to be virgilcrypto.PrivateKey instance.
//
func (s *virgilSigner) Sign(signingString string, key interface{}) (string, error) {
	privateKey, ok := key.(virgilcrypto.PrivateKey)
	if !ok {
		return "", errors.WithMessage(
			jwt.ErrInvalidKeyType,
			"kit.jwt@virgilSigner.Sign [key instance is not a virgilcrypto.PrivateKey]",
		)

	}

	signature, err := virgilcrypto.DefaultCrypto.Sign([]byte(signingString), privateKey)
	if nil != err {
		return "", errors.WithMessage(
			err,
			"kit.jwt@virgilSigner.Sign [signing process failed]",
		)
	}

	return jwt.EncodeSegment(signature), nil
}

//
// Alg returns signer algorithm name.
//
func (s *virgilSigner) Alg() string {

	return VirgilSigningAlgorithm
}

//
// GetVirgilSigningMethod returns an instance of Virgil signing method.
//
func GetVirgilSigningMethod() jwt.SigningMethod {

	return VirgilSigningMethod
}
