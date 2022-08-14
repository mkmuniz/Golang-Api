package user

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	Done     bool   `json:"done"`
}
