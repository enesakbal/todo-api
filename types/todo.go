package types

type Todo struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Complete    bool
	Priority    string
}
