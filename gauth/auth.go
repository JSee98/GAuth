package gauth

import (
	"github.com/jsee98/GAuth/gauth/db"
	"github.com/jsee98/GAuth/gauth/internal/repository"
	"github.com/jsee98/GAuth/gauth/internal/service"
	"github.com/jsee98/GAuth/gauth/types"
)

type GAuth struct {
	DBClient          db.DBClientI
	CredentialService *service.CredentialService
}

func NewGAuthClient(cfg *types.Config) (*GAuth, error) {
	dbClient, err := db.NewClient(cfg.DBConfig)
	if err != nil {
		return nil, err
	}

	credRepo := &repository.CredentialRepo{
		DBClient: dbClient,
	}

	credService := service.NewCredentialService(credRepo, &cfg.PasswordRequirements, &cfg.Argon2Config)

	return &GAuth{
		DBClient:          dbClient,
		CredentialService: credService,
	}, nil
}
