package folderstructure

import (
	"generator/config"
	"path"
)

var (
	PathToModel = path.Join(config.Get().PathToRepositoryRoot, "internal", "models")
)
