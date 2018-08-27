package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-squads/reuni-server/helper"
)

type (
	verifiedUser struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		IAT      int64  `json:"iat"`
		Expire   int64  `json:"exp"`
	}
)

func (v *verifiedUser) Valid() bool {
	return v.ID != 0 && v.Name != "" && v.Username != "" && v.Email != "" && v.IAT != 0 && v.Expire != 0
}

const (
	createUserQuery  = "INSERT INTO users (name, username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	verifyLoginQuery = "SELECT id, name, username, email FROM users WHERE username=$1 AND password=$2"
	getAllUserQuery  = "SELECT id,username,name FROM users"
	getUserDataQuery = "SELECT id, name, username, email FROM users WHERE username=$1"
)

type userRepositoryInterface interface {
	createUser(userstore user) error
	loginUser(loginData userv) ([]byte, []byte, error)
	getAllUser() ([]user, error)
	getUserData(username string) ([]byte, error)
}

type userRepository struct {
	execer helper.QueryExecuter
}

func initRepository(execer helper.QueryExecuter) *userRepository {
	return &userRepository{
		execer: execer,
	}
}

func (u *userRepository) createUser(userstore user) error {
	userstore.CreatedAt = time.Now()
	userstore.UpdatedAt = userstore.CreatedAt
	_, err := u.execer.DoQueryRow(createUserQuery, userstore.Name, userstore.Username, userstore.Password, userstore.Email, userstore.CreatedAt, userstore.UpdatedAt)
	if err != nil {
		log.Println("CREATE USER: " + err.Error())
	}
	return err
}

func (u *userRepository) loginUser(loginData userv) ([]byte, []byte, error) {
	v := verifiedUser{}
	data, err := u.execer.DoQueryRow(verifyLoginQuery, loginData.Username, loginData.Password)
	// data.Scan(&v.ID, &v.Name, &v.Username, &v.Email)
	if err != nil {
		return nil, nil, err
	}
	if data != nil {
		v.ID = int(data["id"].(int64))
		v.Name = data["name"].(string)
		v.Username = data["username"].(string)
		v.Email = data["email"].(string)
		v.IAT = makeTimestamp()
		v.Expire = time.Now().Add(time.Minute * 1).Unix()
	}
	if !v.Valid() {
		return nil, nil, helper.NewHttpError(http.StatusUnauthorized, "Wrong username/password")
	}
	vJSON, err := json.Marshal(v)
	if err != nil {
		return nil, nil, err
	}
	vRefreshToken := &v
	vRefreshToken.Expire = time.Now().Add(time.Hour * 1).Unix()
	vRefreshTokenJSON, err := json.Marshal(vRefreshToken)
	if err != nil {
		return nil, nil, err
	}
	return vJSON, vRefreshTokenJSON, nil
}
func (u *userRepository) getAllUser() ([]user, error) {
	data, err := u.execer.DoQuery(getAllUserQuery)
	if err != nil {
		return nil, err
	}
	var users []user
	err = helper.ParseMap(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) getUserData(username string) ([]byte, error) {
	v := verifiedUser{}
	data, err := u.execer.DoQueryRow(getUserDataQuery, username)
	if err != nil {
		return nil, err
	}
	if data != nil {
		v.ID = int(data["id"].(int64))
		v.Name = data["name"].(string)
		v.Username = data["username"].(string)
		v.Email = data["email"].(string)
		v.IAT = makeTimestamp()
		v.Expire = time.Now().Add(time.Minute * 1).Unix()
	}
	return json.Marshal(v)
}
