package crypto

import (
	"crypto/rand"
	"crypto/subtle"

	"github.com/jsee98/GAuth/gauth/internal/utils"
	"github.com/jsee98/GAuth/gauth/types"
    	"github.com/jsee98/GAuth/gauth/error"

    "golang.org/x/crypto/argon2"
)

type Argon2ParamsInternal struct {
	memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

func GeneratePasswordHash(password string, params types.Argon2Params ) ([]byte, []byte, *error.Error) {
    // Generate a cryptographically secure random salt.
    salt, err := generateRandomBytes(params.SaltLength)
    if err != nil {
        return nil, nil, err
    }
    hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)
    
    return hash, salt, nil
}

func ComparePasswordAndHash(password, hash string, salt string, params types.Argon2Params) (bool, *error.Error) {
    decodedHash, err := utils.FromBase64(hash)
    if err!=nil{
        errToRet := error.NewError("failed comparing passwords")
        errToRet.AddDetails(err.Error())
        return false, errToRet
    }
    otherHash := argon2.IDKey([]byte(password), []byte(salt), params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

    // if needed you can read on timing attacks here
    // https://www.propelauth.com/post/what-does-timing-attack-actually-mean
    if subtle.ConstantTimeCompare(decodedHash, otherHash) == 1 {
        return true, nil
    }
    return false, nil
}

func generateRandomBytes(n uint32) ([]byte, *error.Error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        errToRet := error.NewErrorEmpty()
        errToRet.AddDetails(err.Error())
        return nil, errToRet
    }
    return b, nil
}

