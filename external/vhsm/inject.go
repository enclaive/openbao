package vhsm

import (
	"context"

	"github.com/openbao/openbao/v2/external"
)

func Plugins() map[string]external.Plugin {
	return nil
}

type Core interface{}

func PostUnseal(_ context.Context, _ Core) {}
