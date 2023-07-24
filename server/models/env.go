package models

type Env struct {
	JWT_SECRET     string `validate:"required"`
	ADMIN_USERNAME string `validate:"required"`
	ADMIN_PASSWORD string `validate:"required"`
	MONGODB_URI    string `validate:"required,uri"`
}
