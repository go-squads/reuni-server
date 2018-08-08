package organization

import (
	"errors"

	"github.com/go-squads/reuni-server/helper"
)

type repository interface {
	createNewOrganization(organization_name string) (int64, error)
	addUser(organizationId, userId int64, role string) error
	deleteUser(organizationId, userId int64) error
}

type mainRepository struct {
	execer helper.QueryExecuter
}

const (
	createNewOrganizationQuery = "INSERT INTO organization(name) VALUES($1) RETURNING id"
	addUserQuery               = "INSERT INTO organization_member(organization_id, user_id, role) VALUES ($1,$2,$3) "
	deleteUserFromGroupQuery   = "DELETE FROM organization_member where organization_id=$1 and user_id=$2"
)

func initRepository(execer helper.QueryExecuter) *mainRepository {
	return &mainRepository{
		execer: execer,
	}
}

func (s *mainRepository) createNewOrganization(organizationName string) (int64, error) {
	data, err := s.execer.DoQueryRow(createNewOrganizationQuery, organizationName)
	if err != nil {
		return 0, err
	}
	id, ok := data["id"].(int64)
	if !ok {
		return 0, errors.New("Id cannot be parsed")
	}
	return id, nil
}

func (s *mainRepository) addUser(organizationId, userId int64, role string) error {
	_, err := s.execer.DoQuery(addUserQuery, organizationId, userId, role)
	return err
}

func (s *mainRepository) deleteUser(organizationId, userId int64) error {
	_, err := s.execer.DoQuery(deleteUserFromGroupQuery, organizationId, userId)
	return err
}
