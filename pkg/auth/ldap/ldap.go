package ldap

import (
	"errors"
	"fmt"
	"strings"

	"git.bode.fun/orders/pkg/auth"
	"github.com/go-ldap/ldap/v3"
)

// Errors
// ------------------------------------------------------------------------

var (
	ErrNumEntries error = errors.New("ldap: the entry count is not one")
	ErrWildcard         = errors.New("ldap: the username contains a wildcard character")
)

// Implementation
// ------------------------------------------------------------------------

var _ auth.Authenticator = (*authenticator)(nil)

type authenticator struct {
	conn    *ldap.Conn
	baseDN  string
	filterf string
}

func New(addr, baseDN, filterf, username, password string) (*authenticator, error) {
	conn, err := ldap.DialURL(addr)
	if err != nil {
		return nil, err
	}

	err = conn.Bind(username, password)
	if err != nil {
		return nil, err
	}

	return &authenticator{
		conn:    conn,
		baseDN:  baseDN,
		filterf: filterf,
	}, nil
}

// Public Methods
// ------------------------------------------------------------------------

func (l *authenticator) AuthenticateUser(username, password string) (bool, error) {
	if strings.Contains(username, "*") {
		return false, ErrWildcard
	}

	user, err := l.searchUser(username)
	if err != nil {
		return false, err
	}

	err = l.conn.Bind(user.DN, password)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Private Methods
// ------------------------------------------------------------------------

func (l *authenticator) searchUser(username string) (*ldap.Entry, error) {
	req := &ldap.SearchRequest{
		BaseDN: l.baseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: fmt.Sprintf(l.filterf, username),
	}

	res, err := l.conn.Search(req)
	if err != nil {
		return nil, err
	}

	if len(res.Entries) != 1 {
		return nil, ErrNumEntries
	}

	return res.Entries[0], nil
}
