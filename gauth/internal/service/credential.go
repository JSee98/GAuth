package service

import (
	"context"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/gofrs/uuid/v5"
	"github.com/jsee98/GAuth/gauth/error"
	"github.com/jsee98/GAuth/gauth/interfaces"
	"github.com/jsee98/GAuth/gauth/internal/crypto"
	"github.com/jsee98/GAuth/gauth/internal/utils"
	"github.com/jsee98/GAuth/gauth/types"
)

type CredentialService struct {
	repo             interfaces.CredentialRepoI
	passRequirements *types.PasswordRequirements
	argon2Params     *types.Argon2Params
}

func NewCredentialService(repo interfaces.CredentialRepoI, passRequirements *types.PasswordRequirements, argon2Params *types.Argon2Params) *CredentialService {
	return &CredentialService{
		repo:             repo,
		passRequirements: passRequirements,
		argon2Params:     argon2Params,
	}
}

func (c *CredentialService) CreateCredentials(ctx *context.Context, email, password string) (*types.Credential, *error.Error) {
	credObj := &types.Credential{}

	err := c.validatePasswordRequirements(password)
	if err != nil {
		err.SetMessage("password failed requirements")
		return nil, err
	}

	if !utils.IsValidEmail(email) {
		return nil, error.NewError("invalid email")
	}

	hashedPassword, salt, err := crypto.GeneratePasswordHash(password, *c.argon2Params)
	if err != nil {
		return nil, err
	}

	credObj.Email = email
	credObj.Password = utils.ToBase64(hashedPassword)
	credObj.ID = uuid.NewV5(uuid.NamespaceDNS, email)
	credObj.Salt = utils.ToBase64(salt)

	return c.repo.CreateCredential(ctx, credObj)
}

func (c *CredentialService) GetCredentialByEmail(ctx *context.Context, email string) (*types.Credential, *error.Error) {
	if !utils.IsValidEmail(email) {
		return nil, error.NewError("invalid email")
	}
	return c.repo.GetCredentialByEmail(ctx, email)
}

func (c *CredentialService) GetCredentialByUUID(ctx *context.Context, id uuid.UUID) (*types.Credential, *error.Error) {
	return c.repo.GetCredentialByUUID(ctx, id)
}

func (c *CredentialService) validatePasswordRequirements(password string) *error.Error {
	err := error.NewErrorEmpty()

	if c.passRequirements.MinPasswordLength > utf8.RuneCountInString(password) {
		err.AddDetails(fmt.Sprintf("length of password less than %d", c.passRequirements.MinPasswordLength))
	}

	if c.passRequirements.NeedCapitalLetters {
		// [A-Z] matches any capital letter
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		if !hasUpper {
			err.AddDetails("password needs capital letters")
		}
	}

	if c.passRequirements.NeedNumbers {
		// [A-Z] matches any capital letter
		hasUpper := regexp.MustCompile(`[0-9]`).MatchString(password)
		if !hasUpper {
			err.AddDetails("password needs a number")
		}
	}

	if c.passRequirements.NeedSpecialCharacters {
		// [A-Z] matches any capital letter
		hasUpper := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)
		if !hasUpper {
			err.AddDetails("password needs special character")
		}
	}

	return err.Error()
}
