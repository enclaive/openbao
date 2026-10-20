package jwt

import jwtauth "github.com/openbao/openbao/v2/internal/builtin/credential/jwt"

var Factory = jwtauth.Factory

type AuthBackend = jwtauth.Backend
