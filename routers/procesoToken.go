package routers

import (
	"errors"
	"strings"

	"github.com/ascendere/micro-convocatorias/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string
var Nombre string
var Tk string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDelUniverso")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if len(claims.ID) > 0 {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
		Nombre = claims.Nombre + " " + claims.Apellidos
		Tk = tk

		return claims, true, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
