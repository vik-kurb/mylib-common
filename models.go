package common

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthorMessage struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Action   string `json:"action"`
}
