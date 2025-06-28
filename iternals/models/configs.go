package models

type Configs struct {
	AuthParams     AuthParams     `json:"auth_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
	LogParams      LogParams      `json:"log_params"`
}

type AuthParams struct {
	JwtSecretKey  string `json:"jwt_secret_key"`
	JwtTtlMinutes int    `json:"jwt_ttl_minutes"`
}

type LogParams struct {
	Directory        string `json:"log_directory"`
	Info             string `json:"log_info"`
	Error            string `json:"log_error"`
	Warn             string `json:"log_warn"`
	Debug            string `json:"log_debug"`
	MaxSizeMegabytes int    `json:"max_size_megabytes"`
	MaxBackups       int    `json:"max_backups"`
	MaxAgeDays       int    `json:"max_age_days"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"local_time"`
}

type AppParams struct {
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
	AppVersion string `json:"app_version"`
	PortRun    string `json:"port_run"`
	GinMode    string `json:"gin_mode"`
}

type PostgresParams struct {
	User         string `json:"user"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database"`
}
