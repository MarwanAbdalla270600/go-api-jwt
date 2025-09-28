package entity

type User struct {
	Id             string `db:"id"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Role           string `db:"role"`
	CreatedAt      string `db:"created_at"`
}

type UserDTO struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName"  binding:"required"`
	LastName  string `json:"lastName"   binding:"required"`
	Email     string `json:"email"      binding:"required,email"`
	Password  string `json:"password"   binding:"required,min=8"`
	Role      string `json:"role" binding:"required,oneof=ADMIN USER"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  UserDTO
}
