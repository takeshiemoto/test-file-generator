package dto

type SignUpRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
