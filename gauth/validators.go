package gauth

import "github.com/jsee98/GAuth/gauth/types"
import "github.com/jsee98/GAuth/gauth/error"

/*
	Checks if the auth config is valid for initializing a client.
	If not returns a custom list of errors stating what is missing and what might be needed.
*/

func ValidateConfig(cfg types.Config) *error.Error {
	errObj := error.NewError("Config Validation Failed")

	// check DB Config
	dbErrors := cfg.DBConfig.Validate()
	for _, err := range dbErrors {
		errObj.AddDetails(err)
		errObj.SetPanic()
		return errObj
	}

	

	return nil
}
