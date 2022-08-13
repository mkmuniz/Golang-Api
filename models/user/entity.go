package user

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"user"`
	Password string `json:"password"`
	Done     bool   `json:"done"`
}
