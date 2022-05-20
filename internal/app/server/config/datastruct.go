package config

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

type Config struct {
	DB DB
}
