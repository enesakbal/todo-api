package types

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
}
