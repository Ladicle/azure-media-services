package util

import (
	"fmt"
	"github.com/Ladicle/azure-media-services/template"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

const CONTENT_KEY_ID = "urn:microsoft:azure:mediaservices:contentkeyidentifier"

// Return AES token
func GenerteTestToken(template *template.TokenRestrictionTemplate, tokenKey string, expiration time.Time) string {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CONTENT_KEY_ID: tokenKey[len("nb:kid:UUID:"):],
		"iss": template.Issuer,
		"aud": template.Audience,
		"exp": strconv.FormatInt(expiration.Unix(), 10),
	})

	token, err := t.SignedString(template.PrimaryVerificationKey.DescerializeKey())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return "Bearer " + token
}
