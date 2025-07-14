package common

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthorMessage struct {
	FullName string `json:"full_name"`
	ID       string `json:"id"`
}
