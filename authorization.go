package clickup

// AuthorizationMethod is the authorization method that should be used for integrations with ClickUp
type AuthorizationMethod interface {
	// ApplyAuthorization applies the AuthorizationMethod to a given Request
	ApplyAuthorization(r *Request)
}

// PersonalTokenAuthorization is an authorization method that uses personal tokens
type PersonalTokenAuthorization struct {
	token string
}

// NewPersonalTokenAuthorization creates a new PersonalTokenAuthorization
func NewPersonalTokenAuthorization(token string) *PersonalTokenAuthorization {
	return &PersonalTokenAuthorization{token: token}
}

// ApplyAuthorization applies the AuthorizationMethod to a given Request
func (a *PersonalTokenAuthorization) ApplyAuthorization(r *Request) {
	r.SetHeader(header_authorization, a.token)
}
