package config

import (
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver"
	"github.com/alirezazahiri/gofetch-v2/internal/repository/postgresql"
)

type RepositoryConfig struct {
	Postgres postgresql.Config `koanf:"postgres"`
}

type Config struct {
	Env        string             `koanf:"env"`
	HttpServer httpserver.Config   `koanf:"http_server"`
	Repository RepositoryConfig   `koanf:"repository"`
}
