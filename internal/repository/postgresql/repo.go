package postgresql

type Config struct {
	Host     string `koanf:"host"`
	Port     uint   `koanf:"port"`
	User     string `koanf:"username"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
}
