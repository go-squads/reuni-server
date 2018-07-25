package users

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
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

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
