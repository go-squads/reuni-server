package users

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

type userProcessorInterface interface {
	getRepository() userRepositoryInterface
	createUserProcessor(userdata userv) error
	createUserEncryptPassword(salt string, password string) string
	loginUserProcessor(loginData userv) ([]byte, error)
	getAllUserProcessor() (string, error)
}

type userProcessor struct {
	repo userRepositoryInterface
}

func (u *userProcessor) getRepository() userRepositoryInterface {
	return u.repo
}
func (u *userProcessor) createUserProcessor(userdata userv) error {
	userstore := user{}
	userstore.Name = userdata.Name
	userstore.Username = userdata.Username
	userstore.Password = userdata.Password
	userstore.Email = userdata.Email
	return u.repo.createUser(userstore)
}

func (u *userProcessor) createUserEncryptPassword(salt string, password string) string {
	passwordStore := salt + password
	h := sha256.New()
	h.Write([]byte(passwordStore))

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func (u *userProcessor) loginUserProcessor(loginData userv) ([]byte, error) {
	return u.repo.loginUser(loginData)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (u *userProcessor) getAllUserProcessor() (string, error) {
	data, err := u.repo.getAllUser()
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}
