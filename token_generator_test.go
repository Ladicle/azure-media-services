package util

import (
	"github.com/Ladicle/azure-media-services/template"
	"testing"
	"time"
)

func TestGenerateTestToken(t *testing.T) {
	expectedToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1cm46dGVzdCIsImV4cCI6IjE0MjAwMjgwMDQiLCJpc3MiOiJodHRwOi8vdGVzdGFjcy5jb20iLCJ1cm46bWljcm9zb2Z0OmF6dXJlOm1lZGlhc2VydmljZXM6Y29udGVudGtleWlkZW50aWZpZXIiOiI5Y3NhZDZhNS0yMDVkLTQwNGQtODE1Ni0xMDExMTI3NmNlMTcifQ.mrSuvITU_uZ3sJYXmVw5WkRFpKnzUZOkk5gQpWBY6sk"
	templateString := `<TokenRestrictionTemplate xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/Azure/MediaServices/KeyDelivery/TokenRestrictionTemplate/v1"><AlternateVerificationKeys /><Audience>urn:test</Audience><Issuer>http://testacs.com</Issuer><OpenIdConnectDiscoveryDocument i:nil="true" /><PrimaryVerificationKey i:type="SymmetricVerificationKey"><KeyValue>Tesadfoiwpeonlksaf==</KeyValue></PrimaryVerificationKey><RequiredClaims><TokenClaim><ClaimType>urn:microsoft:azure:mediaservices:contentkeyidentifier</ClaimType><ClaimValue i:nil="true" /></TokenClaim></RequiredClaims><TokenType>JWT</TokenType></TokenRestrictionTemplate>`
	tokenKey := "nb:kid:UUID:9csad6a5-205d-404d-8156-10111276ce17"

	expiration := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
	template := template.NewTokenRestrictionTemplate(templateString)
	token := GenerteTestToken(template, tokenKey, expiration)

	if token != expectedToken {
		t.Error("Generated token is invalid.")
	}
}
