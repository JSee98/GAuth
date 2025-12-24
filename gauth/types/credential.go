package types

import (
	uuid "github.com/gofrs/uuid/v5"
)

/*
	Defines the email password credentials
*/

type Credential struct {
	ID uuid.UUID
	Email string
	Password string
	Salt string
}

type PasswordRequirements struct {
	MinPasswordLength int
	NeedCapitalLetters bool
	NeedSpecialCharacters bool
	NeedNumbers bool
	MinSaltLength int
}