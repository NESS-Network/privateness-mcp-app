package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"encoding/json"
)

type AuthInit struct {
	PubKey   string `json:"pubkey"`   // Privateness blockchain pubkey (hex/base64; your choice)
	Nonce    string `json:"nonce"`
	Signature string `json:"signature"` // Signature over Nonce
}

func randomNonce() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil { return "", err }
	return hex.EncodeToString(b), nil
}

// TODO: Implement real verification against Privateness blockchain state
func verify(pubKey, nonce, signature string) bool {
	// Placeholder: always true for skeleton
	return pubKey != "" && nonce != "" && signature != ""
}

// Performs a very simple JSON handshake on the first stream.
// Server reads AuthInit, verifies signature; returns a JSON {ok:true, session_id:...}
func Authenticate(rw io.ReadWriter) (clientPubKey string, err error) {
	dec := json.NewDecoder(rw)
	var ai AuthInit
	if err := dec.Decode(&ai); err != nil { return "", err }
	if !verify(ai.PubKey, ai.Nonce, ai.Signature) { return "", errors.New("auth failed") }
	return ai.PubKey, nil
}
