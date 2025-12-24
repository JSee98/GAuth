package types

import (
	"context"

	"github.com/jsee98/GAuth/gauth/db"
)

type Token struct {
	AccessToken  string // JWT base64 encoded
	RefreshToken string // JWT base64 encoded
}

type ConsistencyLevelEnum int

const EVENTUAL_CONSISTENCY ConsistencyLevelEnum = 0 // replica is read if available
const HIGH_CONSISTENCY ConsistencyLevelEnum = 1     // master is read

type Config struct {
	DBConfig             db.DBConfig
	Argon2Config         Argon2Params
	PasswordRequirements PasswordRequirements
	DenyList             bool
	ConsistencyLevel     ConsistencyLevelEnum
}

type GAuthI interface {
	SignInWithEmail(ctx *context.Context, email, password string) (Token, error)
	SignUpWithEmail(ctx *context.Context, email, password string) (Token, error)
}
