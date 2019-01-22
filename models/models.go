package models

//Endpoint describes the content of an endpoint
type Endpoint struct {
	Route    string `json:"route,omitempty"`
	Function string `json:"description,omitempty"`
}

//CipherObject contains information about the encrypted text as well as the ciphertext
type CipherObject struct {
	Cipher     Cipher `json:"cipher,omitempty"`
	Key        string `json:"key,omitempty"`
	Encoding   string `json:"encoding,omitempty"`
	Ciphertext string `json:"ciphertext,omitempty"`
}

//Cipher contains metadata about the algorithm etc.
type Cipher struct {
	Algorithm string `json:"algorithm,omitempty"`
	Entropy   string `json:"entropy,omitempty"`
	Mode      string `json:"mode,omitempty"`
}
