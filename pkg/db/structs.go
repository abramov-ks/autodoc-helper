package db

type DatabaseConfig struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
