package utils

 import b64 "encoding/base64"

func ToBase64(data []byte) string {
	return b64.StdEncoding.EncodeToString(data)
}

func FromBase64(data string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(data)
}