package authn

import (
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

// Authn specific claims in addition to standard JWT claims
type AuthnClaims struct {
	jwt.Claims

	// Time when the authentication occurred
	AuthTime jwt.NumericDate `json:"auth_time,omitempty"`
}

// Provides a JSON Web Key from a Key ID
// Wanted to use function signature from go-jose.v2
// but that would make us lose error information
type jwkProvider interface {
	Key(kid string) ([]jose.JSONWebKey, error)
}

// Extracts verified in-built claims from a jwt idToken
type jwtClaimsExtractor interface {
	GetVerifiedClaims(idToken string) (*AuthnClaims, error)
}
