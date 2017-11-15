package jwt

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"

	"encoding/base64"
	"github.com/ameteiko/golang-kit/test/helper"
	"gopkg.in/virgil.v4/virgilcrypto"
)

//
// Testing constants.
//
const (
	InvalidSignatureValue = "Some signature"
	TokenValue            = "Some token value"
)

//
// VirgilSigner.Verify :: invalid key instance :: returns an error.
//
func TestVerifyWithInvalidKey(t *testing.T) {
	signer := new(virgilSigner)
	key := "Some invalid key value"

	err := signer.Verify(TokenValue, InvalidSignatureValue, key)

	assert.Error(t, err)
	helper.AssertError(t, jwt.ErrInvalidKeyType, err)
}

//
// VirgilSigner.Verify :: invalid key instance :: returns an error.
//
func TestVerifyWithAPrivateKeyInsteadOfAPublicKey(t *testing.T) {
	signer := new(virgilSigner)
	privateKey, _ := helper.GenerateKeys()

	err := signer.Verify(TokenValue, InvalidSignatureValue, privateKey)

	assert.Error(t, err)
	helper.AssertError(t, jwt.ErrInvalidKeyType, err)
}

//
// VirgilSigner.Verify :: not a base64 signature value :: returns an error.
//
func TestVerifyWithIncorrectSignatureValue(t *testing.T) {
	signer := new(virgilSigner)
	_, publicKey := helper.GenerateKeys()

	err := signer.Verify(TokenValue, InvalidSignatureValue, publicKey)

	assert.Error(t, err)
	helper.AssertError(t, ErrSignatureDecode, err)
}

//
// VirgilSigner.Verify :: invalid signature value :: returns an error.
//
func TestVerifyWithInvalidSignatureValue(t *testing.T) {
	signer := new(virgilSigner)
	_, publicKey := helper.GenerateKeys()
	signature := base64.StdEncoding.EncodeToString([]byte(InvalidSignatureValue))

	err := signer.Verify(TokenValue, signature, publicKey)

	assert.Error(t, err)
	helper.AssertError(t, jwt.ErrSignatureInvalid, err)
}

//
// VirgilSigner.Verify :: valid signature encoded in a wrong way :: returns an error.
//
func TestVerifyWithASignatureEncodedInAWrongWay(t *testing.T) {
	signer := new(virgilSigner)
	privateKey, publicKey := helper.GenerateKeys()
	rawSignature, _ := virgilcrypto.Signer.Sign([]byte(TokenValue), privateKey)
	signature := base64.StdEncoding.EncodeToString(rawSignature)

	err := signer.Verify(TokenValue, signature, publicKey)

	assert.Error(t, err)
}

//
// VirgilSigner.Verify :: valid signature value :: passes.
//
func TestVerifyWithAValidSignature(t *testing.T) {
	signer := new(virgilSigner)
	privateKey, publicKey := helper.GenerateKeys()
	rawSignature, _ := virgilcrypto.Signer.Sign([]byte(TokenValue), privateKey)
	signature := jwt.EncodeSegment(rawSignature)

	err := signer.Verify(TokenValue, signature, publicKey)

	assert.Empty(t, err)
}

//
// VirgilSigner.Sign :: invalid key instance :: returns an error.
//
func TestSignWithInvalidKey(t *testing.T) {
	signer := new(virgilSigner)
	key := "Some invalid key value"

	_, err := signer.Sign(TokenValue, key)

	assert.Error(t, err)
	helper.AssertError(t, jwt.ErrInvalidKeyType, err)
}

//
// VirgilSigner.Sign :: valid key instance :: passes.
//
func TestSignWithValidKey(t *testing.T) {
	signer := new(virgilSigner)
	privateKey, _ := helper.GenerateKeys()
	rawSignature, _ := virgilcrypto.Signer.Sign([]byte(TokenValue), privateKey)
	expectedSignature := jwt.EncodeSegment(rawSignature)

	s, err := signer.Sign(TokenValue, privateKey)

	assert.Empty(t, err)
	assert.Equal(t, expectedSignature, s)
}

//
// VirgilSigner.Alg :: without parameters :: returns a valid alg value.
//
func TestAlgReturnsAlgorithmValue(t *testing.T) {
	s := new(virgilSigner)

	alg := s.Alg()

	assert.Equal(t, VirgilSigningAlgorithm, alg)
}

//
// VirgilSigner.Sign :: valid key instance :: passes.
//
func TestSignAndVerifyWorkflow(t *testing.T) {
	signer := new(virgilSigner)
	privateKey, publicKey := helper.GenerateKeys()

	s, errSign := signer.Sign(TokenValue, privateKey)
	errVerify := signer.Verify(TokenValue, s, publicKey)

	assert.Empty(t, errSign)
	assert.Empty(t, errVerify)
}
