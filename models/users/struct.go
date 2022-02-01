package users

const Table = "users"
const PrimaryKey = "id"

type Attributes struct {
	ID              int     `json:"id" db:"id"`
	Name            string  `json:"name" db:"name"`
	Email           string  `json:"email" db:"email"`
	EmailVerifiedAt string  `json:"email_verified_at" db:"email_verified_at"`
	Password        string  `json:"password" db:"password"`
	RememberToken   *string `json:"remember_token" db:"remember_token"`
	CreatedAt       string  `json:"created_at" db:"created_at"`
	UpdatedAt       string  `json:"updated_at" db:"updated_at"`
}
