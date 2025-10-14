package config

type Config struct {
	Host string `yaml:"host" json:"host"`
	Port string
}

func LoadFromFile() *Config {
	return nil
}

func LoadFromEnvironment() (*Config, error) {
	return new(Config), nil
}
