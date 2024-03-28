package filesystem

import (
	"errors"
	"fmt"
	"generator/config"
	"log"
	"os"
	"path"
	"strings"
)

var (
	defPerm = os.FileMode(0755)
	root    = config.Get().PathToRepositoryRoot
)

var (
	serviceName = config.Get().ServiceName
	paths       = []string{
		PathToCmd,
		PathToApp,
		PathToConfigs,
		PathToDatabaseConfig,
		PathToPostgresConfig,
		PathToServiceConfig,
		PathToInternal,
		PathToDatabase,
		PathToRepository,
		PathToModel,
		PathToService,
		PathToMapper,
		PathToTransport,
		PathToHttp,
		PathToHttpIn,
		PathToHttpOut,
		PathToRepositoryRegistry,
	}
)

func CreateArchetypeFileSystem() error {
	for _, p := range paths {
		if err := CreateDir(p); err != nil {
			return fmt.Errorf("create archetype file system: %e", err)
		}
	}
	return nil
}

func CreateDir(name string) error {
	if !strings.HasPrefix(name, root) {
		name = path.Join(root, name)
	}
	if err := os.Mkdir(name, defPerm); err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		log.Fatal(err)
	}
	return nil
}
