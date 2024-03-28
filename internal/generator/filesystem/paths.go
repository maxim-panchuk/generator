package filesystem

import (
	"generator/config"
	"path"
)

var p = config.Get().PathToRepositoryRoot

var (
	PathToCmd                = path.Join(p, "cmd")
	PathToApp                = path.Join(p, "cmd", "app")
	PathToConfigs            = path.Join(p, "configs")
	PathToDatabaseConfig     = path.Join(p, "configs", "database")
	PathToPostgresConfig     = path.Join(p, "configs", "database", "postgres")
	PathToServiceConfig      = path.Join(p, "configs", "servicesettings")
	PathToInternal           = path.Join(p, "internal")
	PathToDatabase           = path.Join(p, "internal", "database")
	PathToRepository         = path.Join(p, "internal", "database", "repositories")
	PathToModel              = path.Join(p, "internal", "models")
	PathToService            = path.Join(p, "internal", "service")
	PathToMapper             = path.Join(p, "internal", "mapper")
	PathToTransport          = path.Join(p, "internal", "transport")
	PathToHttp               = path.Join(p, "internal", "transport", "http")
	PathToHttpIn             = path.Join(p, "internal", "transport", "http", "in")
	PathToHttpOut            = path.Join(p, "internal", "transport", "http", "out")
	PathToRepositoryRegistry = path.Join(p, "internal", "database", "registry")
)
