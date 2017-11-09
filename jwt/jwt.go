package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"strings"

	"github.com/ameteiko/golang-kit/errors"
	"gopkg.in/virgil.v4/virgilcrypto"
)

//
// Virgil signing constants.
//
const (
	VirgilSigningAlgorithm = "VIRGIL"
)

//
// init registers Virgil Security token validator.
//
func init() {
	jwt.RegisterSigningMethod(VirgilSigningAlgorithm, GetVirgilSigningMethod)
}

//
// VirgilSigner is a Virgil Security implementation for the token signing.
//
type VirgilSigner struct{}

//
// Verify performs token verification.
// Expects key to be virgilcrypto.PublicKey
//
func (s *VirgilSigner) Verify(token, signature string, key interface{}) error {
	publicKey, ok := key.(virgilcrypto.PublicKey)
	if !ok {
		return errors.WithMessage(
			jwt.ErrInvalidKeyType,
			"kit.jwt@VirgilSigner.Verify [key instance is not a virgilcrypto.PublicKey]",
		)
	}

	decodedSignature, err := jwt.DecodeSegment(signature)
	if nil != err {
		return errors.WrapError(
			err,
			ErrSignatureDecode,
		)
	}

	parts := strings.Split(token, ".")
	signedData := strings.Join(parts[0:2], ".")
	if ok, err = virgilcrypto.DefaultCrypto.Verify([]byte(signedData), decodedSignature, publicKey); !ok || nil != err {
		if !ok {
			return errors.WithMessage(
				ErrSignatureIsInvalid,
				"kit.jwt@VirgilSigner.Verify [signature value is invalid]",
			)
		}

		return errors.WithMessage(
			err,
			"kit.jwt@VirgilSigner.Verify [signature verification failed]",
		)
	}

	return nil
}

//
// Sign performs token signing.
// Expects key to be virgilcrypto.PrivateKey instance.
//
func (s *VirgilSigner) Sign(token string, key interface{}) (string, error) {
	privateKey, ok := key.(virgilcrypto.PrivateKey)
	if !ok {
		return "", errors.WithMessage(
			jwt.ErrInvalidKeyType,
			"kit.jwt@VirgilSigner.Sign [key instance is not a virgilcrypto.PrivateKey]",
		)

	}

	signature, err := virgilcrypto.DefaultCrypto.Sign([]byte(token), privateKey)
	if nil != err {
		return "", errors.WithMessage(
			err,
			"kit.jwt@VirgilSigner.Sign [signing process failed]",
		)
	}

	return jwt.EncodeSegment(signature), nil
}

//
// Alg returns signer algorithm name.
//
func (s *VirgilSigner) Alg() string {

	return VirgilSigningAlgorithm
}

//
// GetVirgilSigningMethod returns an instance of Virgil signing method.
//
func GetVirgilSigningMethod() jwt.SigningMethod {

	return new(VirgilSigner)
}
