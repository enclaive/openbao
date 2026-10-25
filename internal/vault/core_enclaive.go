package vault

import (
	"github.com/openbao/openbao/sdk/v2/logical"
)

func (c *Core) DevMode() bool {
	return c.storageType == "inmem"
}

func (c *Core) Storage() logical.Storage {
	return c.barrier
}
