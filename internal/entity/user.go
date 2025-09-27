package entity

type User struct {
	Id             string
	FirstName      string
	LastName       string
	Email          string
	HashedPassword string
}

type UserDTO struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName"  binding:"required"`
	LastName  string `json:"lastName"   binding:"required"`
	Email     string `json:"email"      binding:"required,email"`
	Password  string `json:"password"   binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
