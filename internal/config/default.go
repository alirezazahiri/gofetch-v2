package config

const EnvPrefix = "GOFETCH_V2_"

var defaultConfig = map[string]any{
	"env": "development",
	"http_server.port": 8080,
	"postgresql.host": "localhost",
	"postgresql.port": 5432,
	"postgresql.username": "postgres",
	"postgresql.password": "postgres",
	"postgresql.dbname": "postgres",
}
