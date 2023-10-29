package User

// User defines the structure for user details
// swagger:User Login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger: Message
type Message struct {
	Message string
}