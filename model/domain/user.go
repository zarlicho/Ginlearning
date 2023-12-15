package domain

type LoginData struct {
	Gmail     string
	Passwords string
	Token     string
	Id        int
}

type RegisterData struct {
	Gmail     string
	Names     string
	Passwords string
	Id        int
}
