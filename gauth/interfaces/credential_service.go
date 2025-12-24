package interfaces

import (
	"context"

	"github.com/jsee98/GAuth/gauth/error"
	"github.com/jsee98/GAuth/gauth/types"
)

type CredentialServiceI interface{
	CreateCredentials(ctx *context.Context, email, password string) (*types.Credential, *error.Error)
	ValidateCredentials(ctx *context.Context, email, password string) (*types.Credential, *error.Error)
}