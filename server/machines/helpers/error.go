package helpers

import (
	"github.com/khulnasoft/meshery/server/models"
	"github.com/layer5io/meshkit/errors"
)

var (
	ErrAutoRegisterCode = "meshery-server-1219"
)

func ErrAutoRegister(err error, connType string) error {
	return errors.New(ErrAutoRegisterCode, errors.Alert, []string{}, []string{}, []string{}, []string{})
}

func IsConnectionUpdateErr(err error) bool {
	return errors.GetCode(err) == models.ErrUpdateConnectionStatusCode
}
