package vhsm

import (
	"context"

	"github.com/hashicorp/cli"
	"github.com/openbao/openbao/v2/external"
)

func Plugins() map[string]external.Plugin {
	return nil
}

type Core interface{}

func PostUnseal(_ context.Context, _ Core) {}

func BeforeServer(_ cli.Ui) bool {
	return true
}
