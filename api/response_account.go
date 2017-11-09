package api

//
// Account response object.
//
type AccountResponse struct {
	ID           string                `json:"id"`
	APIKey       string                `json:"api_key"`
	Applications []ApplicationResponse `json:"applications"`
}

//
// ApplicationResponse is an Account application nested model.
//
type ApplicationResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	PublicKeys  []string `json:"public_keys,omitempty"`
}
