package main

type User struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Points int    `json:"points"`
}

type SignUpRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SignUpResponse struct {
	Message string `json:"message"`
}

type SignInRequest struct {
	ID string `json:"id"`
}

type SignInResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Points  int    `json:"points"`
}

type PointsResponse struct {
	Points int `json:"points"`
}
