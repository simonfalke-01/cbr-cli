package netscape

import (
	"github.com/simonfalke-01/cbr-cli/kooky/internal/cookies"
)

type CookieStore struct {
	cookies.DefaultCookieStore
	IsStrictBool bool
}

// strict netscape cookies.txt format
func (s *CookieStore) IsStrict() bool {
	return s != nil && s.IsStrictBool
}

var _ cookies.CookieStore = (*CookieStore)(nil)
