package web

type LoginRequest struct {
	Gmail     string `validate:"required,min=1,max=200"`
	Passwords string `validate:"required,min=1,max=200"`
}
type RegisterRequest struct {
	Gmail     string `validate:"required,min=1,max=200"`
	Names     string `validate:"required,min=1,max=200"`
	Passwords string `validate:"required,min=1,max=200"`
}
