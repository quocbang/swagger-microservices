package configs

import "time"

type Options struct {
	ServerConfig string `short:"c" long:"server-config" description:"server configuration file" required:"true"`
}

// DBConnection set PostgreSQL connection settings.
type DBConnection struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

// Configs for
type Configs struct {
	DevMode        bool   `yaml:"development_mode"`
	UIDir          string `yaml:"ui_distribution_directory"`
	CreateUIConfig bool   `yaml:"create_ui_configuration"`

	Timeout             time.Duration `yaml:"timeout"`
	WebServiceEndpoint  string        `yaml:"web_service_endpoint"`
	PostgreSQL          DBConnection  `yaml:"postgres"`
	CorsAllowedOrigins  []string      `yaml:"cors_allowed_origins"`
	TokenExpiredSeconds int           `yaml:"token_expired_in_seconds"`
	// FunctionRolePermissions is deprecated.
	FunctionRolePermissions map[string][]string `yaml:"permissions"`
	Printers                map[string]string   `yaml:"printers"`
	FontPath                string              `yaml:"font_path"`
	MesPath                 string              `yaml:"mes_path"`
}
