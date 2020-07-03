package user

//User structure
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewUser(username string, password string) *User {

	return &User{
		Username: username,
		Password: password,
	}
}
