package api

import (
	"github.com/ameteiko/golang-kit/errors"
	"gopkg.in/virgil.v4/virgilcrypto"
	"time"
)

//
// Account model.
//
type Account struct {
	ID           string
	APIKey       virgilcrypto.PublicKey
	Applications []Application
}

//
// Application is an Account application nested model.
//
type Application struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublicKeys  []virgilcrypto.PublicKey
}

//
// NewAccountFromResponse returns a new Account model instantiated from the AccountResponse
//
func NewAccountFromResponse(accResponse AccountResponse) (*Account, error) {
	var err error
	account := Account{ID: accResponse.ID}

	for _, appResponse := range accResponse.Applications {
		app := Application{
			ID:          appResponse.ID,
			Name:        appResponse.Name,
			Description: appResponse.Description,
		}
		app.CreatedAt, err = time.Parse("", appResponse.CreatedAt)
		if nil != err {
			return nil, errors.WithMessage(
				err,
				"kit-api@NewAccountFromResponse [unable to parse CreatedAt time parameter (%s)]",
				appResponse.CreatedAt,
			)
		}
		app.UpdatedAt, err = time.Parse("", appResponse.UpdatedAt)
		if nil != err {
			return nil, errors.WithMessage(
				err,
				"kit-api@NewAccountFromResponse [unable to parse UpdatedAt time parameter (%s)]",
				appResponse.UpdatedAt,
			)
		}
		for _, keyResponse := range appResponse.PublicKeys {
			key, err := virgilcrypto.DecodePublicKey([]byte(keyResponse))
			if nil != err {

			}
			app.PublicKeys = append(app.PublicKeys, key)
		}

		account.Applications = append(account.Applications, app)
	}

	return &account, nil
}

//func parseTime(t string) (time.Time, error) {
//	time, err := time.Parse("", t)
//	if nil != err {
//		return nil, errors.WithMessage(
//			err,
//			"kit-api@NewAccountFromResponse [unable to parse CreatedAt time parameter (%s)]",
//			appResponse.CreatedAt,
//		)
//	}
//}
///
