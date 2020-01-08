package security

import (
	"crypto/rand"
	"fmt"

	"github.com/ahmetcanozcan/reves/sockets/messages"
)

var onAuthenticateSocket func(string, string) (string, bool) = func(username string, password string) (string, bool) {
	return generateSocketID(), true
}

//IsAuthenticationActive :
var IsAuthenticationActive bool = false

//AuthenticateSocket :
func AuthenticateSocket(payload messages.Payload) (string, bool) {
	if !IsAuthenticationActive {
		return generateSocketID(), true
	}
	username, ok := payload["username"]
	if ok {
		password, ok := payload["password"]
		if ok {
			return onAuthenticateSocket(username, password)
		}
	}
	return "", false
}

//OnAuthenticateSocket :
func OnAuthenticateSocket(f func(string, string) (string, bool)) {
	onAuthenticateSocket = f
}

func generateSocketID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
