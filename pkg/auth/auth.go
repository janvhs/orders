package auth

type Authenticator interface {
	AuthenticateUser(username, password string) (bool, error)
}
