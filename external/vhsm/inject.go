package vhsm

import (
	"context"

	"github.com/enclaive/openbao/external/vhsm/command"
	"github.com/hashicorp/cli"
	"github.com/openbao/openbao/v2/external"
)

func Plugins() map[string]external.Plugin {
	return nil
}

type Core any

func PostUnseal(_ context.Context, _ Core) {}

func BeforeServer(_ cli.Ui) bool {
	return true
}

func Commands(_ map[string]cli.CommandFactory, _ func() command.Command) {}
