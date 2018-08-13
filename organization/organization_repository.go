package organization

import (
	"errors"

	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

type repository interface {
	createNewOrganization(organization_name string) (int64, error)
	addUser(organizationId, userId int64, role string) error
	deleteUser(organizationId, userId int64) error
	updateRoleOfUser(newRole string, organizationId, userId int64) error
	getAllMemberOfOrganization(organizationId int64) ([]map[string]interface{}, error)
	getAllOrganization(userId int) ([]OrganizationMember, error)
	translateNameToIdRepository(organizationName string) (int, error)
}

type mainRepository struct {
	execer helper.QueryExecuter
}

const (
	createNewOrganizationQuery      = "INSERT INTO organization(name) VALUES($1) RETURNING id"
	addUserQuery                    = "INSERT INTO organization_member(organization_id, user_id, role) VALUES ($1,$2,$3) "
	deleteUserFromGroupQuery        = "DELETE FROM organization_member where organization_id=$1 and user_id=$2"
	updateRoleOfUserQuery           = "UPDATE organization_member SET role=$1 WHERE organization_id=$2 and user_id=$3"
	getAllMemberOfOrganizationQuery = "SELECT U.username, OM.role, OM.created_at FROM organization_member OM, users U WHERE OM.user_id = U.id AND OM.organization_id=$1"
	getAllOrganizationQuery         = "SELECT O.id,O.name,OM.role FROM organization O, organization_member OM WHERE O.id=OM.organization_id AND OM.user_id = $1"
	translateNameToIdQuery          = "SELECT id FROM organization WHERE name = $1"
)

func initRepository(execer helper.QueryExecuter) *mainRepository {
	return &mainRepository{
		execer: execer,
	}
}

func (s *mainRepository) createNewOrganization(organizationName string) (int64, error) {
	data, err := s.execer.DoQueryRow(createNewOrganizationQuery, organizationName)
	if err != nil {
		return 0, helper.NewHttpError(http.StatusConflict, "Already Exists")
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

func (s *mainRepository) updateRoleOfUser(newRole string, organizationId, userId int64) error {
	_, err := s.execer.DoQuery(updateRoleOfUserQuery, newRole, organizationId, userId)
	return err
}

func (s *mainRepository) getAllMemberOfOrganization(organizationId int64) ([]map[string]interface{}, error) {
	data, err := s.execer.DoQuery(getAllMemberOfOrganizationQuery, organizationId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *mainRepository) getAllOrganization(userId int) ([]OrganizationMember, error) {
	data, err := s.execer.DoQuery(getAllOrganizationQuery, userId)
	if err != nil {
		return nil, err
	}
	var orgs []OrganizationMember
	err = helper.ParseMap(data, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (s *mainRepository) translateNameToIdRepository(organizationName string) (int, error) {
	data, err := s.execer.DoQueryRow(translateNameToIdQuery, organizationName)
	if err != nil {
		return 0, err
	}
	id := int(data["id"].(int64))
	return id, nil
}
