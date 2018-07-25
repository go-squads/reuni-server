package users

import (
	"encoding/json"
	"time"

	context "github.com/go-squads/reuni-server/appcontext"
)

type (
	verifiedUser struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		IAT      int64  `json:"iat"`
	}
)

const createUserQuery = "INSERT INTO users (name, username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
const verifyLoginQuery = "SELECT id, name, username, email FROM users WHERE username=$1 AND password=$2"

func createUser(userstore user) error {
	userstore.CreatedAt = time.Now()
	userstore.UpdatedAt = userstore.CreatedAt
	db := context.GetDB()
	_, err := db.Exec(createUserQuery, userstore.Name, userstore.Username, userstore.Password, userstore.Email, userstore.CreatedAt, userstore.UpdatedAt)

	return err
}

func loginUser(loginData userv) ([]byte, error) {
	v := verifiedUser{}
	db := context.GetDB()
	err := db.QueryRow(verifyLoginQuery, loginData.Username, loginData.Password).Scan(&v.ID, &v.Name, &v.Username, &v.Email)
	v.IAT = makeTimestamp()
	if err != nil {
		return nil, err
	}

	return json.Marshal(v)
}
