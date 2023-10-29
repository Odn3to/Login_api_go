package jwt_func

// swagger: Verfica
type Verfica struct {
	Token string `json:"token"`
}

// swagger: RetornoVerificar
type RetornoVerificar struct {
	ValidToken bool 
}