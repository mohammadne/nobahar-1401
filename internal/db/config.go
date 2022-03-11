package db

type Config struct {
	Driver   string `koanf:"driver"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Database string `koanf:"database"`
	SSLMode  string `koanf:"ssl_mode"`
	Timezone string `koanf:"timezone"`
	PageSize int    `koanf:"page_size"`
}
