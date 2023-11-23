package fcm

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

type tokenProvider struct {
	tokenSource oauth2.TokenSource
}

func newTokenProviderFromBytes(ctx context.Context, jsonKey []byte) (*tokenProvider, error) {
	if len(jsonKey) == 0 {
		return nil, errors.New("empty")
	}

	cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
	if err != nil {
		return nil, errors.Wrapf(err, "fcm: failed to get JWT config for the firebase.messaging scope")
	}

	ts := cfg.TokenSource(ctx)
	return &tokenProvider{
		tokenSource: ts,
	}, nil
}

func newTokenProvider(ctx context.Context, credentialsLocation string) (*tokenProvider, error) {
	jsonKey, err := ioutil.ReadFile(credentialsLocation)
	if err != nil {
		return nil, errors.Wrapf(err, "fcm: failed to read credentials file at: '%s'", credentialsLocation)
	}

	cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
	if err != nil {
		return nil, errors.Wrapf(err, "fcm: failed to get JWT config for the firebase.messaging scope")
	}

	ts := cfg.TokenSource(ctx)
	return &tokenProvider{
		tokenSource: ts,
	}, nil
}

// token is safe for use from multiple go routines. It will request a token if
// one does not exist or is expired.
func (src *tokenProvider) token() (string, error) {
	token, err := src.tokenSource.Token()
	if err != nil {
		return "", errors.Wrapf(err, "fcm: failed to generate Bearer token")
	}

	return token.AccessToken, nil
}
