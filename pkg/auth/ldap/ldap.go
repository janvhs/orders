// TODO: Write tests and documentation
package ldap

import (
	"errors"
	"fmt"
	"strings"

	"git.bode.fun/orders/pkg/auth"
	"github.com/go-ldap/ldap/v3"
)

// Errors
// TODO: Wrap errors. Use a custom error type or something
// ------------------------------------------------------------------------

var (
	ErrConn       error = errors.New("ldap: connection error")
	ErrBind             = errors.New("ldap: bind error")
	ErrSearch           = errors.New("ldap: search error")
	ErrWildcard         = errors.New("ldap: the username contains a wildcard character")
	ErrNotFound         = errors.New("ldap: user not found")
	ErrNumEntries       = errors.New("ldap: the entry count is not one")
)

// Implementation
// ------------------------------------------------------------------------

var _ auth.Authenticator = (*Authenticator)(nil)

type Authenticator struct {
	Conn    *ldap.Conn
	BaseDN  string
	Filterf string
}

// TODO: Remove constructor?
func New(addr, baseDN, filterf, bindDN, bindPassword string) (*Authenticator, error) {
	conn, err := connect(addr)
	if err != nil {
		return nil, err
	}

	err = bind(conn, bindDN, bindPassword)
	if err != nil {
		return nil, err
	}

	return &Authenticator{
		Conn:    conn,
		BaseDN:  baseDN,
		Filterf: filterf,
	}, nil
}

// Public Methods
// ------------------------------------------------------------------------

// TODO: Check if this always returns an error
// TODO: Unbind before closing?
// TODO: Export
func (l *Authenticator) close() error {
	return l.Conn.Close()
}

// TODO: Do I really need the bool, since an error indicates not authenticated?
// TODO: Return the GUID, DN or Entry for the user
func (l *Authenticator) AuthenticateUser(username, password string) (bool, error) {
	if strings.Contains(username, "*") {
		return false, ErrWildcard
	}

	users, err := searchUsers(l.Conn, username, l.BaseDN, l.Filterf)
	if err != nil {
		return false, err
	}

	usersCount := len(users)

	if usersCount == 0 {
		return false, ErrNotFound
	}

	if usersCount != 1 {
		return false, ErrNumEntries
	}

	user := users[0]

	// INFO: ATTENTION this is a check for POSITIVE nil-equality
	// TODO: Does this replace the currently bound user? And therefore need a new connection?
	if err := bind(l.Conn, user.DN, password); err == nil {
		// TODO: Get the GUID from the user
		return true, nil
	}

	return false, err
}

// Private Methods
// ------------------------------------------------------------------------

// Private functions
// ------------------------------------------------------------------------

func connect(addr string) (*ldap.Conn, error) {
	conn, err := ldap.DialURL(addr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConn, err)
	}

	return conn, nil
}

func bind(conn *ldap.Conn, userDN, password string) error {
	err := conn.Bind(userDN, password)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrBind, err)
	}

	return nil
}

func searchUsers(conn *ldap.Conn, username string, baseDN string, filterf string) ([]*ldap.Entry, error) {
	req := &ldap.SearchRequest{
		BaseDN: baseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: fmt.Sprintf(filterf, username),
	}

	res, err := conn.Search(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSearch, err)
	}

	return res.Entries, nil
}
