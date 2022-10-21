package sub

type Http struct {
	Schema   string
	Method   string
	Host     string
	Port     int
	Path     string
	BaseAuth AuthPair
}

type AuthPair struct {
	Username string
	Password string
}
