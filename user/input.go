package user

type RegisterUserInput struct { // struct ini mewakili inputan dari user
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string
}
