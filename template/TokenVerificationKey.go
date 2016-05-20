package template

import (
	"encoding/base64"
	"fmt"
	"os"
)

type TokenVerificationKey struct {
	KeyValue string `xml:"KeyValue"`
	Type     string `xml:"type,attr"`
}

func (t *TokenVerificationKey) DescerializeKey() []byte {
	key, err := base64.StdEncoding.DecodeString(t.KeyValue)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return key
}
