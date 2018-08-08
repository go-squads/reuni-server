package organization

import (
	"net/http"

	"github.com/go-squads/reuni-server/appcontext"
	"github.com/go-squads/reuni-server/helper"
)

type processor interface {
	createNewOrganizationProcessor(orginizationName string, userId int64) error
	addUserProcessor(member *Member) error
	deleteUserFromGroupProcessor(organizationId, userId int64) error
	updateRoleOfUserProcessor(member *Member) error
}

type mainProcessor struct{}

var activeRepository repository

func getRepository() repository {
	if activeRepository == nil {
		activeRepository = initRepository(appcontext.GetHelper())
	}
	return activeRepository
}

func (s *mainProcessor) createNewOrganizationProcessor(organizationName string, userId int64) error {
	id, err := getRepository().createNewOrganization(organizationName)
	if err != nil {
		return err
	}
	return getRepository().addUser(id, userId, "Admin")
}

func (s *mainProcessor) addUserProcessor(member *Member) error {
	if !member.isRoleValid() {
		return helper.NewHttpError(http.StatusBadRequest, "Role is not valid")
	}
	return getRepository().addUser(member.OrgId, member.UserId, member.Role)
}

func (s *mainProcessor) deleteUserFromGroupProcessor(organizationId, userId int64) error {
	return getRepository().deleteUser(organizationId, userId)
}

func (s *mainProcessor) updateRoleOfUserProcessor(member *Member) error {
	if !member.isRoleValid() {
		return helper.NewHttpError(http.StatusBadRequest, "New role is not valid")
	}
	return getRepository().updateRoleOfUser(member.Role, member.OrgId, member.UserId)
}
