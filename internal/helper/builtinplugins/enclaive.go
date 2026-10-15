package builtinplugins

import (
	"fmt"

	"github.com/enclaive/openbao/external/vhsm"
	"github.com/openbao/openbao/sdk/v2/helper/consts"
	"github.com/openbao/openbao/sdk/v2/logical"
)

func init() {
	for name, plugin := range vhsm.Plugins() {
		if err := Registry.Set(name, plugin.Type, plugin.Factory, plugin.DbFactory); err != nil {
			panic(err)
		}
	}
}

func (r *registry) Set(name string, pluginType consts.PluginType, factory logical.Factory, dbFactory BuiltinFactory) error {
	switch pluginType {
	case consts.PluginTypeCredential:
		_, ok := r.credentialBackends[name]
		if ok {
			return fmt.Errorf("credential backend %q already exists", name)
		}
		r.credentialBackends[name] = credentialBackend{Factory: factory}
	case consts.PluginTypeSecrets:
		_, ok := r.logicalBackends[name]
		if ok {
			return fmt.Errorf("logical backend %q already exists", name)
		}
		r.logicalBackends[name] = logicalBackend{Factory: factory}
	case consts.PluginTypeDatabase:
		_, ok := r.databasePlugins[name]
		if ok {
			return fmt.Errorf("database backend %q already exists", name)
		}
		r.databasePlugins[name] = databasePlugin{Factory: dbFactory}
	default:
		return fmt.Errorf("unknown plugin type: %s", pluginType)
	}
	return nil
}
