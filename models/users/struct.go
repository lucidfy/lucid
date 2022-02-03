package users

const Table = "users"
const PrimaryKey = "id"

type Paginate struct {
	Total       int
	PerPage     int
	CurrentPage int
	LastPage    int
	Data        []Attributes
}

type Attributes struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	EmailVerifiedAt string  `json:"email_verified_at"`
	Password        string  `json:"password"`
	RememberToken   *string `json:"remember_token"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
