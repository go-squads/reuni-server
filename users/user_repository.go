package users

import (
	"time"

	context "github.com/go-squads/reuni-server/appcontext"
)

const createUserQuery = "INSERT INTO users (name, username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

func createUser(userstore user) error {
	userstore.CreatedAt = time.Now()
	userstore.UpdatedAt = userstore.CreatedAt
	db := context.GetDB()
	_, err := db.Exec(createUserQuery, userstore.Name, userstore.Username, userstore.Password, userstore.Email, userstore.CreatedAt, userstore.UpdatedAt)

	return err
}
