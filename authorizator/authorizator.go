package authorizator

import (
	"log"

	"github.com/go-squads/reuni-server/helper"
)

type Authorizator interface {
	Authorize(userID, serviceID int, permission rune) bool
}

type authorizator struct {
	helper helper.QueryExecuter
}

func New(helper helper.QueryExecuter) Authorizator {
	return &authorizator{
		helper: helper,
	}
}

const (
	checkUserIDExistingQuery = `
	SELECT  role from organization_member where organization_id= 
	(
		select organization_id FROM services where id=$1
	)
	AND user_id=$2
	`
)

func (a *authorizator) Authorize(userID, serviceID int, permission rune) bool {
	data, err := a.helper.DoQueryRow(checkUserIDExistingQuery, serviceID, userID)
	if err != nil {
		log.Println("Authorizator:", err.Error())
		return false
	}
	role := data["role"].(string)
	switch role {
	case "Admin":
		return true
	case "Developer":
		return permission == 'w' || permission == 'r'
	case "Auditor":
		return permission == 'r'
	}

	return true
}
