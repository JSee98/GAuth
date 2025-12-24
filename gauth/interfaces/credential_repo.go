package interfaces

import (
	"context"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/jsee98/GAuth/gauth/error"
	"github.com/jsee98/GAuth/gauth/types"
)

type CredentialRepoI interface {
	CreateCredential(ctx *context.Context, cred *types.Credential) (*types.Credential, *error.Error)
	GetCredentialByEmail(ctx *context.Context, email string) (*types.Credential, *error.Error)
	GetCredentialByUUID(ctx *context.Context, uuid uuid.UUID) (*types.Credential, *error.Error)
}
