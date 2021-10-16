package clickup

import "fmt"

type AuthorizationMethod interface {
	fmt.Stringer
}

type PersonalTokenAuthorization struct {
	Token string
}

func NewPersonalTokenAuthorization(token string) *PersonalTokenAuthorization {
	return &PersonalTokenAuthorization{
		Token: token,
	}
}

func (a *PersonalTokenAuthorization) String() string { return a.Token }
