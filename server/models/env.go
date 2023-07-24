package models

type Env struct {
	MONGODB_URI string `validate:"required,uri"`
	USERNAME    string `validate:"required"`
	PASSWORD    string `validate:"required"`
}
