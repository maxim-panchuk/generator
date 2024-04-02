package filesystem

import (
	"errors"
	"fmt"
	"generator/config"
	"generator/internal/generator/filesystem/defaultfiles"
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
	files = map[string][]byte{
		path.Join(PathToPostgresConfig, "config.go"): []byte(defaultfiles.PostgresConfig),
	}
)

func CreateArchetypeFileSystem() error {
	for _, p := range paths {
		if err := CreateDir(p); err != nil {
			return fmt.Errorf("create archetype file system: %e", err)
		}
	}

	for p, content := range files {
		if err := CreateFileWithContent(p, content); err != nil {
			return err
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

func CreateFileWithContent(path string, content []byte) error {
	wrF, err := createFile(path)
	defer wrF.Close()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := wrF.Write(content); err != nil {
		log.Fatal(err)
	}

	return nil
}

func createFile(name string) (*os.File, error) {
	if !strings.HasPrefix(name, root) {
		name = path.Join(root, name)
	}
	file, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("create/open file: %w", err)
	}
	return file, nil
}
