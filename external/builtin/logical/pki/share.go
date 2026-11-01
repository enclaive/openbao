package pki

import "github.com/openbao/openbao/v2/internal/builtin/logical/pki"

var Factory = pki.Factory

type Backend = pki.BackendEnclaive
