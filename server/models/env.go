package models

type Env struct {
	JWT_SECRET  string `validate:"required"`
	PASSWORD    string `validate:"required"`
	MONGODB_URI string `validate:"required,uri"`
}
