package models

type ConfigType struct {
	Port          string `validate:"required"`
	Cors          string `validate:"omitempty"`
	AdminPassword string `validate:"required"`
}
