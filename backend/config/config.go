package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Name          string `env:"APP_NAME"`
	Version       string `env:"APP_VERSION"`
	DisableLogin  bool   `env:"DISABLE_LOGIN" env-default:"false"`
	DisableSignup bool   `env:"DISABLE_SIGNUP" env-default:"false"`
	BaseUrl       string `env:"APP_BASE_URL"`
}

type Server struct {
	Port    string `env:"SERVER_PORT" default:"3000"`
	Timeout int    `env:"SERVER_TIMEOUT_SECONDS" env-default:"10"`
}

type Postgres struct {
	PoolMax       int    `env:"PG_POOL_MAX" env-default:"10"`
	URL           string `env:"PG_URL" env-default:"postgres://postgres:postgres@localhost:5432/m_content_development?sslmode=disable"`
	RunMigrations bool   `env:"PG_RUN_MIGRATIONS" env-default:"false"`
	MaxConnection int    `env:"PG_MAX_CONNECTION" env-default:"100"`
	MacIdle       int    `env:"PG_MAX_IDLE" env-default:"10"`
	MaxLifetime   int    `env:"PG_MAX_LIFETIME" env-default:"2"`
}

type Redis struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

type JWT struct {
	Secret    string `env:"JWT_SECRET"`
	Algorithm string `env:"JWT_ALGORITHM"`
	Expires   int    `env:"JWT_EXPIRES_IN_SECONDS"`
}

type Log struct {
	Level string `env:"LOG_LEVEL" env-default:"info"`
}

type OAuthProvider struct {
	ClientID    string `json:"client_id" split_words:"true"`
	Secret      string `json:"secret"`
	CallbackURL string `json:"redirect_uri" split_words:"true"`
	URL         string `json:"url"`
	Enabled     bool   `json:"enabled"`
}

type FiberConfig struct {
}

type Config struct {
	App      App
	Server   Server
	Postgres Postgres
	Redis    Redis
	JWT      JWT
	Log      Log
	Fiber    fiber.Config
	Cors     cors.Config
}

var Cfg Config

func LoadConfig() *Config {
	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		panic(err)
	}

	Cfg.Fiber = fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	}

	Cfg.Cors = cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-With",
		AllowCredentials: true,
	}

	return &Cfg
}
