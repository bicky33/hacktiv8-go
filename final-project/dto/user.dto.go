package dto

type UserCreateRequest struct {
	Age      uint32 `json:"age" validate:"required,numeric,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID       uint32 `json:"id"`
	Age      uint32 `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserUpdateRequest struct {
	ID              uint32 `json:"id,omitempty"`
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
	Email           string `json:"email,omitempty" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
}

type UserUpdateResponse struct {
	ID        uint32 `json:"id"`
	Age       uint32 `json:"age"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	UpdatedAt string `json:"updated_at"`
}

type UserResponse struct {
	Status  uint32      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
