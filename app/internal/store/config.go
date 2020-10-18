package store

// Config ...
type Config struct {
	DatabaseURL  string `toml:"database_url"`
	DatabaseName string `toml:"database"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
