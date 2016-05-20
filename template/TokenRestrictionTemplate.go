package template

import (
	"encoding/xml"
	"fmt"
	"os"
)

type TokenRestrictionTemplate struct {
	XMLName                        xml.Name                       `xml:"TokenRestrictionTemplate"`
	AlternateVerificationKey       AlternateVerificationKeys      `xml:",AlternateVerificationKeys"`
	Audience                       string                         `xml:"Audience"`
	Issuer                         string                         `xml:"Issuer"`
	OpenIdConnectDiscoveryDocument OpenIdConnectDiscoveryDocument `xml:",OpenIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         TokenVerificationKey           `xml:",PrimaryVerificationKey"`
	RequiredClaims                 RequiredClaims                 `xml:",RequiredClaims"`
	TokenType                      string                         `xml:",TokenType"`
	XMLNameSpace                   string                         `xml:"xmlns,attr"`
	XMLNameSpaceXsl                string                         `xml:"xmlns:i,attr"`
}

func NewTokenRestrictionTemplate(template string) *TokenRestrictionTemplate {
	t := new(TokenRestrictionTemplate)
	err := xml.Unmarshal([]byte(template), &t)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return t
}
