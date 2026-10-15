package external

import (
	"github.com/openbao/openbao/sdk/v2/helper/consts"
	"github.com/openbao/openbao/sdk/v2/logical"
)

type Plugin struct {
	Type      consts.PluginType
	Factory   logical.Factory
	DbFactory func() (any, error)
}
