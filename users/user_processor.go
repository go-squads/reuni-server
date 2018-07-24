package users

import (
	"crypto/sha256"
	"encoding/base64"
)

func createUserProcessor(userdata userv) error {
	userstore := user{}
	userstore.Name = userdata.Name
	userstore.Username = userdata.Username
	userstore.Password = createUserEncryptPassword(userdata.Username, userdata.Password)
	userstore.Email = userdata.Email

	return createUser(userstore)
}

func createUserEncryptPassword(salt string, password string) string {
	passwordStore := salt + password
	h := sha256.New()
	h.Write([]byte(passwordStore))

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
