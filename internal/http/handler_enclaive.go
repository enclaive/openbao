package http

import (
	"net/http"

	"github.com/enclaive/openbao/external/vhsm"
	"github.com/openbao/openbao/v2/internal/vault"
)

func init() {
	additionalRoutes = func(mux *http.ServeMux, core *vault.Core) {
		vhsm.RegisterRoutes(mux, core)
	}
}
