package config

type Config struct {
	Port int64
	PostgresDatabase
}

type PostgresDatabase struct {
	Username string
	Password string
	Host     string
	Port     int64
	DBname   string
}
