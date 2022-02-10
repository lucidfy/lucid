package passwordresets

const Table = "password_resets"
const PrimaryKey = "id"

type Paginate struct {
	Total       int
	PerPage     int
	CurrentPage int
	LastPage    int
	Data        []Attributes
}

type Attributes struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Token     string `json:"token" db:"remember_token"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
