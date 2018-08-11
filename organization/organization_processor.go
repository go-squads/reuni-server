package organization

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/reuni-server/helper"
)

type processor interface {
	createNewOrganizationProcessor(orginizationName string, userId int64) error
	addUserProcessor(member *Member) error
	deleteUserFromGroupProcessor(organizationId, userId int64) error
	updateRoleOfUserProcessor(member *Member) error
	getAllMemberOfOrganizationProcessor(organizationId int64) ([]map[string]interface{}, error)
	getAllOrganizationProcessor(userId int) (string, error)
	translateNameToIdProcessor(name string) (int, error)
}

type mainProcessor struct {
	repo repository
}

func (s *mainProcessor) createNewOrganizationProcessor(organizationName string, userId int64) error {
	id, err := s.repo.createNewOrganization(organizationName)
	if err != nil {
		return err
	}
	return s.repo.addUser(id, userId, "Admin")
}

func (s *mainProcessor) addUserProcessor(member *Member) error {
	if !member.isRoleValid() {
		return helper.NewHttpError(http.StatusBadRequest, "Role is not valid")
	}
	return s.repo.addUser(member.OrgId, member.UserId, member.Role)
}

func (s *mainProcessor) deleteUserFromGroupProcessor(organizationId, userId int64) error {
	return s.repo.deleteUser(organizationId, userId)
}

func (s *mainProcessor) updateRoleOfUserProcessor(member *Member) error {
	if !member.isRoleValid() {
		return helper.NewHttpError(http.StatusBadRequest, "New role is not valid")
	}
	return s.repo.updateRoleOfUser(member.Role, member.OrgId, member.UserId)
}

func (s *mainProcessor) getAllMemberOfOrganizationProcessor(organizationId int64) ([]map[string]interface{}, error) {
	return s.repo.getAllMemberOfOrganization(organizationId)
}

func (s *mainProcessor) getAllOrganizationProcessor(userId int) (string, error) {
	res, err := s.repo.getAllOrganization(userId)
	if err != nil {
		return "", err
	}
	if res == nil {
		return "[]", nil
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		return "", err
	}
	return string(resJSON), nil
}

func (p *mainProcessor) translateNameToIdProcessor(name string) (int, error) {
	return p.repo.translateNameToIdRepository(name)
}

func TranslateNameToIdProcessor(q helper.QueryExecuter, name string) (int, error) {
	rep := initRepository(q)
	return rep.translateNameToIdRepository(name)
}
