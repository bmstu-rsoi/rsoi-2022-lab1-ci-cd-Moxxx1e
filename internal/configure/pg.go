package configure

import (
	"os"
)

type PGConfig struct {
	database string
	user     string
	password string
	host     string
	port     string
}

func NewLocal() *PGConfig {
	return &PGConfig{
		database: "postgres",
		user:     "oskolganov",
		password: "postgres",
		host:     "postgres",
		port:     "5432",
	}
}

func (p *PGConfig) GetDSN() string {
	//return "postgres://postgres:postgres@localhost:5432/persons?sslmode=disable"
	//return fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
	//	p.user,
	//	p.password,
	//	p.host,
	//	p.database)
	return os.Getenv("DATABASE_URL")
	//return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.user, p.password, p.host, p.port, p.database)
}
