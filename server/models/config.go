package models

type AppConfig struct {
	StartTime string `json:"start_time" validate:"required"`
	EndTime   string `json:"end_time" validate:"required"`
}

type ServerConfig struct {
	Port          string `validate:"required"`
	Cors          string `validate:"omitempty"`
	AdminPassword string `validate:"required"`
}

type ConfigType struct {
	App    AppConfig
	Server ServerConfig
}
