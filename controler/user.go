package controler

// User content the user information
type User struct {
	ID   string
	Name string
	Pwd  string
}

var users map[string]User
