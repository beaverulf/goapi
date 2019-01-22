package models

// Success response
// swagger:response ok
type swaggerSuccessResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Code int `json:"code"`
	}
}

// Success response
// swagger:response ok
type swaggerCryptoServiceResp struct {
	Body struct {
		// Name of something.
		Name string `json:"name,omitempty"`
		// Cryptographic functions available from the service.
		Functions []string `json:"functions,omitempty"`
	}
}

type CryptoService struct {
	Name      string   `json:"name,omitempty"`
	Functions []string `json:"functions,omitempty"`
}
