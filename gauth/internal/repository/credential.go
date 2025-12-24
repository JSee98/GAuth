package repository

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/jsee98/GAuth/gauth/db"
	"github.com/jsee98/GAuth/gauth/error"
	"github.com/jsee98/GAuth/gauth/types"
)

type CredentialRepo struct {
	DBClient db.DBClientI
}

func (r *CredentialRepo) CreateCredential(ctx *context.Context, cred *types.Credential) (*types.Credential, *error.Error) {
	query := "INSERT INTO gauth.credentials (id, email, password, salt) VALUES (?, ?, ?, ?)"

	_, err := r.DBClient.GetConnection().ExecContext(*ctx, query, cred.ID, cred.Email, cred.Password, cred.Salt)
	if err != nil {
		e := error.NewError("failed to create credential")
		e.AddDetails(err.Error())
		return nil, e
	}

	return cred, nil
}

func (r *CredentialRepo) GetCredentialByEmail(ctx *context.Context, email string) (*types.Credential, *error.Error) {
	query := "SELECT id, email, password, salt FROM gauth.credentials WHERE email = ?"
	row := r.DBClient.GetConnection().QueryRowContext(*ctx, query, email)

	var cred types.Credential
	err := row.Scan(&cred.ID, &cred.Email, &cred.Password, &cred.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			e := error.NewError("credential not found")
			return nil, e
		}
		e := error.NewError("failed to get credential by email")
		e.AddDetails(err.Error())
		return nil, e
	}

	return &cred, nil
}

func (r *CredentialRepo) GetCredentialByUUID(ctx *context.Context, id uuid.UUID) (*types.Credential, *error.Error) {
	query := "SELECT id, email, password, salt FROM gauth.credentials WHERE id = ?"
	row := r.DBClient.GetConnection().QueryRowContext(*ctx, query, id)

	var cred types.Credential
	err := row.Scan(&cred.ID, &cred.Email, &cred.Password, &cred.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			e := error.NewError("credential not found")
			return nil, e
		}
		e := error.NewError("failed to get credential by uuid")
		e.AddDetails(err.Error())
		return nil, e
	}

	return &cred, nil
}
