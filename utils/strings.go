package utils

import (
    "crypto/rand"
    "encoding/base64"
)

//RememberTokenBytes is the numbr of bytes that will be generated.
const RememberTokenBytes = 32

//GenerateBytes function is used to generate random n bytes, or will
//return an error fi the was one, this uses crypto/rand
func GenerateBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}

//GenerateString function is used to generate a byte slic of siz nBytes and
//then it will return a string that's the base64 url encoded version.
//of that byte slice.
func GenerateString(nBytes int) (string, error) {
    b, err := GenerateBytes(nBytes)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}

//RememberToken function is used to generate a string using the remember token bytes.
func RememberToken() (string, error) {
    return GenerateString(RememberTokenBytes)
}
