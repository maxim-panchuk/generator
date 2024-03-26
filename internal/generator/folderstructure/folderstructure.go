package folderstructure

import "generator/config"

var (
	serviceName = config.Get().ServiceName
	paths       = []string{
		PathToModel,
	}
)
