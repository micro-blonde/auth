package auth

import (
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/profile"
)

type AccountProfile[acc account.Model, prof profile.Model] struct {
	Account account.Account[acc]
	Profile profile.Profile[prof]
}
