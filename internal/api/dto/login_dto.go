package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsPaid   bool   `json:"is_paid"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	IsPaid bool   `json:"is_paid"`
}
