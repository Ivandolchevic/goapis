// Package authenticationUtil provides some helpers for the authentication process
package authenticationUtil

import (
	"strings"

	"github.com/nu7hatch/gouuid"
)

// CreateToken generate a bearer token of
func CreateToken() (string, error) {
	u, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	token := strings.Replace(u.String(), "-", "", -1)

	return token, nil
}
